package handlers

import (
    "context"
    "log"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "task-management/internal/models"
    "task-management/internal/database"
    "task-management/internal/services"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    task.ID = primitive.NewObjectID()
    task.CreatedBy = userID
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    _, err := collection.InsertOne(context.Background(), task)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }

    // Generate AI suggestions if configured
    aiService, err := services.NewAIService(os.Getenv("OPENAI_API_KEY"))
    if err == nil {
        suggestions, err := aiService.GenerateTaskSuggestions(task)
        if err != nil {
            log.Printf("Error generating AI suggestions: %v", err)
        } else {
            // Store AI suggestions
            suggCollection := database.Client.Database("taskmanagement").Collection("ai_suggestions")
            suggCollection.InsertOne(context.Background(), models.AITaskSuggestion{
                TaskID:      task.ID,
                Suggestion:  suggestions,
                GeneratedAt: time.Now(),
            })
        }
    }

    // Broadcast task creation to all connected clients
    services.BroadcastMessage("task_created", gin.H{
        "task": task,
        "created_by": userID,
        "timestamp": time.Now(),
    })

    c.JSON(201, task)
}

func UpdateTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var updateData models.Task
    
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    updateData.UpdatedAt = time.Now()

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    result, err := collection.UpdateOne(
        context.Background(),
        bson.M{"_id": taskID},
        bson.M{"$set": updateData},
    )

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to update task"})
        return
    }

    if result.ModifiedCount > 0 {
        // Broadcast task update
        services.BroadcastMessage("task_updated", gin.H{
            "task_id": taskID,
            "updates": updateData,
            "timestamp": time.Now(),
        })
    }

    c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    result, err := collection.DeleteOne(context.Background(), bson.M{"_id": taskID})

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete task"})
        return
    }

    if result.DeletedCount > 0 {
        // Broadcast task deletion
        services.BroadcastMessage("task_deleted", gin.H{
            "task_id": taskID,
            "timestamp": time.Now(),
        })
    }

    c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func GetTasks(c *gin.Context) {
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    
    collection := database.Client.Database("taskmanagement").Collection("tasks")
    cursor, err := collection.Find(context.Background(), bson.M{"assigned_to": userID})
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
        return
    }

    var tasks []models.Task
    if err = cursor.All(context.Background(), &tasks); err != nil {
        c.JSON(500, gin.H{"error": "Failed to decode tasks"})
        return
    }

    c.JSON(200, tasks)
}

func GetAISuggestions(c *gin.Context) {
    taskID, err := primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid task ID"})
        return
    }

    // Get task details
    collection := database.Client.Database("taskmanagement").Collection("tasks")
    var task models.Task
    err = collection.FindOne(context.Background(), bson.M{"_id": taskID}).Decode(&task)
    if err != nil {
        c.JSON(404, gin.H{"error": "Task not found"})
        return
    }

    // Generate AI suggestions
    aiService, err := services.NewAIService(os.Getenv("OPENAI_API_KEY"))
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to initialize AI service"})
        return
    }

    suggestions, err := aiService.GenerateTaskSuggestions(task)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate suggestions"})
        return
    }

    // Store and return suggestions
    suggCollection := database.Client.Database("taskmanagement").Collection("ai_suggestions")
    suggestion := models.AITaskSuggestion{
        TaskID:      taskID,
        Suggestion:  suggestions,
        GeneratedAt: time.Now(),
    }

    _, err = suggCollection.InsertOne(context.Background(), suggestion)
    if err != nil {
        log.Printf("Error storing AI suggestions: %v", err)
    }

    c.JSON(200, suggestion)
}
package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "task-management/internal/models"
    "task-management/internal/database"
    "task-management/internal/services"
    "time"
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

    // Generate AI suggestions
    aiService := services.NewAIService(os.Getenv("GEMINI_API_KEY"))
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

    // Broadcast task creation to all connected clients
    wsService.BroadcastUpdate(gin.H{
        "type": "task_created",
        "task": task,
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
        wsService.BroadcastUpdate(gin.H{
            "type": "task_updated",
            "task_id": taskID,
            "updates": updateData,
        })
    }

    c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func GetTaskSuggestions(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    
    collection := database.Client.Database("taskmanagement").Collection("ai_suggestions")
    var suggestion models.AITaskSuggestion
    err := collection.FindOne(
        context.Background(),
        bson.M{"task_id": taskID},
    ).Decode(&suggestion)

    if err != nil {
        c.JSON(404, gin.H{"error": "No suggestions found"})
        return
    }

    c.JSON(200, suggestion)
}
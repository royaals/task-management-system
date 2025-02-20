package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "task-management/internal/models"
    "task-management/internal/database"  // Add this import
    "time"
)

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

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    task.AssignedTo = userID
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    result, err := collection.InsertOne(context.Background(), task)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }

    task.ID = result.InsertedID.(primitive.ObjectID)
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
    _, err := collection.UpdateOne(
        context.Background(),
        bson.M{"_id": taskID},
        bson.M{"$set": updateData},
    )

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to update task"})
        return
    }

    c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": taskID})

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete task"})
        return
    }

    c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
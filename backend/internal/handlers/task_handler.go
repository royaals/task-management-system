package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "task-management/internal/models"
    "time"
)

func GetTasks(c *gin.Context) {
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    
    collection := client.Database("taskmanagement").Collection("tasks")
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

    collection := client.Database("taskmanagement").Collection("tasks")
    result, err := collection.InsertOne(context.Background(), task)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }

    task.ID = result.InsertedID.(primitive.ObjectID)
    c.JSON(201, task)
}
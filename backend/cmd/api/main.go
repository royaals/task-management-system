// cmd/api/main.go
package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "task-management/internal/handlers"
    "task-management/internal/middleware"
    "task-management/internal/database"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    database.InitDatabase()

    // Initialize Gin router
    r := gin.Default()
    
    // Apply CORS middleware
    r.Use(middleware.CORSMiddleware())

    // API routes
    api := r.Group("/api")
    {
        // Auth routes
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)
        api.POST("/logout", handlers.Logout)
        api.GET("/me", handlers.GetMe)

        // Task routes
        api.GET("/tasks", handlers.GetTasks)
        api.POST("/tasks", handlers.CreateTask)
        api.PUT("/tasks/:id", handlers.UpdateTask)
        api.DELETE("/tasks/:id", handlers.DeleteTask)
        
        // AI routes
        api.POST("/ai/suggestions", handlers.GetAISuggestions)
    }

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
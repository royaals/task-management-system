package main

import (
    "fmt"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "task-management/internal/handlers"
    "task-management/internal/middleware"
    "task-management/internal/database"
)

func main() {
    // Load .env only in development
    if os.Getenv("GO_ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found")
        }
    }

    // Initialize database
    database.InitDatabase()

    // Initialize Gin
    r := gin.Default()
    
    // Add middleware
    r.Use(middleware.CORSMiddleware())

    // Add health check endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "healthy",
            "message": "Server is running",
        })
    })

    // API routes
    api := r.Group("/api")
    {
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)
        api.POST("/logout", handlers.Logout)
        api.GET("/me", handlers.GetMe)
        api.GET("/tasks", handlers.GetTasks)
        api.POST("/tasks", handlers.CreateTask)
        api.PUT("/tasks/:id", handlers.UpdateTask)
        api.DELETE("/tasks/:id", handlers.DeleteTask)
        api.POST("/ai/suggestions", handlers.GetAISuggestions)
    }

    // Get port from environment
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Start server with explicit host binding
    address := fmt.Sprintf("0.0.0.0:%s", port)
    log.Printf("Server starting on %s", address)
    if err := r.Run(address); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
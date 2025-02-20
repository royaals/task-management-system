package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "task-management/internal/handlers"
    "task-management/internal/middleware"
    "task-management/internal/database"
    "task-management/internal/services"  // Add this import
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    database.InitDatabase()

    // Start WebSocket hub
    go services.WebsocketHub.Run()

    // Initialize Gin router
    r := gin.Default()
    r.Use(middleware.CORSMiddleware())

    // API routes
    api := r.Group("/api")
    {
        // Public routes
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            // Task routes
            protected.GET("/tasks", handlers.GetTasks)
            protected.POST("/tasks", handlers.CreateTask)
            protected.PUT("/tasks/:id", handlers.UpdateTask)
            protected.DELETE("/tasks/:id", handlers.DeleteTask)
            
            // AI suggestions routes
            protected.GET("/tasks/:id/suggestions", handlers.GetAISuggestions)
            protected.POST("/tasks/:id/suggestions", handlers.GetAISuggestions)
            
            // WebSocket connection
            protected.GET("/ws", handlers.HandleWebSocket)
        }
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
package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "task-management/internal/handlers"
    "task-management/internal/middleware"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize router
    r := gin.Default()

    // Add CORS middleware
    r.Use(middleware.CORSMiddleware())

    // Routes
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
        }
    }

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
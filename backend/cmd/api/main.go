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

    database.InitDatabase()

    r := gin.Default()
    r.Use(middleware.CORSMiddleware())

    api := r.Group("/api")
    {
        // Auth routes
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
    protected.POST("/tasks/:id/suggestions", handlers.GetAISuggestions) // Generate new suggestions
            
            // WebSocket connection
            protected.GET("/ws", handlers.HandleWebSocket)
        }
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
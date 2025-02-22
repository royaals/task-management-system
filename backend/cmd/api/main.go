
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

    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
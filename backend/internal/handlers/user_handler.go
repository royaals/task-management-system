package handlers

import (
    "context"
    "log"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
    "task-management/internal/models"
    "task-management/internal/database"
)

func Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Log the received data
    log.Printf("Registering user with email: %s", input.Email)
    log.Printf("Original password: %s", input.Password)

    // Generate password hash
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

    log.Printf("Hashed password: %s", string(hashedPassword))

    // Create user
    user := models.User{
        ID:       primitive.NewObjectID(),
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    // Save to database
    collection := database.Client.Database("taskmanagement").Collection("users")
    _, err = collection.InsertOne(context.Background(), user)
    if err != nil {
        log.Printf("Error saving user: %v", err)
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    log.Printf("User registered successfully with ID: %s", user.ID.Hex())

    c.JSON(201, gin.H{
        "id": user.ID.Hex(),
        "email": user.Email,
    })
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

  

    // Find user
    var user models.User
    collection := database.Client.Database("taskmanagement").Collection("users")
    err := collection.FindOne(context.Background(), bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        log.Printf("User not found: %v", err)
        c.JSON(401, gin.H{"error": "Invalid credentials"})
        return
    }

  

    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        log.Printf("Password comparison failed: %v", err)
        c.JSON(401, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID.Hex(),
        "email":   user.Email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        log.Printf("Error generating token: %v", err)
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    

    c.JSON(200, gin.H{
        "token": tokenString,
        "user": gin.H{
            "id":    user.ID.Hex(),
            "email": user.Email,
        },
    })
}
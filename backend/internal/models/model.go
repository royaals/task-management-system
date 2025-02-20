package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string            `bson:"email" json:"email"`
    Password string            `bson:"password" json:"-"`
    Name     string            `bson:"name" json:"name"`
}

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string            `bson:"title" json:"title"`
    Description string            `bson:"description" json:"description"`
    Status      string            `bson:"status" json:"status"`
    Priority    string            `bson:"priority" json:"priority"`
    DueDate     time.Time         `bson:"due_date" json:"due_date"`
    AssignedTo  primitive.ObjectID `bson:"assigned_to" json:"assigned_to"`
    CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
    CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time         `bson:"updated_at" json:"updated_at"`
    Tags        []string          `bson:"tags" json:"tags"`
}

type Notification struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
    Message   string            `bson:"message" json:"message"`
    Type      string            `bson:"type" json:"type"`
    Read      bool              `bson:"read" json:"read"`
    CreatedAt time.Time         `bson:"created_at" json:"created_at"`
}

type AITaskSuggestion struct {
    TaskID      primitive.ObjectID `bson:"task_id" json:"task_id"`
    Suggestion  string            `bson:"suggestion" json:"suggestion"`
    GeneratedAt time.Time         `bson:"generated_at" json:"generated_at"`
}
package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string            `bson:"email" json:"email"`
    Password string            `bson:"password" json:"-"`
}

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string            `bson:"title" json:"title"`
    Description string            `bson:"description" json:"description"`
    Status      string            `bson:"status" json:"status"`
    AssignedTo  primitive.ObjectID `bson:"assigned_to" json:"assigned_to"`
    CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time         `bson:"updated_at" json:"updated_at"`
}
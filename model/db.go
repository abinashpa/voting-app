package model

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

// DB of MySQL
var DB *sql.DB

// Rdb redis client
var Rdb *redis.Client

// Ctx bg for redis
var Ctx = context.Background()

// User struct
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Session store in redis
type Session struct {
	Email  string
	UserID uint8
}

// Poll structure
type Poll struct {
	Title   string `json:"title"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
	Option5 string `json:"option5"`
}

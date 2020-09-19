package model

import (
	"database/sql"
)

// DB of MySQL
var DB *sql.DB

// User struct
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

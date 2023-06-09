// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"encoding/json"
)

type Role struct {
	ID          int32           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Permissions json.RawMessage `json:"permissions"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	UpdatedAt   sql.NullTime    `json:"updated_at"`
}

type User struct {
	ID        int32        `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Phone     string       `json:"phone"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type UserRole struct {
	ID        int32        `json:"id"`
	UserID    int32        `json:"user_id"`
	RoleID    int32        `json:"role_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

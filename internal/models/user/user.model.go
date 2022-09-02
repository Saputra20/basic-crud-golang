package usermdl

import (
	"basic-crud/internal/datastore"
	"time"
)

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Role      datastore.Role
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Name   string `json:"name" binding:"required"`
	RoleID int    `json:"role_id" binding:"required"`
}

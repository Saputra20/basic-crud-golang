package rolemdl

import "time"

type Role struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type RoleRequest struct {
	Name string `json:"name" binding:"required"`
}

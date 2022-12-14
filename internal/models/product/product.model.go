package productmdl

import "time"

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductRequest struct {
	Name string `json:"name" binding:"required"`
}

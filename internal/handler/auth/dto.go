package auth

import (
	"time"
)

// ===== REQUEST DTOs =====

type RegisterRequest struct {
	Name                 string `json:"name" binding:"required,min=2,max=100"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

// ===== RESPONSE DTOs =====

type RegisterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

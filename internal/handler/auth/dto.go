package auth

import (
	"net/mail"
	"strings"
	"time"

	"github.com/hogiabao7725/go-ticket-engine/internal/response"
)

// ===== REQUEST DTOs =====

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterRequest) Validate() []response.ValidationDetail {
	var errs []response.ValidationDetail

	if strings.TrimSpace(req.Name) == "" {
		errs = append(errs, response.ValidationDetail{Field: "Name", Reason: "This field is required"})
	} else if len(req.Name) < 2 {
		errs = append(errs, response.ValidationDetail{Field: "Name", Reason: "Must be at least 2 characters"})
	}

	if strings.TrimSpace(req.Email) == "" {
		errs = append(errs, response.ValidationDetail{Field: "Email", Reason: "This field is required"})
	} else if _, err := mail.ParseAddress(req.Email); err != nil {
		errs = append(errs, response.ValidationDetail{Field: "Email", Reason: "Invalid email format"})
	}

	if len(req.Password) < 6 {
		errs = append(errs, response.ValidationDetail{Field: "Password", Reason: "Must be at least 6 characters"})
	}

	return errs
}

// ===== RESPONSE DTOs =====

type RegisterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

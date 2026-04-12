package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/hogiabao7725/go-ticket-engine/pkg/hash"
)

type User struct {
	id        string
	name      string
	email     string
	password  string // hashed
	role      string // user, organizer, admin
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name, email, plainPassword string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if len(plainPassword) < 6 {
		return nil, ErrWeakPassword
	}

	hashedPassword, err := hash.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &User{
		id:        uuid.New().String(),
		name:      name,
		email:     email,
		password:  hashedPassword,
		role:      "user",
		createdAt: now,
		updatedAt: now,
	}, nil
}

func ReconstructUser(id, name, email, hashedPassword, role string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		email:     email,
		password:  hashedPassword,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) VerifyPassword(plain string) bool {
	return hash.ComparePassword(u.password, plain) == nil
}

func (u *User) ID() string { return u.id }

func (u *User) Name() string { return u.name }

func (u *User) Email() string { return u.email }

func (u *User) Role() string { return u.role }

func (u *User) PasswordHash() string { return u.password }

func (u *User) CreatedAt() time.Time { return u.createdAt }

func (u *User) UpdatedAt() time.Time { return u.updatedAt }

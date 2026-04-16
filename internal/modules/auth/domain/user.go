package domain

import (
	"time"
)

type User struct {
	id        string
	name      Name
	email     Email
	password  HashedPassword
	role      Role
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id string, name Name, email Email, hashedPassword HashedPassword, role Role) (*User, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	now := time.Now()

	return &User{
		id:        id,
		name:      name,
		email:     email,
		password:  hashedPassword,
		role:      role,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func ReconstructUser(id string, name Name, email Email, hashedPassword HashedPassword, role Role, createdAt, updatedAt time.Time) *User {
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

func (u *User) ID() string { return u.id }

func (u *User) Name() Name { return u.name }

func (u *User) Email() Email { return u.email }

func (u *User) Role() Role { return u.role }

func (u *User) PasswordHash() string { return u.password.Value() }

func (u *User) CreatedAt() time.Time { return u.createdAt }

func (u *User) UpdatedAt() time.Time { return u.updatedAt }

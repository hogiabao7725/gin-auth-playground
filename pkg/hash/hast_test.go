package hash

import (
	"testing"

	"github.com/hogiabao7725/go-ticket-engine/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// Table of test cases
	tests := []struct {
		name     string
		password string
		wantErr  bool
		errCode  string
	}{
		{
			name:     "Hash Password Successfully",
			password: "1234",
			wantErr:  false,
		},
		{
			name:     "Hash Empty Password",
			password: "",
			wantErr:  true,
			errCode:  apperror.ErrPasswordEmpty.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := HashPassword(tt.password)
			if tt.wantErr {
				assert.Error(t, err)

				if appErr, ok := err.(*apperror.AppError); ok {
					assert.Equal(t, tt.errCode, appErr.Code)
				} else {
					t.Errorf("expected AppError, got %T", err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hashed)

				// Check if the hashed password contains the bcrypt prefix
				assert.Contains(t, hashed, "$2")
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	// Pre-hash a password for testing
	validPass := "1234"
	validHashed, _ := HashPassword(validPass)

	tests := []struct {
		name       string
		hashedPass string
		password   string
		wantErr    bool
		errCode    string
	}{
		{
			name:       "Correct Password",
			hashedPass: validHashed,
			password:   validPass,
			wantErr:    false,
		},
		{
			name:       "Wrong Password",
			hashedPass: validHashed,
			password:   "wrong",
			wantErr:    true,
			errCode:    apperror.ErrInvalidCredentials.Code,
		},
		{
			name:       "Invalid Hashed Password",
			hashedPass: "invalid-hash",
			password:   validPass,
			wantErr:    true,
			errCode:    apperror.ErrComparePasswordFailed.Code,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ComparePassword(tt.hashedPass, tt.password)
			if tt.wantErr {
				assert.Error(t, err)

				if appErr, ok := err.(*apperror.AppError); ok {
					assert.Equal(t, tt.errCode, appErr.Code)
				} else {
					t.Errorf("expected AppError, got %T", err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

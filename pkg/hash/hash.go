package hash

import (
	"strings"

	"github.com/hogiabao7725/go-ticket-engine/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	valPassword, err := validatePassword(password)
	if err != nil {
		return "", err
	}

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(valPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", apperror.Wrap(
			err,
			apperror.ErrHashFailed.Code,
			apperror.ErrHashFailed.Message,
			apperror.ErrHashFailed.StatusCode,
		)
	}
	return string(hashedByte), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return apperror.ErrInvalidCredentials
		}

		return apperror.Wrap(
			err,
			apperror.ErrComparePasswordFailed.Code,
			apperror.ErrComparePasswordFailed.Message,
			apperror.ErrComparePasswordFailed.StatusCode,
		)
	}
	return nil
}

func validatePassword(password string) (string, error) {
	valPassword := strings.TrimSpace(password)
	if valPassword == "" {
		return "", apperror.ErrPasswordEmpty
	}
	return valPassword, nil
}

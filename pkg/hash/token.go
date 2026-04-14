package hash

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrRefreshTokenEmpty   = errors.New("refresh token cannot be empty")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

func HashRefreshToken(refreshToken string) (string, error) {
	valRefreshToken := strings.TrimSpace(refreshToken)
	if valRefreshToken == "" {
		return "", ErrRefreshTokenEmpty
	}

	sum := sha256.Sum256([]byte(valRefreshToken))
	return hex.EncodeToString(sum[:]), nil
}

func CompareRefreshToken(hashedRefreshToken, refreshToken string) error {
	valHashedRefreshToken := strings.TrimSpace(hashedRefreshToken)
	if valHashedRefreshToken == "" {
		return ErrInvalidRefreshToken
	}

	calculatedHash, err := HashRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare([]byte(valHashedRefreshToken), []byte(calculatedHash)) != 1 {
		return ErrInvalidRefreshToken
	}

	return nil
}

func MustHashRefreshToken(refreshToken string) string {
	hashedRefreshToken, err := HashRefreshToken(refreshToken)
	if err != nil {
		panic(fmt.Errorf("failed to hash refresh token: %w", err))
	}

	return hashedRefreshToken
}

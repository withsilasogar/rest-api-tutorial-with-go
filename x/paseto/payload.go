package paseto

import (
	"errors"
	"test01/x/interfacesx"
	"time"
)

type AuthenticationPayload struct {
	User      interfacesx.UserData `json:"user"`
	IssuedAt  time.Time            `json:"issuedAt"`
	ExpiresAt time.Time            `json:"expiresAt"`
}

var (
	ErrExpiredToken = errors.New("token is expired")
)

func NewAuthenticationPayload(user interfacesx.UserData, duration time.Duration) (
	*AuthenticationPayload, error,
) {
	payload := &AuthenticationPayload{
		User:      user,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *AuthenticationPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}

	return nil
}

package paseto

import (
	"test01/x/interfacesx"
	"time"
)

type Generator interface {
	GenerateToken(user interfacesx.UserData, duration time.Duration) (string, *AuthenticationPayload, error) //token, payload, error
	VerifyToken(token string) (*AuthenticationPayload, error)
}

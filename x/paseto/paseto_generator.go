package paseto

import (
	"fmt"
	"test01/x/interfacesx"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoGenerator struct {
	paseto      *paseto.V2
	symetricKey []byte
}

func NewPasetoGenerator(symetricKey string) (Generator, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(
			"invalid symetric key size, it should be %d characters",
			chacha20poly1305.KeySize,
		)
	}

	generator := &PasetoGenerator{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}

	return generator, nil
}

func (generator *PasetoGenerator) GenerateToken(
	user interfacesx.UserData,
	duration time.Duration,
) (string, *AuthenticationPayload, error) {
	payload, err := NewAuthenticationPayload(user, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := generator.paseto.Encrypt(generator.symetricKey, payload, nil)
	if err != nil {
		return "", payload, err
	}

	return token, payload, nil
}

func (generator *PasetoGenerator) VerifyToken(token string) (*AuthenticationPayload, error) {
	payload := &AuthenticationPayload{}

	err := generator.paseto.Decrypt(token, generator.symetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

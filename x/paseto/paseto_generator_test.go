package paseto

import (
	"test01/x/interfacesx"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoGenerator(t *testing.T) {
	symetricKey := "0123456789abcdef0123456789abcdef"

	generator, err := NewPasetoGenerator(symetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, generator)

	user := interfacesx.UserData{
		Username: "testuser",
	}
	duration := time.Minute

	// Creating new Token
	token, payload, err := generator.GenerateToken(user, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotNil(t, payload)

	// Verify token
	verifiedToken, err := generator.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, verifiedToken)

	// Check if the payload matches
	require.Equal(t, payload.User.Username, verifiedToken.User.Username)
	require.WithinDuration(t, payload.ExpiresAt, verifiedToken.ExpiresAt, time.Second)
	require.WithinDuration(t, payload.IssuedAt, verifiedToken.IssuedAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	symetricKey := "0123456789abcdef0123456789abcdef"

	generator, err := NewPasetoGenerator(symetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, generator)

	user := interfacesx.UserData{
		Username: "testuser",
	}
	duration := -time.Minute

	//Creating a token
	token, payload, err := generator.GenerateToken(user, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotNil(t, payload)

	// Verify the token is expired
	verifiedToken, err := generator.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, verifiedToken)
	require.EqualError(t, err, ErrExpiredToken.Error())
}

//Test for InvalidPasetoToken

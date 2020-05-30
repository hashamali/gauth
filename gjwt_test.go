package gjwt

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	secret := uuid.New().String()
	jwt := JWTAuth{
		Secret: secret,
	}

	userID := uuid.New()
	token, err := jwt.Create(userID.String(), "test", uint(10))
	assert.NoError(t, err)

	extracted, err := jwt.ExtractFromHeader("Bearer " + token)
	assert.NoError(t, err)

	extractedUserIDString, ok := extracted.(string)
	assert.True(t, ok)

	extractedUserID, err := uuid.Parse(extractedUserIDString)
	assert.NoError(t, err)

	assert.Equal(t, userID, extractedUserID)
}

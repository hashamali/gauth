package gauth

import (
	"github.com/hashamali/grand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	username := grand.RandomString(10)
	password := grand.RandomString(10)

	ba := StaticBasicAuth{
		Username: username,
		Password: password,
	}

	otherUsername := grand.RandomString(8)
	otherPassword := grand.RandomString(8)
	assert.Equal(t, false, ba.Validate(otherUsername, otherPassword))
	assert.Equal(t, true, ba.Validate(username, password))
}

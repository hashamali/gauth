package gauth

import "crypto/subtle"

// BasicAuth defines an interface for interaction with the Basic Auth protocol.
type BasicAuth interface {
	Validate(username string, password string) bool
}

// StaticBasicAuth takes a static username/password combo.
type StaticBasicAuth struct {
	Username string
	Password string
}

// Validate that the given username and password are correct.
func (auth *StaticBasicAuth) Validate(username string, password string) bool {
	userMatches := subtle.ConstantTimeCompare([]byte(username), []byte(auth.Username)) == 1
	passwordMatches := subtle.ConstantTimeCompare([]byte(password), []byte(auth.Password)) == 1
	return userMatches && passwordMatches
}

// GetStaticBasicAuth will return a StaticBasicAuth.
func GetStaticBasicAuth() (*StaticBasicAuth, error) {
	c, err := getStaticBasicAuthConfig()
	if err != nil {
		return nil, err
	}

	return &StaticBasicAuth{
		Username: c.User,
		Password: c.Password,
	}, nil
}

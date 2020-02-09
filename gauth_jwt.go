package gauth

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTAuth handles encoding and decoding JWT tokens.
type JWTAuth struct {
	Secret string
}

// Create creates a JWT token for a given user ID.
func (auth *JWTAuth) Create(meta interface{}, issuer string, expirationInMinutes uint) (string, error) {
	ttl := time.Duration(expirationInMinutes) * time.Minute
	return auth.createTokenString(meta, issuer, ttl)
}

// ExtractFromHeader returns the encoded object from the header.
func (auth *JWTAuth) ExtractFromHeader(tokenHeader string) (interface{}, error) {
	if tokenHeader != "" {
		split := strings.Split(tokenHeader, " ")
		if len(split) == 2 {
			meta, err := auth.Extract(split[1])
			if err != nil {
				return nil, err
			}

			return meta, nil
		}
	}

	return nil, errors.New("auth: invalid auth token")
}

// Extract will take a raw token and extract the value.
func (auth *JWTAuth) Extract(rawToken string) (interface{}, error) {
	tk := &customJWT{}
	token, err := jwt.ParseWithClaims(rawToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Secret), nil
	})

	if err == nil && token.Valid {
		return tk.Meta, nil
	}

	return nil, err
}

// GetJWTAuth will return a JWTAuth.
func GetJWTAuth() (*JWTAuth, error) {
	jwtConfig, err := getJWTConfig()
	if err != nil {
		return nil, err
	}

	return &JWTAuth{Secret: jwtConfig.Secret}, nil
}

func (auth *JWTAuth) createTokenString(meta interface{}, issuer string, ttl time.Duration) (string, error) {
	token := &customJWT{
		Meta: meta,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token).SignedString([]byte(auth.Secret))
}

type customJWT struct {
	Meta interface{}
	jwt.StandardClaims
}

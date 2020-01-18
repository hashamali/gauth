package gauth

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTAuth handles encoding and decoding JWT tokens.
type JWTAuth struct {
	Secret            string
	ExpirationInHours uint
}

// Create creates a JWT token for a given user ID.
func (auth *JWTAuth) Create(meta interface{}, issuer string) (string, error) {
	ttl := time.Duration(auth.ExpirationInHours) * time.Hour
	return auth.createTokenString(meta, issuer, ttl)
}

// Extract returns the encoded object from the header.
func (auth *JWTAuth) Extract(tokenHeader string) (interface{}, string, error) {
	if tokenHeader != "" {
		split := strings.Split(tokenHeader, " ")
		if len(split) == 2 {
			tk := &customJWT{}
			token, err := jwt.ParseWithClaims(split[1], tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.Secret), nil
			})

			if err == nil && token.Valid {
				return tk.Meta, split[1], nil
			}

			return nil, "", err
		}
	}

	return nil, "", errors.New("auth: invalid auth token")
}

// GetJWTAuth will return a JWTAuth.
func GetJWTAuth() (*JWTAuth, error) {
	jwtConfig, err := getJWTConfig()
	if err != nil {
		return nil, err
	}

	return &JWTAuth{
		Secret:            jwtConfig.Secret,
		ExpirationInHours: uint(jwtConfig.ExpirationInHours),
	}, nil
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

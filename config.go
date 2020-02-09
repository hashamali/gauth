package gauth

import (
	"github.com/kelseyhightower/envconfig"
)

type staticBasicAuthConfig struct {
	User     string `required:"true"`
	Password string `required:"true"`
}

type jwtConfig struct {
	Secret string `required:"true"`
}

func getStaticBasicAuthConfig() (*staticBasicAuthConfig, error) {
	c := staticBasicAuthConfig{}
	err := envconfig.Process("BASIC_AUTH", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func getJWTConfig() (*jwtConfig, error) {
	c := jwtConfig{}
	err := envconfig.Process("JWT_AUTH", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

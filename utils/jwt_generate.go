package utils

import (
	"time"

	"github.com/aunjaffery/teamsprint/config"
	"github.com/golang-jwt/jwt/v5"
)

func Generate_JWT(paylod string) (t string, error error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": now.Add(time.Hour * 2).Unix(),
		"sub": paylod,
		"iat": now.Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Envs.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

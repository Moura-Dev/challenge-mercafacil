package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secret   string
	issure   string
	customer string
}

type Claim struct {
	jwt.StandardClaims
}

func NewJWTService() *jwtService {
	return &jwtService{
		secret: "secret-key",
		issure: "contacts-api",
	}
}

func (s *jwtService) GenerateToken() (string, error) {

	claim := &Claim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.issure,
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token_encode, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return "", err
	}
	return token_encode, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secret), nil
	})

	return err == nil
}

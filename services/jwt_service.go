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

func (s *jwtService) GenerateToken(customer string) (string, error) {

	claim := &Claim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.issure,
			Subject:   customer,
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

func (s *jwtService) GetCustomerFromToken(t string) (string, error) {
	tokenString := t
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if token.Valid {
		return claims["sub"].(string), nil
	}
	return "", err
}

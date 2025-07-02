package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id, email string) (string, error) {
	expStr := os.Getenv("JWT_EXP") // misal: "1"
	expInt, err := strconv.Atoi(expStr)
	if err != nil {
		fmt.Printf("Invalid JWT_EXP env: %v", err)
	}
	jwtExp := time.Duration(expInt) * time.Hour

	tokenString, err := generateToken(id, email, jwtExp)
	if err != nil {
		return "", fmt.Errorf("error creating refresh token: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func CreateRefreshToken(id, email string) (string, error) {
	
	expStr := os.Getenv("JWT_REFRESH_EXP") // misal: "1"
	expInt, err := strconv.Atoi(expStr)
	if err != nil {
		fmt.Printf("Invalid JWT_REFRESH_EXP env: %v", err)
	}
	jwtExp := time.Duration(expInt) * time.Hour

	tokenString, err := generateToken(id, email, jwtExp)
	if err != nil {
		return "", fmt.Errorf("error creating refresh token: %w", err)
	}

	return tokenString, nil
}

func generateToken(id, email string, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(exp).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

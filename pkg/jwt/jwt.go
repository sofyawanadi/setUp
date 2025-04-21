// jwt.go
package jwt

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

func GenerateToken(userID int, secret string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

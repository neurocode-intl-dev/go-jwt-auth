package utils

import (
    "time"
    "os"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID int64) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(time.Minute * 15).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userID int64) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}

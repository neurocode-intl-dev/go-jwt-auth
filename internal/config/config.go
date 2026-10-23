package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppPort            string
	DBHost             string
	DBPort             int
	DBUser             string
	DBPassword         string
	DBName             string
	DatabaseURL        string
	JWTSecret          string
	RefreshSecret      string
	AccessTokenExpire  int
	RefreshTokenExpire int
}

func LoadConfig() (*Config, error) {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		dbPortStr = "5432"
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}

	// Build DATABASE_URL from individual components
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key"
	}

	refreshSecret := os.Getenv("REFRESH_SECRET")
	if refreshSecret == "" {
		refreshSecret = "your-refresh-secret-key"
	}

	accessTokenExpire := 15
	if val := os.Getenv("ACCESS_TOKEN_EXPIRE"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			accessTokenExpire = parsed
		}
	}

	refreshTokenExpire := 43200
	if val := os.Getenv("REFRESH_TOKEN_EXPIRE"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			refreshTokenExpire = parsed
		}
	}

	return &Config{
		AppPort:            appPort,
		DBHost:             dbHost,
		DBPort:             dbPort,
		DBUser:             dbUser,
		DBPassword:         dbPassword,
		DBName:             dbName,
		DatabaseURL:        databaseURL,
		JWTSecret:          jwtSecret,
		RefreshSecret:      refreshSecret,
		AccessTokenExpire:  accessTokenExpire,
		RefreshTokenExpire: refreshTokenExpire,
	}, nil
}

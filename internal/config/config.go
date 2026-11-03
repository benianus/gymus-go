package config

import (
	"context"
	"os"

	"github.com/benianus/gymus-go/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	database.Queries
	EnvironemntVariables
}

type EnvironemntVariables struct {
	Port      string
	DbUrl     string
	SecretKey string
	Platform  string
}

func (cfg *ApiConfig) LoadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	secretKey := os.Getenv("SECRET_KEY")
	platform := os.Getenv("PLATFORM")

	cfg.Port = port
	cfg.DbUrl = dbUrl
	cfg.SecretKey = secretKey
	cfg.Platform = platform

	return nil
}

func (cfg *ApiConfig) LoadDatabase() error {
	db, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		return err
	}
	dbQueries := database.New(db)

	cfg.Queries = *dbQueries

	return nil
}

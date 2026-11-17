package config

import (
	"database/sql"
	"os"

	"github.com/benianus/gymus-go/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	database.Queries
	EnvironemntVariables
}

type EnvironemntVariables struct {
	Port      string
	DbUrl     string
	Platform  string
	SecretKey string
	Issuer    string
	Audience  string
	Validate  *validator.Validate
}

func (cfg *ApiConfig) LoadEnvVariables() error {
	if err := godotenv.Load("../.env"); err != nil {
		return err
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secretKey := os.Getenv("SECRET_KEY")
	issuer := os.Getenv("ISSUER")
	audience := os.Getenv("AUDIENCE")

	cfg.Port = port
	cfg.DbUrl = dbUrl
	cfg.SecretKey = secretKey
	cfg.Platform = platform
	cfg.Issuer = issuer
	cfg.Audience = audience

	return nil
}

func (cfg *ApiConfig) LoadDatabase() (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DbUrl)

	if err != nil {
		return nil, err
	}
	dbQueries := database.New(db)

	cfg.Queries = *dbQueries

	return db, nil
}

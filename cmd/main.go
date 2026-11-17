package main

import (
	"encoding/json/v2"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/routers"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// initialize single instance of validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	cfg := &config.ApiConfig{
		Validate: validate,
	}

	/**
	 * load environement variables
	 */
	if err := cfg.LoadEnvVariables(); err != nil {
		log.Fatal(err)
		return
	}

	/**
	 * initilize database
	 */
	db, err := cfg.LoadDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	/**
	* initilize server
	* routing & handlers
	 */
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		json.MarshalWrite(w, map[string]string{"message": "Hello, World!"})
	})

	// Auth
	routers.Router(cfg, mux)
	/**
	 * initialize server
	 */
	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	/*
	 * log server informations
	 */
	log.Printf("Serving files from  on port: %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}

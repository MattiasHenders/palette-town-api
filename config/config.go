package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     *DBConfig
	Server *ServerConfig
	API    *APIConfig
}

type ServerConfig struct {
	Port string
	Salt string
}

type DBConfig struct {
	URI      string
	Username string
	Password string
	DBName   string
}

type APIConfig struct {
	ColorMindURL string
}

func GetConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Server config
	server_port := os.Getenv("SERVER_PORT")
	server_salt := os.Getenv("SERVER_SALT")

	// Get DB config
	db_uri := os.Getenv("DB_URI")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	// Get API config
	api_colormind_url := os.Getenv("API_COLORMIND_URL")

	return &Config{
		Server: &ServerConfig{
			Port: server_port,
			Salt: server_salt,
		},
		DB: &DBConfig{
			URI:      db_uri,
			Username: db_username,
			Password: db_password,
			DBName:   db_name,
		},
		API: &APIConfig{
			ColorMindURL: api_colormind_url,
		},
	}
}

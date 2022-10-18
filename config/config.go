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
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
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

	// Get DB config
	// TODO: set up db for users

	// Get API config
	api_colormind_url := os.Getenv("API_COLORMIND_URL")

	return &Config{
		Server: &ServerConfig{
			Port: server_port,
		},
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
		API: &APIConfig{
			ColorMindURL: api_colormind_url,
		},
	}
}

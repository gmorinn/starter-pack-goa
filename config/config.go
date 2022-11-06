package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// API represent api
type API struct {
	Name        string
	Mode        string
	Domain      string
	SSL         bool
	Host        string
	Port        int
	Cors        string
	Security    Security
	DatabaseURL string
	Database    Database
}

// Database represent database
type Database struct {
	Host     string
	Database string
	User     string
	Password string
	Port     int
}

// Security represent security
type Security struct {
	Secret               string
	AccessTokenDuration  int
	RefreshTokenDuration int
	OAuthID              string
	OAuthSecret          string
}

// New config api return config
func New() *API {
	var config API

	godotenv.Load("../.env")

	config.Mode = os.Getenv("ENV")
	config.Name = os.Getenv("PROJECT")
	config.Domain = os.Getenv("API_DOMAIN")

	if os.Getenv("PORT") == "" {
		config.Port, _ = getenvInt("API_PORT")
	} else {
		config.Port, _ = getenvInt("PORT")
	}

	config.Cors = os.Getenv("API_CORS")
	config.SSL, _ = getenvBool("API_SSL")
	config.Host = fmt.Sprintf("%s:%d", config.Domain, config.Port)

	config.Database.Host = os.Getenv("POSTGRES_HOST")
	config.Database.Database = os.Getenv("POSTGRES_DB")
	config.Database.User = os.Getenv("POSTGRES_USER")
	config.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Database.Port, _ = getenvInt("POSTGRES_PORT")

	config.Security.OAuthID = os.Getenv("API_OAUTH_ID")
	config.Security.OAuthSecret = os.Getenv("API_OAUTH_SECRET")

	config.Security.Secret = os.Getenv("API_SECRET")
	config.Security.AccessTokenDuration, _ = getenvInt("API_ACCESS_TOKEN")
	config.Security.RefreshTokenDuration, _ = getenvInt("API_REFRESH_TOKEN")

	if os.Getenv("ENV") == "PROD" {
		config.DatabaseURL = os.Getenv("DATABASE_URL")
	} else {
		config.DatabaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database)
	}

	return &config
}

func getenvInt(key string) (int, error) {
	s := os.Getenv(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Println("Env ", key, " : ", err.Error())
		return 0, err
	}
	return v, nil
}

func getenvBool(key string) (bool, error) {
	s := os.Getenv(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Println("Env ", key, " : ", err.Error())
		return false, err
	}
	return v, nil
}

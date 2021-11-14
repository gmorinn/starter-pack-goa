package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// API represent api
type API struct {
	Mode     string
	Domain   string
	TZ       string
	SSL      bool
	Host     string
	Port     int
	Cert     string
	Cors     string
	Key      string
	FronHost string
	Security Security
	Database Database
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

	godotenv.Load(".env")

	config.Mode = os.Getenv("API_MODE")
	config.Domain = os.Getenv("API_DOMAIN")

	config.TZ = os.Getenv("TZ")
	config.Port, _ = getenvInt("API_PORT")
	config.SSL, _ = getenvBool("API_SSL")
	config.Cert = os.Getenv("API_CERT")
	config.Key = os.Getenv("API_KEY")
	config.Host = os.Getenv("API_HOST")

	config.Cors = os.Getenv("API_CORS")

	config.Cert = os.Getenv("API_CERT")
	config.Key = os.Getenv("API_KEY")

	config.FronHost = os.Getenv("FRONT_HOST")

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

func getenvSliceString(key string) []string {
	s := os.Getenv(key)
	v := strings.Split(s, ",")
	return v
}

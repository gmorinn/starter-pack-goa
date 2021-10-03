package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kataras/golog"
)

// API represent api
type API struct {
	TZ       string
	SSL      bool
	Host     string
	Port     int
	Cert     string
	Key      string
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

// Get return config
func Get() *API {
	var config API
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/api_crud/.env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.TZ = os.Getenv("TZ")
	config.Port, _ = getenvInt("API_PORT")
	config.SSL, _ = getenvBool("API_SSL")
	config.Host = os.Getenv("API_HOST")

	config.Database.Host = os.Getenv("POSTGRES_HOST")
	config.Database.Database = os.Getenv("POSTGRES_DB")
	config.Database.User = os.Getenv("POSTGRES_USER")
	config.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Database.Port, _ = getenvInt("POSTGRES_PORT")

	return &config
}

func getenvInt(key string) (int, error) {
	s := os.Getenv(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		golog.Error("Env ", key, " : ", err.Error())
		return 0, err
	}
	return v, nil
}

func getenvBool(key string) (bool, error) {
	s := os.Getenv(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		golog.Error("Env ", key, " : ", err.Error())
		return false, err
	}
	return v, nil
}

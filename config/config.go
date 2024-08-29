package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// create struct for env
type DB struct {
	DBHOST     string `envconfig:"DB_HOST"`
	DBUSER     string `envconfig:"DB_USER"`
	DBPASSWORD string `envconfig:"DB_PASSWORD"`
	DBNAME     string `envconfig:"DB_NAME"`
	PORT       int    `envconfig:"PORT"`
}

// load env here
func LoadConfig() (DB, error) {
	var cfg DB
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	//fill env data to struct
	err = envconfig.Process("myapp", &cfg)
	if err != nil {
		log.Fatalf("process env failed, error: %v", err)
	}
	return cfg, nil
}

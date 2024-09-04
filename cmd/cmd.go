package cmd

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/pet-project/config"
	"github.com/quocbang/pet-project/repository"
)

func Run() {
	s := echo.New()

	// load env
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(env)
	// init logger

	// connect database
	client, err := repository.ConnectToDatabase(env)
	if err != nil {
		log.Fatal(err)
	}
	db, err := client.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect db done")

	// new router

	// listen http
	if err := s.Start(":8888"); err != nil {
		log.Fatalf("failed to start server, error: %v", err)
	}
}

package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Run() {
	s := echo.New()
	
	// load env

	// init logger

	// connect database

	// new router

	// listen http
	if err := s.Start(":8888"); err != nil {
		log.Fatalf("failed to start server, error: %v", err)
	}
}

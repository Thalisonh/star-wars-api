package main

import (
	"log"

	"github.com/Thalisonh/star-wars-api/server"
	"github.com/joho/godotenv"
)

func main() {
	errDotEnv := godotenv.Load()

	if errDotEnv != nil {
		log.Fatal("Error loading .env files")
	}

	s := server.Init()
	s.Run()
}

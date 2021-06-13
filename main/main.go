package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() bool {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading environment variables")
		return false
	}
	log.Println(os.Getenv("DB_PASSWORD"))
	return true
}

func main() {
	// loadEnv()
}

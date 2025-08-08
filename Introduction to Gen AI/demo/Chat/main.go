package main

import (
	//	"context"
	"fmt"
	"log"
	"os"
	// "github.com/openai/openai-go"
	// "github.com/openai/openai-go/option"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv("MODEL_RUNNER_BASE_URL")
	fmt.Println("MODEL_RUNNER_BASE_URL:", value)
}

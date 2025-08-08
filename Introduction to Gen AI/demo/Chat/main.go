package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	fmt.Print(os.Getenv("MODEL_RUNNER_BASE_URL"))
}

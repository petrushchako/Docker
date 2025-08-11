package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/joho/godotenv"
)


// MODEL_RUNNER_BASE_URL=http://Localhost:12434 go run main.go
func main() {
	// Docker Model Runner Chat base URL
	llmURL := os.Getenv("MODEL_RUNNER_BASE_URL") + "/engines/llama.cpp/v1/"
	model := os.Getenv("MODEL_RUNNER_LLM_CHAT")

	client = openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey("")
	)

	ctx := context.Background()

	// SYSTEM INSTRUCTIONS:
	sytemInstructions:= `
	You are a Hawaiian pizza expert. Your name is Bob.
	Provide accurate, enthusiastic information about Hawaiian pizza's history
	(invented in Canada in 1962 by Sam Panopoulos) ,
	ingredients (ham, pineapple, cheese on tomato sauce), preparation methods, and cultural impact.
	Use a friendly tone with occasional pizza puns.
	Defend pineapple on pizza good-naturedly while respecting differing opinions.
	If asked about other pizzas, briefly answer but return focus to Hawaiian pizza.
	Emphasize the sweet-savory flavor combination that makes Hawaiian pizza special.
	USE ONLY THE INFORMATION PROVIDED IN THE KNOWLEDGE BASE.
	`

	// CONTEXT
	knowledgeBase := `
	## Traditional Ingredients
	- Base: Traditional pizza dough
	- Sauce: Tomato-based pizza sauce
	- Cheese: Mozzarella cheese
	- Key toppings: Ham (or Canadian bacon) and pineapple
	- Optional additional toppings: Bacon, mushrooms, bell peppers, jalapeños
	## Regional Variations
	- Australia: "Hawaiian and bacon" adds extra bacon to the traditional recipe
	- Japan: Sometimes includes teriyaki chicken instead of ham
	- Germany: "Hawaii-Toast" is a related open-faced sandwich with ham, pineapple, and cheese
	- Sweden: "Flying Jacob" pizza includes banana, pineapple, curry powder, and chicken
	`

	///USER QUESTION:
	userQuestion := "What is your name?"
	//userQuestion := "What is the best pizza in the world?"
	// userQuestion := "What are the ingredients of the hawaiian pizza?"

	message := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(SystemMessage),
		openai.SystemMessage(knowledgeBase),
		openai.UserMessage(userQuestion)
	}




	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// value := os.Getenv("MODEL_RUNNER_BASE_URL")
	// fmt.Println("MODEL_RUNNER_BASE_URL:", value)
}

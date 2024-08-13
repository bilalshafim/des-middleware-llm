package handler

import (
	"fmt"
)

// Define the type for your intent handling functions
type IntentFunc func()

func greetIntent() {
	fmt.Println("Hello! How can I assist you today?")
}

func productLLMQueryIntent() {
	fmt.Println("Product intent matched!")
}

var intents = map[string]IntentFunc{
	"greet":             greetIntent,
	"product_llm_query": productLLMQueryIntent,
}

func HandleIntent(intent *string) {
	if handler, exists := intents[*intent]; exists {
		handler()
	} else {
		fmt.Println("Intent not recognized")
	}
}

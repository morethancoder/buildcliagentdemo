package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {
	url := "https://openrouter.ai/api/v1"
	apiKey := os.Getenv("OPENROUTER_API_KEY")

	if apiKey == "" {
		log.Fatal("Api key is required!")
	}

	client := openai.NewClient(
		option.WithBaseURL(url),
		option.WithAPIKey(apiKey),
	)

	messages := []openai.ChatCompletionMessageParamUnion{}

	messages = append(messages, openai.UserMessage("what model are you?"))

	model := "x-ai/grok-4-fast:free" 

	params := openai.ChatCompletionNewParams{
		Model: model,
		Messages: messages,
	}
	
	ctx := context.Background()

	res, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Choices[0].Message.Content)

}

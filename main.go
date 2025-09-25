package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	model := "x-ai/grok-4-fast:free"

	params := openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messages,
	}

	ctx := context.Background()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		params.Messages = append(params.Messages, openai.UserMessage(input))

		res, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(res.Choices[0].Message.Content)

	}

}

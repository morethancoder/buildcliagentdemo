package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
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

	fmt.Println(color.CyanString(
		fmt.Sprintf("CLI Agent powered by (%s)", model),
	))

	for {
		fmt.Print(color.GreenString("\n> "))

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if input == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		if input == "exit" {
			fmt.Println(color.WhiteString("see ya"))
			break
		}

		params.Messages = append(params.Messages, openai.UserMessage(input))

		res, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			log.Println(color.RedString(err.Error()))
		}

		output := res.Choices[0].Message.Content
		fmt.Println("\n", color.YellowString(output))

		params.Messages = append(params.Messages, openai.AssistantMessage(output))

		messagesCount := len(params.Messages)
		fmt.Println(color.WhiteString(
			fmt.Sprintf("\nmessages count: %d", messagesCount),
		))

		if messagesCount > 10 {
			params.Messages = params.Messages[messagesCount - 4:]
			//clean up the view
			fmt.Println(color.WhiteString("cleaned up chat history!"))
		} 

	}

}

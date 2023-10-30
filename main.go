package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	value := os.Getenv("OPENAI_API_KEY")
	if len(value) == 0 {
		slog.Error("No OpenAI API key found in env vars")
		os.Exit(11)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		slog.Info("Can't read given text")
	}

	client := openai.NewClient(value)
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: text,
				},
			},
		},
	)
	if err != nil {
		slog.Error("Could not get a response from OpenAI")
	}

	fmt.Println(response.Choices[0].Message.Content)
}

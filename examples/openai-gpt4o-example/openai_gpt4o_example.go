package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jaysteelmind/langchaingo/llms"
	"github.com/jaysteelmind/langchaingo/llms/openai"
	"github.com/jaysteelmind/langchaingo/llms/streaming"
)

func main() {
	llm, err := openai.New(openai.WithModel("gpt-4o"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a go programming expert"),
		llms.TextParts(llms.ChatMessageTypeHuman, "Why is go a good language for production LLM applications?"),
	}

	completion, err := llm.GenerateContent(
		ctx,
		content,
		llms.WithStreamingFunc(func(_ context.Context, chunk streaming.Chunk) error {
			fmt.Println(chunk.String())
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = completion
}

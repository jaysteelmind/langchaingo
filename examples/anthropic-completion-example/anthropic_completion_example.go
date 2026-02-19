package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jaysteelmind/langchaingo/llms"
	"github.com/jaysteelmind/langchaingo/llms/anthropic"
	"github.com/jaysteelmind/langchaingo/llms/streaming"
)

func main() {
	llm, err := anthropic.New(
		anthropic.WithModel("claude-3-5-sonnet-20240620"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, "Hi claude, write a poem about golang powered AI systems",
		llms.WithTemperature(0.8),
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

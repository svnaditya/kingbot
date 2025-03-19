package llm

import (
	"context"
	"log"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/ollama/ollama/api"
	"github.com/svnaditya/telegram-x-bot/config"
)

func GeneratePost(prompt string) string {
	switch config.Cfg.LLM {
	case "chatgpt":
		return callChatGPT(prompt)
	case "ollama":
		return callOllama(prompt)
	default:
		return "Unsupported LLM"
	}
}

func callChatGPT(prompt string) string {
	ctx := context.Background()
	client := gpt3.NewClient(config.Cfg.LLMToken)
	resp, err := client.CompletionWithEngine(ctx, gpt3.GPT3Dot5Turbo, gpt3.CompletionRequest{
		Prompt:      []string{config.Cfg.Preamble + prompt},
		MaxTokens:   gpt3.IntPtr(30),
		Temperature: gpt3.Float32Ptr(0),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Choices[0].Text
}

func callOllama(prompt string) string {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  config.Cfg.Model,
		Prompt: config.Cfg.Preamble + prompt,
		Stream: new(bool),
	}

	ctx := context.Background()
	var result string
	respFunc := func(resp api.GenerateResponse) error {
		result = resp.Response
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

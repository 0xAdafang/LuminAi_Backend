package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetEmbedding(text string) ([]float32, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: []string{text},
			Model: openai.SmallEmbedding3,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Data[0].Embedding, nil
}

func GenerateResponse(question string, contextText string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an expert assistant. Answer the question using ONLY the context provided below. If the answer is not in the context, say you don't know.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Contexte: %s\n\nQuestion: %s", contextText, question),
				},
			},
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

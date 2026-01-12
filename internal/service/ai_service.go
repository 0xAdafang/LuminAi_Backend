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
					Role: openai.ChatMessageRoleSystem,

					Content: "Tu es un assistant expert spécialisé dans l'analyse de documents. Réponds de manière précise à la question en utilisant les informations du contexte fourni. Si l'information est absente, explique poliment que tu n'as pas assez de données pour répondre.",
				},
				{
					Role: openai.ChatMessageRoleUser,

					Content: fmt.Sprintf("--- DEBUT DU CONTEXTE ---\n%s\n--- FIN DU CONTEXTE ---\n\nQuestion de l'utilisateur : %s", contextText, question),
				},
			},
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

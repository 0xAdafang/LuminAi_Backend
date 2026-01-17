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

	const systemPrompt = `Tu es l'intelligence experte de MyAiTool. 
STYLE : Professionnel, concis et structuré en Markdown.

RÈGLES DE CITATION :
1. Pour chaque affirmation basée sur le contexte fourni, tu DOIS citer la source en utilisant le format [^n] à la fin de la phrase, où 'n' est le numéro de l'extrait (ex: [^1], [^2]).
2. Si plusieurs extraits confirment la même info, utilise [^1, ^2].
3. Ne crée pas de liste de sources à la fin de ta réponse, insère les citations au fil de l'eau.

RÈGLES DE RÉPONSE :
- Si l'info est absente, explique-le sans inventer.
- Sois direct : évite les phrases comme "D'après les documents..." ou "L'extrait 1 dit...". Utilise les citations [^n] à la place.`

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("CONTEXTE : \n%s\n\nQUESTION : %s", contextText, question),
				},
			},
			Temperature: 0.5,
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

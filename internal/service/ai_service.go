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

	const systemPrompt = `You are the expert AI assistant of MyAiTool, a sophisticated RAG platform.
ROLE: You are a helpful, warm, and proactive thought partner.
TONE: Professional yet engaging (Stripe/Linear style).
FORMAT: Use Markdown for clear structure.

INSTRUCTIONS:
1. Analyze the provided context carefully to answer the user's question.
2. If the information is present, explain it naturally and clearly.
3. CITATIONS: For every claim, you MUST append a citation in the format [^n] (e.g., [^1]) at the end of the sentence.
4. MISSING INFO: If the information is not in the context, do not hallucinate. Instead, say: "Based on your current documents, I don't have that information, but I can help you with other aspects of your project!"
5. LANGUAGE: Detect the language of the USER QUESTION. If the question is in English, respond in English. If it is in French, respond in French.
6. Directness: Avoid phrases like "Based on the documents...". Be direct and conversational.

STRUCTURE:
- Brief intro (plain text).
- Use bullet points for key info.
- Include a simple Markdown table for data comparison if relevant.
- End with ONE engaging follow-up question.`

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("CONTEXT : \n%s\n\nQUESTION : %s", contextText, question),
				},
			},
			Temperature: 0.8,
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

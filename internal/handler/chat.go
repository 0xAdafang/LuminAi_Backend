package handler

import (
	"MyAiTool/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *IngestHandler) HandleChat(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(UserIDKey).(string)

	var req struct {
		Question string `json:"question"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	queryVector, err := service.GetEmbedding(req.Question)
	if err != nil {
		http.Error(w, "AI Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	articles, err := h.Repo.SearchSimilarArticles(queryVector, 3, userID)
	if err != nil {
		http.Error(w, "Database Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Number of items found :  %d\n", len(articles))

	var contextBuilder strings.Builder
	if len(articles) > 0 {
		for i, art := range articles {
			contextBuilder.WriteString(fmt.Sprintf("--- EXTRACT %d (Source: %s) ---\n%s\n\n", i+1, art.Title, art.Content))
		}
	} else {
		contextBuilder.WriteString("No relevant articles were found in the knowledge base.")
	}

	contextText := contextBuilder.String()

	limit := 100
	if len(contextText) < limit {
		limit = len(contextText)
	}

	fmt.Printf("DEBUG CONTEXT (%d chars) : %s\n", limit, contextText[:limit])

	answer, err := service.GenerateResponse(req.Question, contextText)
	if err != nil {
		http.Error(w, "AI Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"answer":   answer,
		"articles": articles,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Encoding Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

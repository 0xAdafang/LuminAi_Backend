package handler

import (
	"MyAiTool/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
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

	fmt.Printf("Nombre d'articles trouvés : %d\n", len(articles))

	contextText := "Aucun document trouvé dans la base de données."
	if len(articles) > 0 {
		fmt.Printf("Meilleur match : %s\n", articles[0].Title)

		contextText = articles[0].Content
	}

	limit := 100
	if len(contextText) < limit {
		limit = len(contextText)
	}

	fmt.Printf("DEBUG CONTEXT (%d chars) : %s\n", limit, contextText[:limit])

	answer, err := service.GenerateResponse(req.Question, contextText)

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

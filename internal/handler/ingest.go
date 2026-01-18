package handler

import (
	"MyAiTool/internal/models"
	"MyAiTool/internal/repository"
	"MyAiTool/internal/service"
	"encoding/json"
	"net/http"
)

type IngestHandler struct {
	Repo *repository.Repository
}

func (h *IngestHandler) HandleIngest(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(UserIDKey).(string)

	var req models.IngestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	title, content, err := service.ExtractTextFromURL(req.URL)
	if err != nil {
		http.Error(w, "Error during extractionn : "+err.Error(), http.StatusInternalServerError)
		return
	}

	embedding, err := service.GetEmbedding(content)
	if err != nil {
		http.Error(w, "AI Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	article := models.Article{
		Title:     title,
		Content:   content,
		URL:       req.URL,
		Summary:   content[:200] + "...",
		Embedding: embedding,
		UserID:    userID,
	}

	if err := h.Repo.SaveArticle(article); err != nil {
		http.Error(w, "Database saving Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Article successfully indexed ! ✨"})

}

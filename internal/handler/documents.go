package handler

import (
	"encoding/json"
	"net/http"
)

func (h *IngestHandler) HandleListDocuments(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(UserIDKey).(string)

	docs, err := h.Repo.GetAllDocuments(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)

}

func (h *IngestHandler) HandleDeleteDocument(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL du document manquante", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteDocument(url); err != nil {
		http.Error(w, "Erreur lors de la suppression : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Document supprimé avec succès"})
}

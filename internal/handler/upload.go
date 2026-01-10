package handler

import (
	"MyAiTool/internal/models"
	"MyAiTool/internal/service"
	"io"
	"net/http"
)

func (h *IngestHandler) HandleFileUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error Retrieving the File: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, _ := io.ReadAll(file)

	text, err := service.ExtractTextFromPDF(fileBytes)
	if err != nil {
		http.Error(w, "Error Extracting Text: "+err.Error(), http.StatusInternalServerError)
		return
	}

	embedding, err := service.GetEmbedding(text)

	article := models.Article{
		Title:     header.Filename,
		Content:   text,
		URL:       "file://" + header.Filename,
		Summary:   "Document PDF importé localement",
		Embedding: embedding,
	}

	h.Repo.SaveArticle(article)
	w.WriteHeader(http.StatusCreated)
}

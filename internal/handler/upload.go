package handler

import (
	"MyAiTool/internal/models"
	"MyAiTool/internal/service"
	"fmt"
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

	chunks := service.SplitText(text, 1000, 200)

	for i, chunk := range chunks {

		embedding, err := service.GetEmbedding(chunk)
		if err != nil {
			continue
		}

		article := models.Article{

			Title:     fmt.Sprintf("%s (Partie %d)", header.Filename, i+1),
			Content:   chunk,
			URL:       "file://" + header.Filename,
			Summary:   fmt.Sprintf("Segment %d du document %s", i+1, header.Filename),
			Embedding: embedding,
		}

		h.Repo.SaveArticle(article)
	}

	w.WriteHeader(http.StatusCreated)
}

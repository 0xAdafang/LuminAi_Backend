package repository

import (
	"MyAiTool/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveArticle(art models.Article) error {

	// Construire un littéral de vecteur PostgreSQL: '[0.1,0.2,...]'
	vecParts := make([]string, len(art.Embedding))
	for i, v := range art.Embedding {
		vecParts[i] = fmt.Sprintf("%f", v)
	}
	vecString := fmt.Sprintf("[%s]", strings.Join(vecParts, ","))

	query := `INSERT INTO articles (title, content, url, summary, embedding) 
              VALUES ($1, $2, $3, $4, $5::vector)`

	_, err := r.db.Exec(query, art.Title, art.Content, art.URL, art.Summary, vecString)
	return err
}

func (r *Repository) SearchSimilarArticles(queryEmbedding []float32, limit int) ([]models.Article, error) {
	vecParts := make([]string, len(queryEmbedding))
	for i, v := range queryEmbedding {
		vecParts[i] = fmt.Sprintf("%f", v)
	}
	vecString := fmt.Sprintf("[%s]", strings.Join(vecParts, ","))

	query := `SELECT id, title, summary, url FROM articles 
              ORDER BY embedding <=> $1::vector LIMIT $2`

	rows, err := r.db.Query(query, vecString, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Summary, &a.URL); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

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

func (r *Repository) SearchSimilarArticles(queryEmbedding []float32, limit int, userID string) ([]models.Article, error) {
	vecParts := make([]string, len(queryEmbedding))
	for i, v := range queryEmbedding {
		vecParts[i] = fmt.Sprintf("%f", v)
	}
	vecString := fmt.Sprintf("[%s]", strings.Join(vecParts, ","))

	query := `SELECT id, title, content, summary, url FROM articles 
              WHERE user_id = $3
              ORDER BY embedding <=> $1::vector LIMIT $2`

	rows, err := r.db.Query(query, vecString, limit, userID)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Summary, &a.URL); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func (r *Repository) GetAllDocuments() ([]string, error) {
	rows, err := r.db.Query("SELECT DISTINCT url FROM articles ORDER BY url")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		docs = append(docs, url)
	}
	return docs, nil
}

func (r *Repository) DeleteDocument(url string) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE url = $1", url)
	return err
}

func (r *Repository) CreateUser(email, passwordHash string) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err := r.db.Exec(query, email, passwordHash)
	return err
}

func (r *Repository) GetUserByEmail(email string) (int, string, string, error) {
	var id int
	var emailDB, hash string
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id, &emailDB, &hash)
	return id, emailDB, hash, err
}

package main

import (
	"MyAiTool/internal/handler"
	"log"
	"net/http"
	"os"

	"MyAiTool/internal/repository"
	"MyAiTool/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Panicln("Note: No .env file found, using system env")
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer database.Close()

	repo := repository.NewRepository(database)
	log.Printf("Repository initialized: %v", repo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := &handler.IngestHandler{Repo: repo}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/ingest", h.HandleIngest)
	r.Post("/api/chat", h.HandleChat)
	r.Post("/api/upload", h.HandleFileUpload)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK Backend - Ok DB connected"))
	})

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

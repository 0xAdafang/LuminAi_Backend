# 🚀 LuminAi Backend (Go)

High-performance RAG (Retrieval-Augmented Generation) engine built with Go. This service handles document ingestion, text extraction, and vector-based search.

## ✨ Features
- **URL & PDF Ingestion**: Automatic text extraction from web pages and documents.
- **Vector Embeddings**: Generates semantic vectors using OpenAI's `text-embedding-3-small` model.
- **Bilingual Brain**: Adaptive AI response system supporting English and French.
- **Secure Authentication**: JWT-based user context management.
- **Real-time Processing**: Fast response times optimized for RAG workflows.

## 🛠 Tech Stack
- **Language**: Go 1.22+
- **AI**: OpenAI API (GPT-4o mini)
- **Database**: PostgreSQL with `pgvector`
- **Tools**: Go-OpenAI SDK

## ⚙️ Setup
1. Define your environment variables in `.env`:
   ```env
   OPENAI_API_KEY=your_key
   DATABASE_URL=postgres://...
   JWT_SECRET=your_secret

go run cmd/main.go

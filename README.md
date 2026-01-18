# 🚀 Lumina Backend (Go)

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

### 2. Frontend README (Next.js)
À placer dans `my-ai-tool-ui/README.md`.

# 🌌 Lumina UI

A modern, high-end bilingual interface for document intelligence. Built with a focus on speed, aesthetics, and user experience.

## 📸 Screenshots

| Landing Page | Chat Interface |
| :---: | :---: |
| ![Landing](./img/page1.png) | ![Chat](./img/chat1.png) |

| Bilingual Support | Secure Auth |
| :---: | :---: |
| ![Language Switch](./img/chat2.png) | ![Login](./img/login.png) |

## ✨ Features
- **Bilingual Interface**: Seamlessly switch between English and French with global state management.
- **Intelligent Chat**: Markdown rendering with a custom citation system linking directly to document sources.
- **Modern UI**: Dark mode glassmorphism with an animated star-field background.
- **Responsive Layout**: Sidebar-driven document management with real-time sync.

## 🛠 Tech Stack
- **Framework**: Next.js 15 (App Router)
- **Styling**: Tailwind CSS + Typography plugin
- **Animations**: Framer Motion
- **Icons**: Lucide React
- **Notifications**: Sonner (Toast system)

## 🚀 Getting Started
1. Install dependencies:
   ```bash
   npm install
Set up your .env.local:

Extrait de code

NEXT_PUBLIC_API_URL=http://localhost:8080/api
Launch development server:

Bash

npm run dev

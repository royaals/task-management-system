# TaskAI - AI-Powered Task Management System


## 📋 Description

TaskAI is a modern, AI-powered task management system that combines intelligent task organization with real-time collaboration features. Built with cutting-edge technologies, it helps teams and individuals manage tasks more efficiently through AI-driven suggestions, real-time updates, and smart task breakdowns.

## ✨ Features

### Core Features
- 🔐 Secure JWT-based authentication
- ✅ Intuitive task creation and management
- 🤖 AI-powered task suggestions and breakdowns
- 🔄 Real-time updates via WebSocket
- 👥 Team collaboration and task assignment
- 📊 Analytics and progress tracking



## 🛠️ Tech Stack

### Backend
- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT
- **AI Integration:** OpenAI API

### Frontend
- **Framework:** Next.js 14 
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **UI Components:** Shadcn UI

### DevOps
- **Containerization:** Docker
- **Orchestration:** Kubernetes

## 📦 Installation

### Prerequisites
- Go 1.21+
- Node.js 18+
- MongoDB
- Docker (optional)
- OpenAI API Key

### Local Development Setup

1. **Clone the repository**
```bash
git clone https://github.com/royaals/task-system.git
cd task-system
```
2. backend
```bash
cd backend
replace the .env.example to .env
go mod tidy
docker run -d --name mongodb -p 27017:27017 mongo
go run cmd/api/main.go
```
3. frontend
```bash
cd frontend
npm install
npm run dev
```

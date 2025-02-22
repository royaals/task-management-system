# TaskAI - AI-Powered Task Management System


## 📋 Description

TaskAI is a modern, AI-powered task management system that combines intelligent task organization with real-time collaboration features. Built with cutting-edge technologies, it helps teams and individuals manage tasks more efficiently through AI-driven suggestions, real-time updates, and smart task breakdowns.

## ✨ Features

### Core Features
- 🔐 Secure JWT-based authentication
- ✅ Intuitive task creation and management
- 🤖 AI-powered task suggestions and breakdowns
- 👥 task assignment


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
- Docker 
- OpenAI API Key

## Local Development Setup

Follow the steps below to set up and run the Task System locally.

### Clone the Repository

```bash
git clone git clone https://github.com/royaals/task-system.git
cd task-management-system
```

---

## Backend Setup

1. Navigate to the backend directory:

```bash
cd backend
```

2. Copy the example environment file:

```bash
cp .env.example .env
```

3. Update `.env` with your configurations:
   - **MongoDB URI**
   - **JWT secret**
   - **OpenAI API key**

4. Install Go dependencies:

```bash
go mod tidy
```

5. Run MongoDB:

```bash
docker run -d --name mongodb -p 27017:27017 mongo
```

6. Start the backend server:

```bash
go run cmd/api/main.go
```

---

## Frontend Setup

1. Navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Create and update environment variables:

```bash
cp .env.example .env.local
```

4. Start the development server:

```bash
npm run dev
```

---

## Docker Setup

### Using Docker Compose

1. Build and start all services:

```bash
docker-compose up --build
```

2. Stop all services:

```bash
docker-compose down
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.


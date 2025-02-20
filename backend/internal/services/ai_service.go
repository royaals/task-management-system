package services

import (
    "context"
    "google.golang.org/api/option"
    "cloud.google.com/go/vertexai/genai"
    "task-management/internal/models"
)

type AIService struct {
    client *genai.Client
}

func NewAIService(apiKey string) (*AIService, error) {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        return nil, err
    }
    return &AIService{client: client}, nil
}

func (s *AIService) GenerateTaskSuggestions(task models.Task) (string, error) {
    ctx := context.Background()
    model := s.client.GenerativeModel("gemini-pro")

    prompt := generateAIPrompt(task)
    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", err
    }

    return resp.Candidates[0].Content.Parts[0].(genai.Text).Text, nil
}

func generateAIPrompt(task models.Task) string {
    return `Please analyze this task and provide suggestions:
    Title: ` + task.Title + `
    Description: ` + task.Description + `
    Current Status: ` + task.Status + `
    Priority: ` + task.Priority + `
    
    Please provide:
    1. Task breakdown suggestions
    2. Estimated time requirements
    3. Potential challenges and solutions
    4. Best practices and recommendations`
}
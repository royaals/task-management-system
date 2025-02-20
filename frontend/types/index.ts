export interface User {
    id: string;
    email: string;
  }
  

// types/index.ts
export interface Task {
    id: string;
    title: string;
    description: string;
    status: 'todo' | 'in_progress' | 'completed';
    priority: 'low' | 'medium' | 'high';
    due_date?: string;
    assigned_to?: string;
    created_by: string;
    created_at: string;
    updated_at: string;
}

  export interface AIResponse {
    suggestions: string;
    generated_at: string;
  }
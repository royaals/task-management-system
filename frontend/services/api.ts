import axios from 'axios';
import { Task } from '@/types';

const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'
});

// Request interceptor for adding auth token
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Response interceptor for handling errors
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export const taskApi = {
    getTasks: async (): Promise<Task[]> => {
        const response = await api.get('/tasks');
        return response.data;
    },

    createTask: async (task: Partial<Task>): Promise<Task> => {
        const response = await api.post('/tasks', task);
        return response.data;
    },

    updateTask: async (id: string, updates: Partial<Task>): Promise<Task> => {
        const response = await api.put(`/tasks/${id}`, updates);
        return response.data;
    },

    deleteTask: async (id: string): Promise<void> => {
        await api.delete(`/tasks/${id}`);
    },

     getAISuggestions: async (prompt: string) => {
        try {
            const response = await api.post('/ai/suggestions', {
                prompt: prompt
            });
            return response.data;
        } catch (error: any) {
            console.error('AI Service Error:', error.response?.data || error.message);
            throw error;
        }
   
  }
};

export const authApi = {
    login: async (email: string, password: string) => {
        const response = await api.post('/login', { email, password });
        return response.data;
    },

    register: async (email: string, password: string) => {
        const response = await api.post('/register', { email, password });
        return response.data;
    }
};
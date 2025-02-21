'use client';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { useRouter } from 'next/navigation';
import { toast } from 'react-hot-toast';
import axios from 'axios';

interface User {
    id: string;
    email: string;
}

interface AuthContextType {
    user: User | null;
    login: (email: string, password: string) => Promise<void>;
    logout: () => void;
    isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
        checkAuth();
    }, []);

    const checkAuth = async () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                setIsLoading(false);
                router.push('/login');
                return;
            }

            // Get user data using the token
            const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/tasks`, {
                headers: { Authorization: `Bearer ${token}` }
            });

            if (response.status === 200) {
                // If we can fetch tasks, the token is valid
                setUser({ id: 'user_id', email: 'user_email' }); // You can store these in the token claims
            } else {
                localStorage.removeItem('token');
                router.push('/login');
            }
        } catch (error) {
            console.error('Auth check failed:', error);
            localStorage.removeItem('token');
            router.push('/login');
        } finally {
            setIsLoading(false);
        }
    };

    const login = async (email: string, password: string) => {
        try {
            const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
                email,
                password
            });

            const { token, user: userData } = response.data;
            localStorage.setItem('token', token);
            setUser(userData);
            router.push('/dashboard');
            toast.success('Login successful!');
        } catch (error) {
            console.error('Login failed:', error);
            toast.error('Invalid credentials');
            throw error;
        }
    };

    const logout = () => {
        localStorage.removeItem('token');
        setUser(null);
        router.push('/login');
        toast.success('Logged out successfully');
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, isLoading }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
}
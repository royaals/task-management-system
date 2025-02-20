'use client';

import { useEffect, useRef, useState } from 'react';
import { toast } from 'react-hot-toast';

export const useWebSocket = (onMessage?: (data: any) => void) => {
    const ws = useRef<WebSocket | null>(null);
    const [isConnected, setIsConnected] = useState(false);
    const reconnectTimeout = useRef<NodeJS.Timeout>();
    const maxReconnectAttempts = 5;
    const reconnectAttempts = useRef(0);

    const connect = () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                console.log('No token found, skipping WebSocket connection');
                return;
            }

            // Close existing connection if any
            if (ws.current) {
                ws.current.close();
            }

            // Create new WebSocket connection with token
            const wsUrl = `${process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8080/api/ws'}?token=${token}`;
            console.log('Connecting to WebSocket...');
            
            ws.current = new WebSocket(wsUrl);

            ws.current.onopen = () => {
                console.log('WebSocket connected');
                setIsConnected(true);
                reconnectAttempts.current = 0;
            };

            ws.current.onmessage = (event) => {
                try {
                    const data = JSON.parse(event.data);
                    console.log('WebSocket message received:', data);
                    
                    if (onMessage) {
                        onMessage(data);
                    }

                    // Handle notifications
                    if (data.type) {
                        switch (data.type) {
                            case 'task_created':
                                toast.success('New task created!');
                                break;
                            case 'task_updated':
                                toast.success('Task updated!');
                                break;
                            case 'task_deleted':
                                toast.success('Task deleted!');
                                break;
                        }
                    }
                } catch (error) {
                    console.error('Error parsing WebSocket message:', error);
                }
            };

            ws.current.onclose = (event) => {
                console.log('WebSocket disconnected:', event);
                setIsConnected(false);

                // Attempt to reconnect if not at max attempts
                if (reconnectAttempts.current < maxReconnectAttempts) {
                    reconnectAttempts.current += 1;
                    console.log(`Attempting to reconnect (${reconnectAttempts.current}/${maxReconnectAttempts})`);
                    reconnectTimeout.current = setTimeout(connect, 3000);
                } else {
                    toast.error('Connection lost. Please refresh the page.');
                }
            };

            ws.current.onerror = (error) => {
                console.error('WebSocket error:', error);
            };

        } catch (error) {
            console.error('Error creating WebSocket connection:', error);
            setIsConnected(false);
        }
    };

    useEffect(() => {
        connect();

        return () => {
            if (reconnectTimeout.current) {
                clearTimeout(reconnectTimeout.current);
            }
            if (ws.current) {
                ws.current.close();
            }
        };
    }, []);

    const sendMessage = (type: string, data: any) => {
        if (!ws.current || ws.current.readyState !== WebSocket.OPEN) {
            console.warn('WebSocket is not connected');
            return;
        }

        try {
            ws.current.send(JSON.stringify({ type, data }));
        } catch (error) {
            console.error('Error sending message:', error);
        }
    };

    return {
        isConnected,
        sendMessage,
        reconnect: connect
    };
};
'use client';

import { useEffect, useRef } from 'react';
import { toast } from 'react-hot-toast';

export const useWebSocket = (onMessage?: (data: any) => void) => {
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) return;

    const connect = () => {
      ws.current = new WebSocket(`ws://localhost:8080/api/ws`);

      ws.current.onopen = () => {
        console.log('Connected to WebSocket');
        ws.current?.send(JSON.stringify({ type: 'auth', token }));
      };

      ws.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (onMessage) {
          onMessage(data);
        }

        // Handle notifications
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
      };

      ws.current.onclose = () => {
        console.log('WebSocket disconnected. Attempting to reconnect...');
        setTimeout(connect, 3000);
      };

      ws.current.onerror = (error) => {
        console.error('WebSocket error:', error);
      };
    };

    connect();

    return () => {
      ws.current?.close();
    };
  }, [onMessage]);

  return ws.current;
};
'use client';

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import TaskList from '@/components/TaskList';
import Sidebar from '@/components/SideBar';
import AIChat from '@/components/AIChat';
import { useWebSocket } from '@/hooks/useWebSocket';
import { Task } from '@/types';

export default function Dashboard() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const ws = useWebSocket();

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    // Implement task fetching
  };

  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="flex-1 overflow-y-auto bg-gradient-radial from-white to-gray-50 p-8">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
        >
          <h1 className="text-4xl font-bold text-gray-900 mb-8">
            Task Dashboard
          </h1>
          <TaskList tasks={tasks} />
        </motion.div>
        <AIChat />
      </main>
    </div>
  );
}
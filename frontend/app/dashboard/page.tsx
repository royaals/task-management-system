//@ts-nocheck
'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { taskApi } from '@/services/api';
import TaskList from '@/components/TaskList';
import Sidebar from '@/components/SideBar';
import AIChat from '@/components/AIChat';
import CreateTaskModal from '@/components/CreateTaskModal';
import { Task } from '@/types';
import { useAuth } from '@/contexts/AuthContext';
import { toast } from 'react-hot-toast';
import {
    ChartBarIcon,
    ClockIcon,
    CheckCircleIcon,
    PlusIcon,
    ChartPieIcon,
    ExclamationCircleIcon,
} from '@heroicons/react/24/outline';
import LoadingSpinner from '@/components/LoadingSpinner';

interface Statistics {
    total: number;
    completed: number;
    inProgress: number;
    todo: number;
}

const initialStatistics: Statistics = {
    total: 0,
    completed: 0,
    inProgress: 0,
    todo: 0,
};

const StatCard = ({ title, value, icon: Icon, color }: {
    title: string;
    value: number;
    icon: any;
    color: string;
}) => (
    <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        className="bg-white rounded-xl shadow-sm p-6 hover:shadow-md transition-shadow"
    >
        <div className="flex items-center justify-between">
            <div>
                <p className="text-sm text-gray-600">{title}</p>
                <p className="text-2xl font-semibold mt-2">{value}</p>
            </div>
            <div className={`p-3 rounded-full ${color}`}>
                <Icon className="h-6 w-6 text-white" />
            </div>
        </div>
    </motion.div>
);

export default function Dashboard() {
    
    const { user, isLoading: isAuthLoading } = useAuth();
    const [tasks, setTasks] = useState<Task[]>([]);
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [selectedFilter, setSelectedFilter] = useState('all');
    const [statistics, setStatistics] = useState<Statistics>(initialStatistics);

    // Fetch tasks when authenticated
    useEffect(() => {
        if (!isAuthLoading && user) {
            fetchTasks();
            const interval = setInterval(fetchTasks, 30000);
            return () => clearInterval(interval);
        }
    }, [isAuthLoading, user]);

    // Calculate statistics when tasks change
    useEffect(() => {
        if (!tasks) return;
        
        const stats = tasks.reduce(
            (acc, task) => {
                acc.total++;
                switch (task.status) {
                    case 'completed':
                        acc.completed++;
                        break;
                    case 'in_progress':
                        acc.inProgress++;
                        break;
                    case 'todo':
                        acc.todo++;
                        break;
                }
                return acc;
            },
            { ...initialStatistics }
        );
        setStatistics(stats);
    }, [tasks]);

    const fetchTasks = async () => {
        try {
            setIsLoading(true);
            const fetchedTasks = await taskApi.getTasks();
            setTasks(fetchedTasks || []);
        } catch (error) {
            console.error('Error fetching tasks:', error);
            toast.error('Failed to fetch tasks');
            setTasks([]);
        } finally {
            setIsLoading(false);
        }
    };

    const handleCreateTask = async (taskData: Partial<Task>) => {
        try {
            await taskApi.createTask(taskData);
            await fetchTasks();
            setIsCreateModalOpen(false);
            toast.success('Task created successfully');
        } catch (error) {
            console.error('Error creating task:', error);
            toast.error('Failed to create task');
        }
    };

    const handleUpdateTask = async (taskId: string, updates: Partial<Task>) => {
        try {
            await taskApi.updateTask(taskId, updates);
            await fetchTasks();
            toast.success('Task updated successfully');
        } catch (error) {
            console.error('Error updating task:', error);
            toast.error('Failed to update task');
        }
    };

    const handleDeleteTask = async (taskId: string) => {
        try {
            await taskApi.deleteTask(taskId);
            await fetchTasks();
            toast.success('Task deleted successfully');
        } catch (error) {
            console.error('Error deleting task:', error);
            toast.error('Failed to delete task');
        }
    };

    const filteredTasks = tasks?.filter(task => {
        if (selectedFilter === 'all') return true;
        return task.status === selectedFilter;
    }) || [];

    if (isAuthLoading) {
        return (
            <div className="flex justify-center items-center h-screen">
                <LoadingSpinner size="large" />
            </div>
        );
    }

    if (!user) {
        return null;
    }

    return (
        <div className="flex h-screen bg-gray-50">
            <Sidebar />
            <main className="flex-1 overflow-y-auto p-8">
                <motion.div
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    className="max-w-7xl mx-auto"
                >
                    <div className="flex justify-between items-center mb-8">
                        <div>
                            <h1 className="text-3xl font-bold text-gray-900">Task Dashboard</h1>
                            <p className="text-gray-600 mt-1">
                                Welcome back, {user.email}
                            </p>
                        </div>
                        <button
                            onClick={() => setIsCreateModalOpen(true)}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center space-x-2"
                        >
                            <PlusIcon className="h-5 w-5" />
                            <span>Create Task</span>
                        </button>
                    </div>

                    {/* Statistics Grid */}
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
                        <StatCard
                            title="Total Tasks"
                            value={statistics.total}
                            icon={ChartBarIcon}
                            color="bg-blue-500"
                        />
                        <StatCard
                            title="In Progress"
                            value={statistics.inProgress}
                            icon={ClockIcon}
                            color="bg-yellow-500"
                        />
                        <StatCard
                            title="Completed"
                            value={statistics.completed}
                            icon={CheckCircleIcon}
                            color="bg-green-500"
                        />
                        <StatCard
                            title="Todo"
                            value={statistics.todo}
                            icon={ChartPieIcon}
                            color="bg-red-500"
                        />
                    </div>

                    {/* Tasks Section */}
                    <div className="bg-white rounded-xl shadow-sm p-6">
                        <div className="flex justify-between items-center mb-6">
                            <h2 className="text-xl font-semibold text-gray-900">Tasks</h2>
                            <div className="flex space-x-2">
                                <select
                                    value={selectedFilter}
                                    onChange={(e) => setSelectedFilter(e.target.value)}
                                    className="rounded-lg border border-gray-300 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                                >
                                    <option value="all">All Tasks</option>
                                    <option value="todo">Todo</option>
                                    <option value="in_progress">In Progress</option>
                                    <option value="completed">Completed</option>
                                </select>
                            </div>
                        </div>

                        {isLoading ? (
                            <div className="flex justify-center items-center h-64">
                                <LoadingSpinner size="large" />
                            </div>
                        ) : filteredTasks.length === 0 ? (
                            <div className="text-center py-12">
                                <ExclamationCircleIcon className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                                <h3 className="text-lg font-medium text-gray-900 mb-2">
                                    No tasks found
                                </h3>
                                <p className="text-gray-600">
                                    {selectedFilter === 'all'
                                        ? 'Get started by creating your first task'
                                        : 'No tasks match the selected filter'}
                                </p>
                            </div>
                        ) : (
                            <TaskList
    tasks={filteredTasks}
    onUpdateTask={handleUpdateTask}
    onDeleteTask={handleDeleteTask}
/>
                        )}
                    </div>
                </motion.div>

                {/* AI Chat Component */}
                <AIChat />

                {/* Create Task Modal */}
                <AnimatePresence>
                    {isCreateModalOpen && (
                        <CreateTaskModal
                            isOpen={isCreateModalOpen}
                            onClose={() => setIsCreateModalOpen(false)}
                            onCreateTask={handleCreateTask}
                        />
                    )}
                </AnimatePresence>
            </main>
        </div>
    );
}
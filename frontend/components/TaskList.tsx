import { motion } from 'framer-motion';
import {Task} from '@/types';

export default function TaskList({ tasks }: { tasks: Task[] }) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {tasks.map((task, index) => (
        <motion.div
          key={task.id}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.3, delay: index * 0.1 }}
          className="bg-white rounded-xl shadow-sm p-6 hover:shadow-md transition-shadow"
        >
          <h3 className="text-lg font-semibold text-gray-900 mb-2">
            {task.title}
          </h3>
          <p className="text-gray-600 mb-4">{task.description}</p>
          <div className="flex items-center justify-between">
            <span className={`px-3 py-1 rounded-full text-sm ${
              task.status === 'completed' ? 'bg-green-100 text-green-800' :
              task.status === 'in_progress' ? 'bg-yellow-100 text-yellow-800' :
              'bg-gray-100 text-gray-800'
            }`}>
              {task.status}
            </span>
            <span className="text-sm text-gray-500">
              {new Date(task.due_date).toLocaleDateString()}
            </span>
          </div>
        </motion.div>
      ))}
    </div>
  );
}
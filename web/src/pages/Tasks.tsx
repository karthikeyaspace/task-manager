import { useState, useEffect } from "react";
import { Plus } from "lucide-react";
import { Loader2 } from "lucide-react";
import { TaskItem, TaskType } from "../components/Task";
import { createTask, deleteTask, getAllTasks, updateTask } from "../utils/api";

export const Tasks = () => {
  const [tasks, setTasks] = useState<TaskType[]>([]);
  const [loading, setLoading] = useState(true);
  const [isOpen, setIsOpen] = useState(false);
  const [newTask, setNewTask] = useState<TaskType>({
    taskid: "",
    title: "",
    description: "",
    priority: 1,
    completed: false,
  });

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const fetchedTasks = await getAllTasks();
      setTasks(fetchedTasks);
    } catch (error) {
      console.error("Failed to fetch tasks:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateTask = async (newTask: TaskType) => {
    try {
      const res = await createTask(newTask);
      if (res.success) {
        setIsOpen(false);
        setNewTask({
          taskid: "",
          title: "",
          description: "",
          priority: 1,
          completed: false,
        });
        const t = { ...newTask, taskid: res.taskid };
        if (tasks === null) {
          setTasks([t]);
        } else setTasks(tasks.concat(t));
      }
    } catch (error) {
      console.error("Failed to create task:", error);
    }
  };

  const handleUpdateTask = async (updatedTask: TaskType) => {
    try {
      await updateTask(updatedTask);
      fetchTasks();
    } catch (error) {
      console.error("Failed to update task:", error);
    }
  };

  const handleDeleteTask = async (id: string) => {
    try {
      await deleteTask(id);
      fetchTasks();
    } catch (error) {
      console.error("Failed to delete task:", error);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    handleCreateTask(newTask);
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <Loader2 className="animate-spin" size={32} />
      </div>
    );
  }

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-2xl font-bold mb-6">Task Manager</h1>
      <div className="mb-6">
        {!isOpen ? (
          <button
            onClick={() => setIsOpen(true)}
            className="flex items-center gap-2 px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
          >
            <Plus size={20} />
            Add New Task
          </button>
        ) : (
          <form
            onSubmit={handleSubmit}
            className="space-y-4 bg-white p-4 rounded-lg shadow-sm"
          >
            <div>
              <input
                type="text"
                placeholder="Task title"
                value={newTask.title}
                onChange={(e) =>
                  setNewTask({ ...newTask, title: e.target.value })
                }
                className="w-full p-2 border rounded"
                required
              />
            </div>
            <div>
              <textarea
                placeholder="Task description"
                value={newTask.description}
                onChange={(e) =>
                  setNewTask({ ...newTask, description: e.target.value })
                }
                className="w-full p-2 border rounded"
              />
            </div>
            <div className="flex gap-2">
              <button
                type="submit"
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                Create Task
              </button>
              <button
                type="button"
                onClick={() => setIsOpen(false)}
                className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
              >
                Cancel
              </button>
            </div>
          </form>
        )}
      </div>

      <div className="space-y-4">
        {tasks ?
          tasks
            .sort((a, b) => b.priority - a.priority)
            .map((task) => (
              <TaskItem
                key={task.taskid}
                task={task}
                onUpdate={handleUpdateTask}
                onDelete={handleDeleteTask}
              />
            ))
          : <p>No tasks</p>
          }
      </div>
    </div>
  );
};

export default Tasks;

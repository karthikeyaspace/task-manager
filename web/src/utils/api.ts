import { TaskType } from "../components/Task";

const API_URL = "http://localhost:8080";

const getAllTasks = async () => {
  const response = await fetch(`${API_URL}/tasks`);
  const data = await response.json();
  if (data.success) return data.tasks;
  else throw new Error("Failed to fetch tasks");
};

const createTask = async (task: TaskType) => {
  const response = await fetch(`${API_URL}/create-task`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title: task.title,
      description: task.description,
      completed: task.completed,
      priority: task.priority,
    }),
  });
  const data = await response.json();
  return data;
};

const updateTask = async (task: TaskType) => {
  const response = await fetch(`${API_URL}/update-task`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(task),
  });
  const data = await response.json();
  return data;
};

const deleteTask = async (id: string) => {
  const response = await fetch(`${API_URL}/delete-task?tid=${id}`, {
    method: "DELETE",
  });
  const data = await response.json();
  return data;
};

export { getAllTasks, createTask, updateTask, deleteTask };

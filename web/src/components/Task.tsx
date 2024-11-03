import { useState } from "react";
import {
  CheckCircle2,
  Circle,
  Edit,
  Trash2,
  Star,
  StarOff,
} from "lucide-react";

export interface TaskType {
  taskid: string;
  title: string;
  description: string;
  priority: number;
  completed: boolean;
}

interface TaskItemProps {
  task: TaskType;
  onUpdate: (task: TaskType) => void;
  onDelete: (id: string) => void;
}

export const TaskItem = ({ task, onUpdate, onDelete }: TaskItemProps) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editedTask, setEditedTask] = useState(task);

  const handleUpdate = () => {
    onUpdate(editedTask);
    setIsEditing(false);
  };

  const toggleComplete = () => {
    onUpdate({ ...task, completed: !task.completed });
  };

  const togglePriority = () => {
    onUpdate({ ...task, priority: task.priority === 1 ? 0 : 1 });
  };

  return (
    <div className="flex items-center p-4 bg-white rounded-lg shadow-sm mb-3 hover:shadow-md transition-shadow">
      {isEditing ? (
        <div className="flex-1 space-y-2">
          <input
            type="text"
            value={editedTask.title}
            onChange={(e) =>
              setEditedTask({ ...editedTask, title: e.target.value })
            }
            className="w-full p-2 border rounded"
          />
          <textarea
            value={editedTask.description}
            onChange={(e) =>
              setEditedTask({ ...editedTask, description: e.target.value })
            }
            className="w-full p-2 border rounded"
          />
          <div className="flex gap-2">
            <button
              onClick={handleUpdate}
              className="px-3 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              Save
            </button>
            <button
              onClick={() => setIsEditing(false)}
              className="px-3 py-1 bg-gray-300 rounded hover:bg-gray-400"
            >
              Cancel
            </button>
          </div>
        </div>
      ) : (
        <>
          <button onClick={toggleComplete} className="mr-3">
            {task.completed ? (
              <CheckCircle2 className="text-green-500" />
            ) : (
              <Circle className="text-gray-400" />
            )}
          </button>
          <div className="flex-1">
            <h3
              className={`font-medium ${
                task.completed ? "line-through text-gray-400" : ""
              }`}
            >
              {task.title}
            </h3>
            <p className="text-sm text-gray-600">{task.description}</p>
          </div>
          <div className="flex gap-2">
            <button onClick={togglePriority}>
              {task.priority === 1 ? (
                <Star className="text-yellow-500" />
              ) : (
                <StarOff className="text-gray-400" />
              )}
            </button>
            <button onClick={() => setIsEditing(true)}>
              <Edit className="text-blue-500" />
            </button>
            <button onClick={() => onDelete(task.taskid)}>
              <Trash2 className="text-red-500" />
            </button>
          </div>
        </>
      )}
    </div>
  );
};

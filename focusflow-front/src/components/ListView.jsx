import React, { useEffect, useState } from 'react';
import {
  getGoals
} from '../api/goals';
import {
  getTasksByGoalID,
  createTask,
  updateTask,
  deleteTask,
} from '../api/tasks';
import { getTags, createTag, deleteTag, updateTag } from '../api/tags';

export default function ListView({ user }) {
  const [goals, setGoals] = useState([]);
  const [selectedGoalId, setSelectedGoalId] = useState(null);
  const [tasks, setTasks] = useState([]);
  const [tags, setTags] = useState([]);

  const [editTaskId, setEditTaskId] = useState(null);
  const [editTaskTitle, setEditTaskTitle] = useState('');
  const [editTaskDescription, setEditTaskDescription] = useState('');
  const [editTaskTagIds, setEditTaskTagIds] = useState([]);

  const [showCreateForm, setShowCreateForm] = useState(false);
  const [newTaskTitle, setNewTaskTitle] = useState('');
  const [newTaskDescription, setNewTaskDescription] = useState('');
  const [newTaskTagIds, setNewTaskTagIds] = useState([]);

  const [newTagName, setNewTagName] = useState('');
  const [editingTags, setEditingTags] = useState({});
  const [tagError, setTagError] = useState('');

  const loadTags = async () => {
    const data = await getTags();
    setTags(data);
  };

  const loadGoals = async () => {
    const data = await getGoals();
    setGoals(data);
  };

  useEffect(() => {
    loadGoals();
    loadTags();
  }, []);

  useEffect(() => {
    if (selectedGoalId) {
      loadTasks(selectedGoalId);
    }
  }, [selectedGoalId]);

  const loadTasks = async (goalId) => {
    const result = await getTasksByGoalID(goalId);
    setTasks(result);
  };

  const handleCreateTag = async () => {
    if (!newTagName.trim()) return;
    try {
      await createTag({ name: newTagName });
      setNewTagName('');
      setTagError('');
      loadTags();
    } catch (err) {
      setTagError(err.message);
    }
  };

  const handleDeleteTag = async (tagId) => {
    if (window.confirm('Удалить тег?')) {
      try {
        await deleteTag(tagId);
        loadTags();
      } catch (err) {
        alert('Ошибка удаления тега');
      }
    }
  };

  const handleUpdateTag = async (id) => {
    const newName = editingTags[id];
    if (!newName?.trim()) return;
    try {
      await updateTag(id, { name: newName });
      setEditingTags(prev => {
        const updated = { ...prev };
        delete updated[id];
        return updated;
      });
      loadTags();
    } catch (err) {
      alert('Ошибка обновления тега');
    }
  };

  const handleMarkDone = async (taskId, isDone) => {
    const task = tasks.find(t => t.ID === taskId);
    if (!task) return;

    await updateTask(taskId, {
      title: task.title,
      description: task.description,
      goal_id: task.goal_id,
      user_id: task.user_id,
      tags: task.tags,
      is_done: isDone,
    });

    loadTasks(selectedGoalId);
  };

  const handleDelete = async (taskId) => {
    if (window.confirm('Удалить задачу?')) {
      await deleteTask(taskId);
      loadTasks(selectedGoalId);
    }
  };

  const handleEditClick = (task) => {
    setEditTaskId(task.ID);
    setEditTaskTitle(task.title);
    setEditTaskDescription(task.description);
    setEditTaskTagIds(task.tags?.map(t => String(t.id)) || []);
    setShowCreateForm(false);
  };

  const handleUpdate = async () => {
    await updateTask(editTaskId, {
      title: editTaskTitle,
      description: editTaskDescription,
      goal_id: selectedGoalId,
      tags: editTaskTagIds.map(id => ({ id: parseInt(id) })),
      is_done: false
    });

    setEditTaskId(null);
    setEditTaskTitle('');
    setEditTaskDescription('');
    setEditTaskTagIds([]);
    loadTasks(selectedGoalId);
  };

  const handleCreate = async () => {
    if (!selectedGoalId) {
      alert('Сначала выберите цель!');
      return;
    }

    if (!newTaskTitle || newTaskTagIds.length === 0) {
      alert('Введите название и выберите хотя бы один тег');
      return;
    }

    await createTask({
      title: newTaskTitle,
      description: newTaskDescription,
      goal_id: selectedGoalId,
      tags: newTaskTagIds.map(id => ({ id: parseInt(id) }))
    });

    setShowCreateForm(false);
    setNewTaskTitle('');
    setNewTaskDescription('');
    setNewTaskTagIds([]);
    loadTasks(selectedGoalId);
  };

  const renderTagManager = () => (
    <div className="col-md-3 border-start">
      <h5>Теги</h5>
      <ul className="list-group">
        {tags.map(tag => (
          <li key={tag.ID} className="list-group-item d-flex justify-content-between align-items-center">
            {editingTags[tag.ID] !== undefined ? (
              <>
                <input
                  type="text"
                  className="form-control form-control-sm me-2"
                  value={editingTags[tag.ID]}
                  onChange={(e) => setEditingTags(prev => ({ ...prev, [tag.ID]: e.target.value }))}
                />
                <button className="btn btn-sm btn-success me-1" onClick={() => handleUpdateTag(tag.ID)}>💾</button>
                <button className="btn btn-sm btn-outline-secondary" onClick={() => setEditingTags(prev => {
                  const updated = { ...prev };
                  delete updated[tag.ID];
                  return updated;
                })}>✖</button>
              </>
            ) : (
              <>
                {tag.name}
                <span>
                  <button className="btn btn-sm btn-outline-secondary me-1" onClick={() =>
                    setEditingTags(prev => ({ ...prev, [tag.ID]: tag.name }))
                  }>✏</button>
                  <button className="btn btn-sm btn-outline-danger" onClick={() => handleDeleteTag(tag.ID)}>🗑</button>
                </span>
              </>
            )}
          </li>
        ))}
      </ul>
      <div className="input-group mt-3">
        <input
          className="form-control"
          placeholder="Новый тег"
          value={newTagName}
          onChange={(e) => setNewTagName(e.target.value)}
        />
        <button className="btn btn-outline-primary" onClick={handleCreateTag}>➕</button>
      </div>
      {tagError && <div className="text-danger mt-2 small">{tagError}</div>}
    </div>
  );

  if (goals.length === 0) {
    return (
      <div className="row">
        <div className="col-md-9">
          <p className="text-muted p-3">Чтобы поставить задачу, пожалуйста, создайте цель в Календаре.</p>
        </div>
        {user?.role === 'superadmin' && renderTagManager()}
      </div>
    );
  }

  return (
    <div className="row">
      <div className="col-md-3 border-end">
        <h5>Цели</h5>
        <ul className="list-group">
          {goals.map(goal => (
            <li
              key={goal.id}
              className={`list-group-item ${selectedGoalId === goal.id ? 'active' : ''}`}
              style={{ cursor: 'pointer' }}
              onClick={() => setSelectedGoalId(goal.id)}
            >
              {goal.title}
            </li>
          ))}
        </ul>
      </div>

      <div className="col-md-6">
        <h5>Задачи</h5>
        {tasks
          .sort((a, b) => a.is_done - b.is_done)
          .map(task => (
            <div key={task.ID} className="border rounded p-2 mb-2 d-flex justify-content-between align-items-center">
              <div>
                <h6 className="mb-1">{task.title}</h6>
                <p className="text-muted small mb-0">{task.description}</p>
              </div>
              <div className="ms-2">
                {!task.is_done && (
                  <button className="btn btn-sm btn-success me-1" onClick={() => handleMarkDone(task.ID, true)}>Done</button>
                )}
                <button className="btn btn-sm btn-outline-primary me-1" onClick={() => handleEditClick(task)}>Edit</button>
                <button className="btn btn-sm btn-outline-danger" onClick={() => handleDelete(task.ID)}>Delete</button>
              </div>
            </div>
          ))}

        {editTaskId && (
          <div className="card p-3 mt-3">
            <h6 className="mb-3">Редактировать задачу</h6>
            <input
              type="text"
              className="form-control mb-2"
              placeholder="Название задачи"
              value={editTaskTitle}
              onChange={(e) => setEditTaskTitle(e.target.value)}
            />
            <textarea
              className="form-control mb-2"
              placeholder="Описание (необязательно)"
              value={editTaskDescription}
              onChange={(e) => setEditTaskDescription(e.target.value)}
              rows={2}
            />
            <select
              multiple
              className="form-select mb-3"
              value={editTaskTagIds}
              onChange={(e) => setEditTaskTagIds(Array.from(e.target.selectedOptions, o => o.value))}
            >
              {tags.map(tag => (
                <option key={tag.ID} value={tag.ID}>{tag.name}</option>
              ))}
            </select>
            <button className="btn btn-warning w-100" onClick={handleUpdate}>💾 Сохранить изменения</button>
          </div>
        )}

        <button
          className="btn btn-primary w-100 mt-3"
          disabled={!selectedGoalId}
          onClick={() => setShowCreateForm(prev => !prev)}
        >
          {showCreateForm ? 'Отмена' : 'Создать задачу'}
        </button>

        {showCreateForm && (
          <div className="card p-3 mt-3">
            <h6 className="mb-3">Новая задача</h6>
            <input
              type="text"
              className="form-control mb-2"
              placeholder="Название задачи"
              value={newTaskTitle}
              onChange={(e) => setNewTaskTitle(e.target.value)}
            />
            <textarea
              className="form-control mb-2"
              placeholder="Описание (необязательно)"
              value={newTaskDescription}
              onChange={(e) => setNewTaskDescription(e.target.value)}
              rows={2}
            />
            <select
              multiple
              className="form-select mb-3"
              value={newTaskTagIds}
              onChange={(e) => setNewTaskTagIds(Array.from(e.target.selectedOptions, o => o.value))}
            >
              {tags.map(tag => (
                <option key={tag.ID} value={tag.ID}>{tag.name}</option>
              ))}
            </select>
            <button className="btn btn-success w-100" onClick={handleCreate}>✅ Добавить задачу</button>
          </div>
        )}
      </div>

      {user?.role === 'superadmin' && renderTagManager()}
    </div>
  );
}

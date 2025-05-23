import React, { useEffect, useState } from 'react';
import { getTeamTasks } from '../api/tasks';
import {
  DndContext,
  closestCenter,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core';
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

function SortableItem({ task }) {
  const {
    attributes, listeners, setNodeRef, transform, transition,
  } = useSortable({ id: task.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    padding: '8px 12px',
    border: '1px solid #ddd',
    borderRadius: '5px',
    backgroundColor: '#fff',
    marginBottom: '8px',
    cursor: 'grab',
  };

  return (
    <div ref={setNodeRef} style={style} {...attributes} {...listeners}>
      <div><strong>{task.title}</strong></div>
      <div className="text-muted small">{task.tags.map(t => t.name).join(', ')}</div>
      {task.user && (
        <div className="small text-secondary">👤 {task.user.username}</div>
      )}
    </div>
  );
}

export default function KanbanBoard() {
  const [tasks, setTasks] = useState([]);
  const [selectedTag, setSelectedTag] = useState('all');
  const [tagOptions, setTagOptions] = useState([]);

  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, { coordinateGetter: sortableKeyboardCoordinates })
  );

  useEffect(() => {
    getTeamTasks().then((data) => {
      setTasks(data);

      const tags = new Set();
      data.forEach(task => task.tags.forEach(tag => tags.add(tag.name)));
      setTagOptions(['all', ...Array.from(tags)]);
    });
  }, []);

  const columns = {
    todo: tasks.filter(t => !t.is_done),
    done: tasks.filter(t => t.is_done),
  };

  const handleDragEnd = (event) => {
    // опционально — можно реализовать перемещение между статусами
  };

  return (
    <>
      <div className="mb-3">
        <label className="form-label fw-bold">Фильтр по тегу:</label>
        <select
          className="form-select"
          value={selectedTag}
          onChange={(e) => setSelectedTag(e.target.value)}
        >
          {tagOptions.map(tag => (
            <option key={tag} value={tag}>
              {tag === 'all' ? 'Все теги' : tag}
            </option>
          ))}
        </select>
      </div>

      <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
        <div className="d-flex gap-4">
          {Object.entries(columns).map(([status, items]) => {
            const filtered = selectedTag === 'all'
              ? items
              : items.filter(t => t.tags.some(tag => tag.name === selectedTag));

            return (
              <div key={status} style={{ width: '45%' }}>
                <h5 className="text-center">{status === 'todo' ? '🕒 В процессе' : '✅ Завершено'}</h5>
                <SortableContext items={filtered.map(t => t.id)} strategy={verticalListSortingStrategy}>
                  <div className="bg-light p-3 rounded" style={{ minHeight: '200px' }}>
                    {filtered.map(task => (
                      <SortableItem key={task.id} task={task} />
                    ))}
                  </div>
                </SortableContext>
              </div>
            );
          })}
        </div>
      </DndContext>
    </>
  );
}

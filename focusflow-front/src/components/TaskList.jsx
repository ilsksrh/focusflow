import React, { useState } from 'react';
import EditGoalModal from './EditGoalModal';
import { deleteGoal, updateGoal, getGoals } from '../api/goals';

export default function TaskList({ goals, selectedDate, setGoals }) {
  const [showEditModal, setShowEditModal] = useState(false);
  const [goalToEdit, setGoalToEdit] = useState(null);
  

  const dateStr = `${selectedDate.getFullYear()}-${(selectedDate.getMonth() + 1)
    .toString()
    .padStart(2, '0')}-${selectedDate.getDate().toString().padStart(2, '0')}`;

  const matchingGoals = goals.filter(goal => goal.date === dateStr);

  const handleDelete = async (goalId) => {
    if (window.confirm('Удалить эту цель?')) {
      try {
        await deleteGoal(goalId);
        const updatedGoals = await getGoals(); 
        setGoals(updatedGoals); 
      } catch (err) {
        console.error('Ошибка удаления цели:', err);
      }
    }
  };
  

  const handleEditClick = (goal) => {
    setGoalToEdit(goal);
    setShowEditModal(true);
  };

  const handleSaveEdit = async (updatedGoal) => {
    try {
      const saved = await updateGoal(updatedGoal.id, {
        title: updatedGoal.title,
        description: updatedGoal.description,
        date: updatedGoal.date,
      });
      setGoals(prev =>
        prev.map(g => g.id === saved.id ? saved : g)
      );
      setShowEditModal(false);
    } catch (err) {
      console.error('Ошибка при сохранении:', err);
    }
  };

  return (
    <div>
      <h5>Goals for {dateStr}</h5>

      {matchingGoals.length === 0 && <p>No goals for this date.</p>}

      {matchingGoals.map(goal => (
        <div key={goal.id} className="mb-3 p-3 border rounded shadow-sm">
          <h6>🎯 {goal.title}</h6>
          <p className="text-muted">{goal.description}</p>

          <div className="d-flex gap-2">
            <button
              className="btn btn-sm btn-warning"
              onClick={() => handleEditClick(goal)}
            >
              Edit
            </button>
            <button
  className="btn btn-sm btn-danger"
  
  onClick={() => handleDelete(goal.id)} 
> Delete
</button>

          </div>
        </div>
      ))}

      {showEditModal && (
        <EditGoalModal
          show={showEditModal}
          onClose={() => setShowEditModal(false)}
          onSave={handleSaveEdit}
          goal={goalToEdit}
        />
      )}
    </div>
  );
}

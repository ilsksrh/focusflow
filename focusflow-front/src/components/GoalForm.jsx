import React, { useState } from 'react';
import { createGoal } from '../api/goals';
import { format } from 'date-fns';

export default function GoalForm({ onCreate, selectedDate }) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const dateStr = format(selectedDate, 'yyyy-MM-dd');

    const newGoal = {
      title,
      description,
      date: dateStr,
    };

    await onCreate(newGoal);

    setTitle('');
    setDescription('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-2">
        <input
          type="text"
          className="form-control"
          placeholder="Goal Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </div>
      <div className="mb-2">
        <textarea
          className="form-control"
          placeholder="Goal Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
      </div>
      <button type="submit" className="btn btn-primary btn-sm">Save Goal</button>
    </form>
  );
}

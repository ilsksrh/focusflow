export async function getGoals() {
    const response = await fetch('http://localhost:8080/goals', {
      method: 'GET',
      credentials: 'include',
    });
  
    if (!response.ok) {
      throw new Error('Ошибка загрузки целей');
    }
  
    return await response.json();
  }
  
  export async function createGoal(goal) {
    const response = await fetch('http://localhost:8080/goals', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(goal),
    });
  
    if (!response.ok) {
      throw new Error('Ошибка создания цели');
    }
  
    return await response.json(); 
  }
  
  export async function deleteGoal(id) {
    const res = await fetch(`http://localhost:8080/goals/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
    if (!res.ok) {
      throw new Error('Ошибка удаления цели');
    }
  }
  
  export async function updateGoal(id, goalData) {
    const res = await fetch(`http://localhost:8080/goals/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(goalData),
    });
    if (!res.ok) {
      throw new Error('Ошибка обновления цели');
    }
    const data = await res.json();
    return data;
  }
  
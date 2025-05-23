export async function getTasks() {
    const res = await fetch('http://localhost:8080/tasks', {
      credentials: 'include',
    });
    return await res.json();
  }
  
  export async function createTask(task) {
    const res = await fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(task),
    });
    return await res.json();
  }
  
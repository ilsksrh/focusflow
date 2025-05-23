export async function getTags() {
    const res = await fetch('http://localhost:8080/tags', {
      credentials: 'include',
    });
    return await res.json();
  }
  
  export async function createTag(tag) {
    const res = await fetch('http://localhost:8080/tags', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify(tag),
    });
    return await res.json();
  }
  
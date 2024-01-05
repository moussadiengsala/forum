

export async function getUserInfo(data:any) {   
    let response = await fetch('http://localhost:8000/api/auth/signin', {
      method: 'POST',
      body: JSON.stringify(data),
    })

    if (!response.ok) {
        const errorMessage = await response.json();
        throw new Error(errorMessage.message || "An error occurred during signin");
    }

    return response.json();
}
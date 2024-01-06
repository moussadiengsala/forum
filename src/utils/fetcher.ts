
type FetcherProps = {
  data: any,
  endpoint: string,
  method: string,
}

export async function fetcher({data, endpoint, method} : FetcherProps) {   
    let response = await fetch(`http://localhost:8000/api${endpoint}`, {
      method: method,
      body: JSON.stringify(data),
    })

    if (!response.ok) {
        const errorMessage = await response.json();
        throw new Error(errorMessage.message || "An error occurred during signin");
    }

    return response.json();
}
export let apiUrl = import.meta.env.VITE_API_URL

if (!apiUrl) {
    apiUrl = "http://localhost:8080"
}
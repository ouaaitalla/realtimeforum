import { apiFetch } from "../utils/fetch.js";

export function registerRequest(user) {
    return apiFetch("/register", {
        method: "POST",
        body: JSON.stringify(user),
    });
}

export function loginRequest(credentials) {
    return apiFetch("/login", {
        method: "POST",
        body: JSON.stringify(credentials),
    });
}

export function meRequest() {
    return apiFetch("/me", {
        method: "GET",
    });
}

export async function logoutRequest() {
    const response = await fetch("http://localhost:8080/logout", {
        method: "POST",
        credentials: "include",
    });

    return await response.json();
}


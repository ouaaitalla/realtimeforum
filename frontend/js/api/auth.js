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

export function logoutRequest() {
    return apiFetch("/logout", {
        method: "POST",
    });
}
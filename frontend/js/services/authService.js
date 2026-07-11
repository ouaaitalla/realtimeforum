import { registerRequest, loginRequest } from "../api/auth.js";
import { logoutRequest } from "../api/auth.js";
import { meRequest } from "../api/auth.js";


export async function register(user) {
    return (await registerRequest(user)).data;
}

export async function login(credentials) {
    return (await loginRequest(credentials)).data;
}

export async function logout() {
    await logoutRequest();
}

export async function checkAuth() {
    try {
        await meRequest();
        return true;
    } catch (error) {
        return false;
    }
}

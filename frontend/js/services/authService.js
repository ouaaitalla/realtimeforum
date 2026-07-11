import { registerRequest, loginRequest } from "../api/auth.js";
import { logoutRequest } from "../api/auth.js";

export async function register(user) {
    const response = await registerRequest(user);

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export async function login(credentials) {
    const response = await loginRequest(credentials);

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

import { meRequest } from "../api/auth.js";


export async function checkAuth() {
    try {
        await meRequest();
        return true;
    } catch (error) {
        return false;
    }
}



export async function logout() {
    const response = await logoutRequest();

    if (!response.success) {
        throw new Error(response.message);
    }
}

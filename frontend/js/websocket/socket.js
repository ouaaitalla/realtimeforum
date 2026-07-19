import { API_BASE_URL } from "../utils/constants.js";

let socket = null;

const listeners = new Map();


export function connectWebSocket() {
    if (socket && socket.readyState === WebSocket.OPEN) {
        return socket;
    }

    const wsURL = API_BASE_URL
        .replace("http://", "ws://")
        .replace("https://", "wss://");

    socket = new WebSocket(`${wsURL}/ws`);

    socket.onopen = () => {
        console.log("WebSocket connected");
    };

    socket.onmessage = (event) => {
        const data = JSON.parse(event.data);

        const callbacks = listeners.get(data.type);

        if (!callbacks) return;

        callbacks.forEach(callback => callback(data.payload));
    };

    socket.onclose = () => {
        console.log("WebSocket disconnected");

        setTimeout(() => {
            connectWebSocket();
        }, 2000);
    };

    socket.onerror = (err) => {
        console.error("WebSocket Error:", err);
    };

    return socket;
}

export function disconnectWebSocket() {
    if (socket) {
        socket.close();
        socket = null;
    }
}

export function sendEvent(type, payload) {
    if (!socket || socket.readyState !== WebSocket.OPEN) {
        return;
    }

    socket.send(JSON.stringify({
        type,
        payload,
    }));
}

export function on(type, callback) {
    if (!listeners.has(type)) {
        listeners.set(type, []);
    }

    listeners.get(type).push(callback);
}

export function off(type, callback) {
    if (!listeners.has(type)) {
        return;
    }

    const callbacks = listeners.get(type);

    const index = callbacks.indexOf(callback);

    if (index !== -1) {
        callbacks.splice(index, 1);
    }
}

export function getSocket() {
    return socket;
}

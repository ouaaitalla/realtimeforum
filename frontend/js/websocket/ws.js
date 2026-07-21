import { API_BASE_URL } from "../utils/constants.js";

let socket = null;
let reconnect = true;

const listeners = new Map();

export const ws = {

    init() {
        this.connect();
    },

    connect() {

        reconnect = true;

        if (
            socket &&
            (
                socket.readyState === WebSocket.OPEN ||
                socket.readyState === WebSocket.CONNECTING
            )
        ) {
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

            callbacks.forEach(callback => {
                callback(data.payload);
            });

        };

        socket.onclose = () => {

            console.log("WebSocket disconnected");

            socket = null;

            if (!reconnect) {
                return;
            }

            setTimeout(() => {
                this.connect();
            }, 2000);

        };

        socket.onerror = (err) => {
            console.error("WebSocket error:", err);
        };

        return socket;
    },

    disconnect() {

        reconnect = false;

        if (socket) {
            socket.close();
            socket = null;
        }

    },

    send(type, payload) {

        if (
            !socket ||
            socket.readyState !== WebSocket.OPEN
        ) {
            return false;
        }

        socket.send(
            JSON.stringify({
                type,
                payload,
            })
        );

        return true;
    },

    sendMessage(receiverID, content) {

        return this.send("message", {
            receiver_id: receiverID,
            content,
        });

    },

    sendTyping(receiverID, isTyping) {

        return this.send("typing", {
            receiver_id: receiverID,
            is_typing: isTyping,
        });

    },

    on(type, callback) {

        if (!listeners.has(type)) {
            listeners.set(type, []);
        }

        listeners.get(type).push(callback);

    },

    off(type, callback) {

        if (!listeners.has(type)) {
            return;
        }

        const callbacks = listeners.get(type);

        const index = callbacks.indexOf(callback);

        if (index !== -1) {
            callbacks.splice(index, 1);
        }

    },

    clearListeners() {
        listeners.clear();
    },

    getSocket() {
        return socket;
    }

};

export default ws;
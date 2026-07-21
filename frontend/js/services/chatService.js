import { apiFetch } from "../utils/fetch.js";
import ws from "../websocket/ws.js";

export async function getUsers() {
    const response = await apiFetch("/users");

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export async function getConversation(userID, limit = 10, offset = 0) {
    const response = await apiFetch(`/messages/${userID}?limit=${limit}&offset=${offset}`);

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export function sendMessage(receiverID, content) {
    ws.send("message", {
        receiver_id: receiverID,
        content,
    });
}

export function sendTyping(receiverID, isTyping) {
    ws.send("typing", {
        receiver_id: receiverID,
        is_typing: isTyping,
    });
}


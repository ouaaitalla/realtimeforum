import { apiFetch } from "../utils/fetch.js";
import { sendEvent } from "../websocket/socket.js";

export async function getUsers() {
    const response = await apiFetch("/users");

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export async function getConversation(userID) {
    const response = await apiFetch(`/messages/${userID}`);

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export async function markConversationAsRead(userID) {
    const response = await apiFetch(`/messages/read/${userID}`, {
        method: "POST",
    });

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

export function sendMessage(receiverID, content) {
    sendEvent("message", {
        receiver_id: receiverID,
        content,
    });
}

export function sendTyping(receiverID, isTyping) {
    sendEvent("typing", {
        receiver_id: receiverID,
        is_typing: isTyping,
    });
}


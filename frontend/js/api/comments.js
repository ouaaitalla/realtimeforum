import { apiFetch } from "../utils/fetch.js";

export function getCommentsRequest(postId) {
    return apiFetch(`/posts/${postId}/comments`, {
        method: "GET",
    });
}

export function createCommentRequest(postId, comment) {
    return apiFetch(`/posts/${postId}/comments`, {
        method: "POST",
        body: JSON.stringify({
            content: comment,
        }),
    });
}


import { apiFetch } from "../utils/fetch.js";

export function getCommentsRequest(postId, limit = 10, offset = 0) {
    return apiFetch(`/posts/${postId}/comments?limit=${limit}&offset=${offset}`, {
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


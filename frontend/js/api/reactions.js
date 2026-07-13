import { apiFetch } from "../utils/fetch.js";

export function togglePostReactionRequest(postId, reaction) {
    return apiFetch(`/posts/${postId}/reaction`, {
        method: "POST",
        body: JSON.stringify({
            reaction,
        }),
    });
}

export function toggleCommentReactionRequest(commentId, reaction) {
    return apiFetch(`/comments/${commentId}/reaction`, {
        method: "POST",
        body: JSON.stringify({
            reaction,
        }),
    });
}

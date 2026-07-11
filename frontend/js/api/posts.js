
import { apiFetch } from "../utils/fetch.js";

export function createPostRequest(post) {
    return apiFetch("/posts", {
        method: "POST",
        body: JSON.stringify(post),
    });
}

export function getPostsRequest() {
    return apiFetch("/posts", {
        method: "GET",
    });
}

export function getPostRequest(id) {
    return apiFetch(`/posts/${id}`, {
        method: "GET",
    });
}



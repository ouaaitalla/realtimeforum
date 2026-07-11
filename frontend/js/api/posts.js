
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



import { apiFetch } from "../utils/fetch.js";

export function createPostRequest(post) {
    return apiFetch("/posts", {
        method: "POST",
        body: JSON.stringify(post),
    });
}


export function getPostsRequest(filters = {}) {

    const params = new URLSearchParams();

    if (filters.category) {
        params.append("category", filters.category);
    }

    if (filters.mine) {
        params.append("mine", "true");
    }

    if (filters.liked) {
        params.append("liked", "true");
    }

    if (filters.sort) {
        params.append("sort", filters.sort);
    }

    let url = "/posts";

    if (params.toString()) {
        url += `?${params.toString()}`;
    }

    return apiFetch(url, {
        method: "GET",
    });
}


export function getPostRequest(id) {
    return apiFetch(`/posts/${id}`, {
        method: "GET",
    });
}



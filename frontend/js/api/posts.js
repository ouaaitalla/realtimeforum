


export function createPostRequest(post) {
    return apiFetch("/posts", {
        method: "POST",
        body: JSON.stringify(post),
    });
}

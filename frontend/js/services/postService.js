import { createPostRequest } from "../api/posts.js";

export async function createPost(post) {
    const response = await createPostRequest(post);

    if (!response.success) {
        throw new Error(response.message);
    }

    return response.data;
}

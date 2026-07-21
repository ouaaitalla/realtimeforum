
import {createPostRequest, getPostsRequest , getPostRequest} from "../api/posts.js";

export async function createPost(post) {
    return (await createPostRequest(post)).data;
}  

export async function getPosts(filters = {}) {
    const data = (await getPostsRequest(filters)).data;
    return data.posts || [];
}

export async function getPostsWithCursor(filters = {}) {
    const response = await getPostsRequest(filters);
    return response.data;
}


export async function getPost(id) {
    return (await getPostRequest(id)).data;
}


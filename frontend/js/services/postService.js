
import {createPostRequest, getPostsRequest , getPostRequest} from "../api/posts.js";

export async function createPost(post) {
    return (await createPostRequest(post)).data;
}  

export async function getPosts(filters = {}) {
    return (await getPostsRequest(filters)).data;
}


export async function getPost(id) {
    return (await getPostRequest(id)).data;
}


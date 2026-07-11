
import {createPostRequest, getPostsRequest , getPostRequest} from "../api/posts.js";

export async function createPost(post) {
    return (await createPostRequest(post)).data;
}  

export async function getPosts() {
    return (await getPostsRequest()).data;
}


export async function getPost(id) {
    return (await getPostRequest(id)).data;
}



import {createPostRequest, getPostsRequest,} from "../api/posts.js";

export async function createPost(post) {
    return (await createPostRequest(post)).data;
}  

export async function getPosts() {
    return (await getPostsRequest()).data;
}

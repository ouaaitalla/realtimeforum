import { getCommentsRequest , createCommentRequest } from "../api/comments.js";

export async function getComments(postId) {
    return (await getCommentsRequest(postId)).data;
}

export async function createComment(postId, content) {
    return (await createCommentRequest(postId, content)).data;
}


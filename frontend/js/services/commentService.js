import { getCommentsRequest , createCommentRequest } from "../api/comments.js";

export async function getComments(postId, limit = 10, offset = 0) {
    const response = await getCommentsRequest(postId, limit, offset);
    return response.data;
}

export async function createComment(postId, content) {
    return (await createCommentRequest(postId, content)).data;
}


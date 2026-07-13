import {
    togglePostReactionRequest,
    toggleCommentReactionRequest,
} from "../api/reactions.js";

export async function togglePostReaction(postId, reaction) {

    return (
        await togglePostReactionRequest(
            postId,
            reaction,
        )
    ).data;
}

export async function toggleCommentReaction(commentId, reaction) {

    return (
        await toggleCommentReactionRequest(
            commentId,
            reaction,
        )
    ).data;
}


import {
    togglePostReaction,
    toggleCommentReaction,
} from "../services/reactionService.js";

let postReactionsInitialized = false;
let commentReactionsInitialized = false;

/**
 * Sets up a single delegated click listener on the document for post like/dislike buttons.
 * Works on both the feed page (post cards in #posts-container) and the details page.
 */
export function initPostReactions() {

    if (postReactionsInitialized) return;
    postReactionsInitialized = true;

    document.addEventListener("click", async (e) => {

        const likeBtn = e.target.closest(".post-like-btn");
        if (likeBtn) {
            e.stopPropagation();
            const postId = likeBtn.dataset.postId;
            await updatePostReaction(postId, 1);
            return;
        }

        const dislikeBtn = e.target.closest(".post-dislike-btn");
        if (dislikeBtn) {
            e.stopPropagation();
            const postId = dislikeBtn.dataset.postId;
            await updatePostReaction(postId, -1);
            return;
        }

    });

}

async function updatePostReaction(postId, reaction) {

    try {

        const data = await togglePostReaction(
            postId,
            reaction,
        );

        const likeBtn = document.querySelector(
            `.post-like-btn[data-post-id="${postId}"]`
        );

        const dislikeBtn = document.querySelector(
            `.post-dislike-btn[data-post-id="${postId}"]`
        );

        likeBtn.textContent = `👍 ${data.likes}`;
        dislikeBtn.textContent = `👎 ${data.dislikes}`;

        likeBtn.classList.toggle(
            "active",
            data.user_reaction === 1,
        );

        dislikeBtn.classList.toggle(
            "active",
            data.user_reaction === -1,
        );

    } catch (err) {
        console.error(err);
    }

}

/**
 * Sets up a single delegated click listener on the document for comment like/dislike buttons.
 */
export function initCommentReactions() {

    if (commentReactionsInitialized) return;
    commentReactionsInitialized = true;

    document.addEventListener("click", async (e) => {

        const likeBtn = e.target.closest(".comment-like-btn");
        if (likeBtn) {
            e.stopPropagation();
            const commentId = likeBtn.dataset.commentId;
            await updateCommentReaction(commentId, 1);
            return;
        }

        const dislikeBtn = e.target.closest(".comment-dislike-btn");
        if (dislikeBtn) {
            e.stopPropagation();
            const commentId = dislikeBtn.dataset.commentId;
            await updateCommentReaction(commentId, -1);
            return;
        }

    });

}

async function updateCommentReaction(commentId, reaction) {

    try {

        const data = await toggleCommentReaction(
            commentId,
            reaction,
        );

        const likeBtn = document.querySelector(
            `.comment-like-btn[data-comment-id="${commentId}"]`
        );

        const dislikeBtn = document.querySelector(
            `.comment-dislike-btn[data-comment-id="${commentId}"]`
        );

        likeBtn.textContent = `👍 ${data.likes}`;
        dislikeBtn.textContent = `👎 ${data.dislikes}`;

        likeBtn.classList.toggle(
            "active",
            data.user_reaction === 1,
        );

        dislikeBtn.classList.toggle(
            "active",
            data.user_reaction === -1,
        );

    } catch (err) {
        console.error(err);
    }

}

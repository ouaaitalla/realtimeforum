import {
    togglePostReaction,
    toggleCommentReaction,
} from "../services/reactionService.js";

export function initPostReactions() {

    document.querySelectorAll(".post-like-btn").forEach(button => {

        button.addEventListener("click", async (e) => {

            e.stopPropagation();

            const postId = button.dataset.postId;

            await updatePostReaction(
                postId,
                1,
            );

        });

    });

    document.querySelectorAll(".post-dislike-btn").forEach(button => {

        button.addEventListener("click", async (e) => {

            e.stopPropagation();

            const postId = button.dataset.postId;

            await updatePostReaction(
                postId,
                -1,
            );

        });

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

export function initCommentReactions() {

    document.querySelectorAll(".comment-like-btn").forEach(button => {

        button.addEventListener("click", async (e) => {

            e.stopPropagation();

            const commentId = button.dataset.commentId;

            await updateCommentReaction(
                commentId,
                1,
            );

        });

    });

    document.querySelectorAll(".comment-dislike-btn").forEach(button => {

        button.addEventListener("click", async (e) => {

            e.stopPropagation();

            const commentId = button.dataset.commentId;

            await updateCommentReaction(
                commentId,
                -1,
            );

        });

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

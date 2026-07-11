import { createComment } from "../services/commentService.js";
import { showNotification } from "./notification.js";
import { commentCard } from "./commentCard.js";

export function commentForm() {
    return `
        <form id="comment-form" class="comment-form">

            <textarea
                id="comment-content"
                class="comment-input"
                placeholder="Write a comment..."
                 maxlength="1000"
            ></textarea>

            <button type="submit" class="comment-submit">
                Comment
            </button>

        </form>
    `;
}

export function initCommentForm(postId) {

    const form = document.getElementById("comment-form");

    if (!form) return;

    form.addEventListener("submit", async (event) => {

        event.preventDefault();

        const textarea = document.getElementById("comment-content");

        const content = textarea.value.trim();

        if (!content) {

            showNotification(
                "Comment cannot be empty",
                "error"
            );

            return;
        }

        try {

            const comment = await createComment(
                postId,
                content
            );

            const commentsList =
                document.getElementById("comments-list");

            const empty =
                commentsList.querySelector(".empty-comments");

            if (empty) {
                empty.remove();
            }

            commentsList.insertAdjacentHTML(
                "beforeend",
                commentCard(comment)
            );

            textarea.value = "";

            showNotification(
                "Comment added successfully!",
                "success"
            );

        } catch (error) {

            showNotification(
                error.message,
                "error"
            );

        }

    });

}
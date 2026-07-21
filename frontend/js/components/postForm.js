import { createPost } from "../services/postService.js";
import { closeModal } from "./modal.js";
import { showNotification } from "./notification.js";
import { postCard, initPostCards } from "./postCard.js";
import { initPostReactions } from "./reactions.js";



export function postForm() {
    return `
        <form id="create-post-form" class="post-form">

            <h2>Create Post</h2>

            <div class="form-group">
                <label for="post-title">Title</label>

                <input
                    type="text"
                    id="post-title"
                    placeholder="Enter post title"
                    maxlength="100"
                >
            </div>

            <div class="form-group">
                <label for="post-content">Content</label>

                <textarea
                    id="post-content"
                    rows="6"
                    placeholder="What's on your mind?"
                ></textarea>
            </div>

            <div class="form-group">

                <label>Categories</label>

                <div class="categories">

                    <label>
                        <input type="checkbox" value="1">
                        General
                    </label>

                    <label>
                        <input type="checkbox" value="2">
                        Programming
                    </label>

                    <label>
                        <input type="checkbox" value="3">
                        Sports
                    </label>

                    <label>
                        <input type="checkbox" value="4">
                        Technology
                    </label>

                    <label>
                        <input type="checkbox" value="5">
                        News
                    </label>

                </div>

            </div>

            <button type="submit" class="publish-btn">
                Publish
            </button>

        </form>
    `;
}


export function initCreatePostForm() {

    const form = document.getElementById("create-post-form");

    if (!form) return;

    form.addEventListener("submit", async (event) => {

        event.preventDefault();

        const title = document.getElementById("post-title").value.trim();
        const content = document.getElementById("post-content").value.trim();

        const categories = Array.from(
            document.querySelectorAll(".categories input:checked")
        ).map(input => Number(input.value));

        const post = {
            title,
            content,
            categories,
        };

        try {
            const newPost = await createPost(post);
            closeModal();
            showNotification(
                "Post created successfully!",
                "success"
            );
            const postsContainer = document.getElementById("posts-container");

            if (postsContainer) {
                postsContainer.insertAdjacentHTML(
                    "afterbegin",
                    postCard(newPost)
                );

                initPostCards();
                initPostReactions();
            }
        } catch (error) {
            showNotification(
                error.message,
                "error"
            );

        }

    });

}



import { navigate } from "../router.js";
import { escapeHTML } from "../utils/escapeHTML.js";

let postCardsInitialized = false;

export function postCard(post) {
    return `
        <article
            class="post-card"
            data-post-id="${post.id}"
        >

            <div class="post-header">

                <div class="post-user">

                    <h4 class="post-author">
                        ${post.author}
                    </h4>

                    <span class="post-date">
                        ${new Date(post.created_at).toLocaleString()}
                    </span>

                </div>

                <div class="post-categories">
                    ${post.categories.map(category => `
                        <span class="category-badge">
                            ${category}
                        </span>
                    `).join("")}
                </div>

            </div>

            <h2 class="post-title">${escapeHTML(post.title)}</h2>

            <p class="post-content">${escapeHTML(post.content)}</p>

          <div class="post-footer">

            <button class="post-action-btn post-like-btn ${post.user_reaction === 1 ? "active" : ""}"data-post-id="${post.id}">👍 ${post.likes}</button>

            <button class="post-action-btn post-dislike-btn ${post.user_reaction === -1 ? "active" : ""}"data-post-id="${post.id}">👎 ${post.dislikes}</button>

            <button class="post-action-btn">💬 ${post.comments}</button>

        </div>

        </article>
    `;
}

/**
 * Sets up a single delegated click listener on the document for post card navigation.
 * Only navigates to the post detail page when clicking a post card on the feed
 * (i.e., inside #posts-container). Reaction button clicks are ignored.
 */
export function initPostCards() {

    if (postCardsInitialized) return;
    postCardsInitialized = true;

    document.addEventListener("click", (e) => {

        // Only handle clicks on post cards inside the feed container
        const container = document.getElementById("posts-container");
        if (!container) return;

        // Don't navigate if clicking reaction buttons
        if (e.target.closest(".post-like-btn") || e.target.closest(".post-dislike-btn")) {
            return;
        }

        const card = e.target.closest(".post-card");
        if (card && container.contains(card)) {
            const postId = card.dataset.postId;
            navigate(`/posts/${postId}`);
        }
    });

}


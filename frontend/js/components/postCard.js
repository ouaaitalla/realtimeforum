
import { navigate } from "../router.js";


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

            <h2 class="post-title">
                ${post.title}
            </h2>

            <p class="post-content">
                ${post.content}
            </p>

          <div class="post-footer">

            <button class="post-action-btn post-like-btn ${post.user_reaction === 1 ? "active" : ""}"data-post-id="${post.id}">👍 ${post.likes}</button>

            <button class="post-action-btn post-dislike-btn ${post.user_reaction === -1 ? "active" : ""}"data-post-id="${post.id}">👎 ${post.dislikes}</button>

            <button class="post-action-btn">💬 ${post.comments}</button>

        </div>

        </article>
    `;
}

export function initPostCards() {

    const cards = document.querySelectorAll(".post-card");

    cards.forEach(card => {

        card.addEventListener("click", (e) => {

            if (
                e.target.closest(".post-like-btn") ||
                e.target.closest(".post-dislike-btn")
            ) {
                return;
            }

            const postId = card.dataset.postId;

            navigate(`/posts/${postId}`);

        });

    });

}


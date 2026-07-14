import { escapeHTML } from "../utils/escapeHTML.js";

export function commentCard(comment) {
    return `
        <article class="comment-card">

            <div class="comment-header">

                <div>
                    <h4 class="comment-author">
                        ${comment.author}
                    </h4>

                    <span class="comment-date">
                        ${new Date(comment.created_at).toLocaleString()}
                    </span>
                </div>

            </div>

            <p class="comment-content">${escapeHTML(comment.content)}</p>

        <div class="comment-footer">

            <button
                class="comment-like-btn ${comment.user_reaction === 1 ? "active" : ""}"
                data-comment-id="${comment.id}"
            >
                👍 ${comment.likes}
            </button>

            <button
                class="comment-dislike-btn ${comment.user_reaction === -1 ? "active" : ""}"
                data-comment-id="${comment.id}"
            >
                👎 ${comment.dislikes}
            </button>

        </div>

        </article>
    `;
}

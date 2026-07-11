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

            <p class="comment-content">
                ${comment.content}
            </p>

            <div class="comment-footer">

                <button class="comment-like-btn">
                    👍 0
                </button>

                <button class="comment-dislike-btn">
                    👎 0
                </button>

            </div>

        </article>
    `;
}

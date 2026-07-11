export function postCard(post) {
    return `
        <article class="post-card">

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
                    ${(post.categories || []).map(category => `
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

                <button class="post-action like-btn" data-id="${post.id}">
                    👍
                    <span>${post.likes ?? 0}</span>
                </button>

                <button class="post-action dislike-btn" data-id="${post.id}">
                    👎
                    <span>${post.dislikes ?? 0}</span>
                </button>

                <button class="post-action comments-btn" data-id="${post.id}">
                    💬
                    <span>${post.comments_count ?? 0}</span>
                </button>

            </div>

        </article>
    `;
}

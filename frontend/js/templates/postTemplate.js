import { postCard } from "../components/postCard.js";
import { commentCard } from "../components/commentCard.js";
import { commentForm } from "../components/commentForm.js";

export function postDetailsTemplate(post, comments, hasMoreComments = false) {

    return `

        <section class="post-details">

            ${postCard(post)}

            <div class="comments-section">

                <h2>
                    Comments (${comments.length})
                </h2>

                ${commentForm()}

                <div id="comments-list">

                    ${
                        comments.length
                        ? comments.map(commentCard).join("")
                        : `
                            <p class="empty-comments">
                                No comments yet.
                            </p>
                        `
                    }

                </div>

                ${
                    hasMoreComments
                        ? `<div class="load-more-comments">
                            <button id="load-more-comments-btn" class="load-more-btn">
                                Load more comments
                            </button>
                        </div>`
                        : ""
                }

            </div>

        </section>

    `;
}

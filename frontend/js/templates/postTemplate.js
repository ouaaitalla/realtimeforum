import { postCard } from "../components/postCard.js";
import { commentCard } from "../components/commentCard.js";
import { commentForm } from "../components/commentForm.js";

export function postDetailsTemplate(post, comments) {

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

            </div>

        </section>

    `;
}

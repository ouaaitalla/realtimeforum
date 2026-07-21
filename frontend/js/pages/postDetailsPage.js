import { render } from "../utils/render.js";
import { appLayout } from "../layouts/appLayout.js";
import { postDetailsTemplate } from "../templates/postTemplate.js";
import { initNavbar } from "../components/navbar.js";
import { getPost } from "../services/postService.js";
import { getComments } from "../services/commentService.js";
import { initCommentForm } from "../components/commentForm.js";
import { initPostReactions, initCommentReactions } from "../components/reactions.js";
import { commentCard } from "../components/commentCard.js";

const PAGE_SIZE = 10;

export async function postDetailsPage(id) {

    let currentOffset = 0;
    let allComments = [];
    let hasMore = true;

    try {

        const post = await getPost(id);

        // Load first page of comments
        const firstPage = await getComments(id, PAGE_SIZE, 0);

        allComments = firstPage || [];
        currentOffset = allComments.length;
        hasMore = allComments.length >= PAGE_SIZE;

        render(
            appLayout(
                postDetailsTemplate(post, allComments, hasMore)
            )
        );

        initNavbar();

        initPostReactions();

        initCommentReactions();

        initCommentForm(post.id);

        // Setup "Load more" button
        const loadMoreBtn = document.getElementById("load-more-comments-btn");
        if (loadMoreBtn) {
            loadMoreBtn.addEventListener("click", async () => {
                loadMoreBtn.disabled = true;
                loadMoreBtn.textContent = "Loading...";

                try {
                    const moreComments = await getComments(id, PAGE_SIZE, currentOffset);

                    if (!moreComments || moreComments.length === 0) {
                        hasMore = false;
                        loadMoreBtn.remove();
                        return;
                    }

                    const commentsList = document.getElementById("comments-list");

                    moreComments.forEach(comment => {
                        commentsList.insertAdjacentHTML(
                            "beforeend",
                            commentCard(comment)
                        );
                    });

                    initCommentReactions();

                    currentOffset += moreComments.length;
                    hasMore = moreComments.length >= PAGE_SIZE;

                    if (!hasMore) {
                        loadMoreBtn.remove();
                    } else {
                        loadMoreBtn.disabled = false;
                        loadMoreBtn.textContent = "Load more comments";
                    }
                } catch (err) {
                    loadMoreBtn.disabled = false;
                    loadMoreBtn.textContent = "Load more comments";
                    console.error("Failed to load more comments:", err);
                }
            });
        }

    } catch (error) {

        render(
            appLayout(`
                <h2>
                    ${error.message}
                </h2>
            `)
        );

    }

}

import { appLayout } from "../layouts/appLayout.js";
import { initNavbar } from "../components/navbar.js";
import {homeTemplate} from "../templates/homeTemplate.js";
import { render } from "../utils/render.js";
import { getPostsWithCursor } from "../services/postService.js";
import { postCard } from "../components/postCard.js";
import { showNotification } from "../components/notification.js";
import { initPostCards } from "../components/postCard.js";
import { initPostReactions } from "../components/reactions.js";
import { initFilterModal } from "../components/filterModal.js";
import { debounce } from "../utils/debounce.js";

const PAGE_SIZE = 10;

let currentCursor = null;
let currentCursorID = null;
let hasMorePosts = true;
let isLoadingPosts = false;
let currentFilters = {};


export async function homePage() {

    // Reset pagination state
    currentCursor = null;
    hasMorePosts = true;
    isLoadingPosts = false;
    currentFilters = {};

    render(
        appLayout(homeTemplate())
    );

    initNavbar();

    initFilterModal();

    // Load initial batch
    await loadPosts();

    // Setup infinite scroll
    setupInfiniteScroll();

}


export async function loadPosts(filters = {}, append = false) {

    // When called with explicit filters (from filter modal), reset pagination
    if (!append) {
        currentFilters = { ...filters };
        currentCursor = null;
        hasMorePosts = true;
    }

    const container = document.getElementById("posts-container");

    if (!container) return;

    // Show loading state on first page
    if (!append) {
        container.innerHTML = `
            <div class="loading-posts">
                Loading posts...
            </div>
        `;
    }

    try {

        const data = await getPostsWithCursor({
            ...currentFilters,
            cursor: currentCursor,
            cursor_id: currentCursorID,
            limit: PAGE_SIZE,
        });

        const posts = data.posts;
        const nextCursor = data.next_cursor;
        const nextCursorID = data.next_cursor_id;
        const hasMore = data.has_more;

        if (!append) {
            // First page: replace content
            if (posts.length === 0) {
                container.innerHTML = `
                    <div class="empty-posts">
                        No posts yet.
                    </div>
                `;
                hasMorePosts = false;
                return;
            }

            container.innerHTML = posts
                .map(post => postCard(post))
                .join("");

        } else {
            // Next pages: append
            if (posts.length === 0) {
                hasMorePosts = false;
                return;
            }

            posts.forEach(post => {
                container.insertAdjacentHTML(
                    "beforeend",
                    postCard(post)
                );
            });
        }

        initPostCards();
        initPostReactions();

        currentCursor = nextCursor || null;
        currentCursorID = nextCursorID || null;
        hasMorePosts = hasMore;

    } catch (error) {

        if (!append) {
            container.innerHTML = `
                <div class="empty-posts">
                    Failed to load posts.
                </div>
            `;
        }
    }
}


function setupInfiniteScroll() {

    const handleScroll = debounce(() => {

        if (isLoadingPosts || !hasMorePosts) return;

        const scrollBottom = window.innerHeight + window.scrollY;
        const pageHeight = document.documentElement.scrollHeight;

        // Load more when within 300px of the bottom
        if (scrollBottom >= pageHeight - 300) {
            isLoadingPosts = true;

            loadPosts({}, true).finally(() => {
                isLoadingPosts = false;
            });
        }

    }, 200);

    window.addEventListener("scroll", handleScroll);

}


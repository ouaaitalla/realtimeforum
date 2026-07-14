import { appLayout } from "../layouts/appLayout.js";
import { initNavbar } from "../components/navbar.js";
import {homeTemplate} from "../templates/homeTemplate.js";
import { render } from "../utils/render.js";
import { getPosts } from "../services/postService.js";
import { postCard } from "../components/postCard.js";
import { showNotification } from "../components/notification.js";
import { initPostCards } from "../components/postCard.js";
import { initPostReactions } from "../components/reactions.js";
import { initFilterModal } from "../components/filterModal.js";


export async function homePage() {

    render(
        appLayout(homeTemplate())
    );

    initNavbar();
    
    initFilterModal();
    
    await loadPosts();
    

}

export async function loadPosts(filters = {}) {

    const container = document.getElementById("posts-container");

    container.innerHTML = `
        <div class="loading-posts">
            Loading posts...
        </div>
    `;

    try {

        const posts = await getPosts(filters);

        if (posts.length === 0) {
            container.innerHTML = `
                <div class="empty-posts">
                    No posts yet.
                </div>
            `;
            return;
        }

        container.innerHTML = posts
            .map(post => postCard(post))
            .join("");

        initPostCards();
        initPostReactions();

    } catch (error) {

        container.innerHTML = `
            <div class="empty-posts">
                Failed to load posts.
            </div>
        `;
    }
}


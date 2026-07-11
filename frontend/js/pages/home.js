import { appLayout } from "../layouts/appLayout.js";
import { initNavbar } from "../components/navbar.js";
import {homeTemplate} from "../templates/homeTemplate.js";
import { render } from "../utils/render.js";
import { getPosts } from "../services/postService.js";
import { postCard } from "../components/postCard.js";
import { showNotification } from "../components/notification.js";
import { initPostCards } from "../components/postCard.js";


export async function homePage() {

    render(
        appLayout(homeTemplate())
    );

    initNavbar();

    await loadPosts();

}

export async function loadPosts() {

    const container = document.getElementById("posts-container");

    container.innerHTML = `
        <div class="loading-posts">
            Loading posts...
        </div>
    `;

    try {

        const posts = await getPosts();

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

    } catch (error) {

        container.innerHTML = `
            <div class="empty-posts">
                Failed to load posts.
            </div>
        `;
    }
}


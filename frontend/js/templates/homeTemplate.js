import { filterModal } from "./filterModal.js";

export function homeTemplate() {
    return `
        <section class="feed-header">

            <h2 class="feed-title">
                Recent Posts
            </h2>

            <button
                id="open-filter-modal"
                class="filter-btn"
            >
                🔍 Filter
            </button>

        </section>

        <section id="posts-container"></section>

        ${filterModal()}
    `;
}

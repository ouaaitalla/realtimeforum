import { getCategories } from "../services/categoryService.js";

export function filterModal() {
    return `
        <div
            id="filter-modal"
            class="modal-overlay hidden"
        >

            <div class="modal-content">

                <button
                    id="close-filter-modal"
                    class="modal-close"
                >
                    &times;
                </button>

                <h2>
                    Filter Posts
                </h2>

                <div class="filter-section">

                    <h3>Categories</h3>

                    <div id="filter-categories">

                    </div>

                </div>

                <div class="filter-section">

                    <h3>Sort</h3>

                    <div class="sort-options">

                        <label class="sort-chip">

                            <input
                                type="radio"
                                name="sort"
                                value=""
                                checked
                            >

                            <span>Recent</span>

                        </label>

                        <label class="sort-chip">

                            <input
                                type="radio"
                                name="sort"
                                value="mostliked"
                            >

                            <span>Most Liked</span>

                        </label>

                    </div>

                </div>

                <div class="filter-section">

                    <div class="filter-checkboxes">

                        <label>

                            <input
                                id="filter-mine"
                                type="checkbox"
                            >

                            My Posts

                        </label>

                        <label>

                            <input
                                id="filter-liked"
                                type="checkbox"
                            >

                            Liked Posts

                        </label>

                    </div>

                </div>

                <div class="filter-actions">

                    <button id="reset-filters">
                        Reset
                    </button>

                    <button id="apply-filters">
                        Apply
                    </button>

                </div>

            </div>

        </div>
    `;
}

export async function loadCategories() {

    const container = document.getElementById("filter-categories");

    if (!container) return;

    const categories = await getCategories();

    container.innerHTML = "";

    categories.forEach(category => {

        container.innerHTML += `
            <label class="category-chip">

                <input
                    type="checkbox"
                    name="category"
                    value="${category.name}"
                >

                <span>${category.name}</span>

            </label>
        `;

    });

}

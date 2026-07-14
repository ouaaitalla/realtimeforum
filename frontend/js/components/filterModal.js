import { loadPosts } from "../pages/home.js";
import { loadCategories } from "../templates/filterModal.js";

export function initFilterModal() {

    const openBtn = document.getElementById("open-filter-modal");
    const closeBtn = document.getElementById("close-filter-modal");

    const modal = document.getElementById("filter-modal");

    const applyBtn = document.getElementById("apply-filters");
    const resetBtn = document.getElementById("reset-filters");

    if (!modal) return;

    loadCategories();

    openBtn?.addEventListener("click", () => {

        modal.classList.remove("hidden");

    });

    closeBtn?.addEventListener("click", () => {

        modal.classList.add("hidden");

    });

    modal.addEventListener("click", (e) => {

        if (e.target === modal) {

            modal.classList.add("hidden");

        }

    });

    applyBtn?.addEventListener("click", async () => {

        const filters = {};

        // Categories
        const selectedCategories = [
            ...document.querySelectorAll(
                'input[name="category"]:checked'
            )
        ].map(input => input.value);

        if (selectedCategories.length > 0) {

            filters.categories = selectedCategories;

        }

        // Sort
        const sort = document.querySelector(
            'input[name="sort"]:checked'
        );

        if (sort && sort.value !== "") {

            filters.sort = sort.value;

        }

        // My Posts
        if (document.getElementById("filter-mine").checked) {

            filters.mine = true;

        }

        // Liked Posts
        if (document.getElementById("filter-liked").checked) {

            filters.liked = true;

        }

        modal.classList.add("hidden");

        await loadPosts(filters);

    });

    resetBtn?.addEventListener("click", async () => {

        document
            .querySelectorAll('input[name="category"]')
            .forEach(input => input.checked = false);

        document
            .querySelector('input[name="sort"][value=""]')
            .checked = true;

        document.getElementById("filter-mine").checked = false;

        document.getElementById("filter-liked").checked = false;

        modal.classList.add("hidden");

        await loadPosts({});

    });

}
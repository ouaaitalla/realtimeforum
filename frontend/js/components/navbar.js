import { logout } from "../services/authService.js";
import { navigate } from "../router.js";
import { postForm } from "./postForm.js";
import { openModal } from "./modal.js";


export function navbar() {
    return `
        <nav class="navbar">

            <div class="logo">
                Forum
            </div>

            <div class="nav-links">

                <a id="home-link">
                    Home
                </a>

                <button id="create-post-btn">
                    Create Post
                </button>

                <a id="messages-link">
                    Messages
                </a>

                <button id="logout-btn">
                    Logout
                </button>

            </div>

        </nav>
    `;
}

export function initNavbar() {

    const logoutBtn = document.getElementById("logout-btn");
    const homeLink = document.getElementById("home-link");
    const messagesLink = document.getElementById("messages-link");
    const createPostBtn = document.getElementById("create-post-btn");

    if (homeLink) {
        homeLink.addEventListener("click", () => {
            navigate("/");
        });
    }

    if (messagesLink) {
        messagesLink.addEventListener("click", () => {
            navigate("/messages");
        });
    }

    if (createPostBtn) {
        createPostBtn.addEventListener("click", () => {
            // Open Create Post Modal
            openModal(postForm());
            initCreatePostForm();
        });
    }

    if (logoutBtn) {
        logoutBtn.addEventListener("click", async () => {
            try {
                await logout();
                navigate("/login");
            } catch (error) {
                console.error(error);
            }
        });
    }

}
import { navbar } from "../components/navbar.js";

export function appLayout(content) {
    return `
        ${navbar()}

        <main class="main-content">
            ${content}
        </main>
    `;
}
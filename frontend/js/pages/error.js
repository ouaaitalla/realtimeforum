

export function errorPage() {

    const app = document.getElementById("app");

    app.innerHTML = `
        <div class="error-page">
            <h1>404</h1>
            <p>Page not found.</p>
        </div>
    `;
}

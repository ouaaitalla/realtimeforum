

export function render(content) {

    // Remove any leftover modal that may exist outside #app
    const leftover = document.getElementById("modal");
    if (leftover) {
        leftover.remove();
    }

    const app = document.getElementById("app");

    app.innerHTML = content;
}
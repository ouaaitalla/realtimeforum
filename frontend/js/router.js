import { homePage } from "./pages/home.js";
import { loginPage } from "./pages/login.js";
import { registerPage } from "./pages/register.js";
import { postDetailsPage } from "./pages/postDetailsPage.js";
import { checkAuth } from "./services/authService.js";
import { errorPage } from "./pages/error.js";

const routes = {
    "/": homePage,
    "/login": loginPage,
    "/register": registerPage,
};

export async function router() {

    const path = window.location.pathname;

    const isAuthenticated = await checkAuth();

    // User not logged in
    if (!isAuthenticated && path !== "/login" && path !== "/register") {
        navigate("/login");
        return;
    }

    // User already logged in
    if (isAuthenticated && (path === "/login" || path === "/register")) {
        navigate("/");
        return;
    }

    // Dynamic Route: /posts/:id
    if (path.startsWith("/posts/")) {

        const id = path.split("/")[2];

        if (!id || isNaN(id)) {
            errorPage();
            return;
        }

        postDetailsPage(Number(id));
        return;
    }

    const page = routes[path];

    if (!page) {
        errorPage();
        return;
    }

    page();

}

export function navigate(path) {
    window.history.pushState({}, "", path);
    router();
}

window.addEventListener("popstate", router);
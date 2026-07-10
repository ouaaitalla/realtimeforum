// import { homePage } from "./pages/home.js";
import { loginPage } from "./pages/login.js";
import { registerPage } from "./pages/register.js";
// import { profilePage } from "./pages/profile.js";
import { checkAuth } from "./services/authService.js";

const routes = {
    // "/": homePage,
    "/login": loginPage,
    "/register": registerPage,
    // "/profile": profilePage,
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

    const page = routes[path] || routes["/"];

    page();
}


export function navigate(path) {
    window.history.pushState({}, "", path);
    router();
}


window.addEventListener("popstate", router);
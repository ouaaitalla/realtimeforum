import { appLayout } from "../layouts/appLayout.js";
import { initNavbar } from "../components/navbar.js";
import {homeTemplate} from "../templates/homeTemplate.js"

export function homePage() {

    const app = document.getElementById("app");

    app.innerHTML = appLayout(
        homeTemplate()
    );

    initNavbar();

}
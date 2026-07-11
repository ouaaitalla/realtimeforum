import { loginTemplate } from "../templates/loginTemplate.js";
import { login } from "../services/authService.js";
import { showNotification } from "../components/notification.js";
import { validateLogin } from "../utils/validator.js";
import { navigate } from "../router.js";


export function loginPage() {

    const app = document.getElementById("app");

    app.innerHTML = loginTemplate();


    const form = document.getElementById("login-form");


    form.addEventListener("submit", async (event) => {

        event.preventDefault();

        const credentials = {

            email: document
                .getElementById("email")
                .value
                .trim(),

            password: document
                .getElementById("password")
                .value

        };


        // Frontend validation
        const errors = validateLogin(credentials);


        if (errors.length > 0) {

            showNotification(
                errors[0],
                "error"
            );

            return;
        }


        try {

            await login(credentials);


            showNotification(
                "Login successful",
                "success"
            );


            setTimeout(() => {
                navigate("/");
            }, 700);

        } catch (error) {

            showNotification(
                error.message,
                "error"
            );

        }

    });
}

import { registerTemplate } from "../templates/registerTemplate.js";
import { register } from "../services/authService.js";
import { validateRegister } from "../utils/validator.js";
import { showNotification } from "../components/notification.js";
import { navigate } from "../router.js";
import { render } from "../utils/render.js";

export function registerPage() {
   render(registerTemplate());

    const form = document.getElementById("register-form");

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        console.log("Submit clicked");
        const user = {
            nickname: document.getElementById("nickname").value.trim(),
            first_name: document.getElementById("first-name").value.trim(),
            last_name: document.getElementById("last-name").value.trim(),
            email: document.getElementById("email").value.trim(),
            password: document.getElementById("password").value,
            age: Number(document.getElementById("age").value),
            gender: document.getElementById("gender").value,
        };

        try {
            const errors = validateRegister(user);
            if (errors.length > 0) {
                showNotification(   
                errors[0],
                "error"
                );
                return;
            }
            await register(user);

            showNotification(
                "Registration successful!",
                "success"
            );
            setTimeout(() => {
                navigate("/login");
            }, 700);
        } catch (error) {
            showNotification(
                error.message,
                "error"
            );
        }
    });
}


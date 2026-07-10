export function registerTemplate() {
    return `
        <section class="auth-page">

            <div class="auth-container">

                <form id="register-form" class="auth-form">

                    <h2>Create Account</h2>

                    <div id="auth-message" class="auth-error"></div>


                    <input
                        type="text"
                        id="nickname"
                        name="nickname"
                        placeholder="Nickname"
                        required
                    >


                    <input
                        type="text"
                        id="first-name"
                        name="first_name"
                        placeholder="First Name"
                        required
                    >


                    <input
                        type="text"
                        id="last-name"
                        name="last_name"
                        placeholder="Last Name"
                        required
                    >


                    <input
                        type="email"
                        id="email"
                        name="email"
                        placeholder="Email"
                        required
                    >


                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        required
                    >


                    <input
                        type="number"
                        id="age"
                        name="age"
                        placeholder="Age"
                        min="13"
                    >


                    <select id="gender" name="gender">

                        <option value="">
                            Select Gender
                        </option>

                        <option value="Male">
                            Male
                        </option>

                        <option value="Female">
                            Female
                        </option>

                    </select>


                    <button type="submit">
                        Register
                    </button>


                    <p class="auth-link">

                        Already have an account?

                        <a href="/login" data-link>
                            Login
                        </a>

                    </p>


                </form>

            </div>

        </section>
    `;
}

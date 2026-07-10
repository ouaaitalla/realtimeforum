export function loginTemplate() {
    return `
        <section class="auth-page">

            <div class="auth-container">

                <form id="login-form" class="auth-form">

                    <h2>Welcome Back</h2>

                    <div id="auth-message" class="auth-error"></div>


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


                    <button type="submit">
                        Login
                    </button>


                    <p class="auth-link">

                        Don't have an account?

                        <a href="/register" data-link>
                            Register
                        </a>

                    </p>


                </form>

            </div>

        </section>
    `;
}

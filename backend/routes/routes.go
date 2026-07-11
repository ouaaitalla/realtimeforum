package routes

import (
	"net/http"

	"real-time-forum/backend/handlers/authHandlers"
	posthandlers "real-time-forum/backend/handlers/post-Handlers"
	"real-time-forum/backend/middleware"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("/register", authHandlers.RegisterHandler)

	mux.HandleFunc("/login", authHandlers.LoginHandler)

	mux.HandleFunc("/logout", authHandlers.LogoutHandler)

	mux.HandleFunc("/me", authHandlers.MeHandler)

	// Posts routes
	postsHandler := middleware.AuthMiddleware(
		http.HandlerFunc(posthandlers.PostsHandler),
	)

	mux.Handle("/posts", postsHandler)
	mux.Handle("/posts/", postsHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Server Running"))
	})

	return mux
}

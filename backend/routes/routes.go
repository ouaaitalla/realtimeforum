package routes

import (
	"net/http"

	"real-time-forum/backend/handlers/authHandlers"
	categoryhandlers "real-time-forum/backend/handlers/category-Handlers"
	commenthandlers "real-time-forum/backend/handlers/comment-Handlers"
	posthandlers "real-time-forum/backend/handlers/post-Handlers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/websocket"
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

	// Comments routes
	commentsHandler := middleware.AuthMiddleware(
		http.HandlerFunc(commenthandlers.CommentsHandler),
	)

	categoriesHandler := middleware.AuthMiddleware(
		http.HandlerFunc(categoryhandlers.GetCategoriesHandler),
	)

	mux.Handle("/categories", categoriesHandler)

	mux.Handle("/comments/", commentsHandler)

	mux.HandleFunc("/ws", websocket.HandleWebSocket)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Server Running"))
	})

	return mux
}

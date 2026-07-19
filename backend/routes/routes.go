package routes

import (
	"net/http"
	"strings"

	"real-time-forum/backend/handlers/authHandlers"
	categoryhandlers "real-time-forum/backend/handlers/category-Handlers"
	commenthandlers "real-time-forum/backend/handlers/comment-Handlers"
	messagehandlers "real-time-forum/backend/handlers/message-Handlers"
	posthandlers "real-time-forum/backend/handlers/post-Handlers"
	userhandlers "real-time-forum/backend/handlers/user-Handlers"
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

	wsHandler := middleware.AuthMiddleware(
		http.HandlerFunc(websocket.HandleWebSocket),
	)

	usersHandler := middleware.AuthMiddleware(
		http.HandlerFunc(userhandlers.UsersHandler),
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

	mux.Handle("/ws", wsHandler)

	mux.Handle("/users", usersHandler)

	messagesHandler := middleware.AuthMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/messages/read/") {
				messagehandlers.MarkAsReadHandler(w, r)
				return
			}

			messagehandlers.GetConversationHandler(w, r)
		}),
	)

	mux.Handle("/messages/", messagesHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Server Running"))
	})

	return mux
}

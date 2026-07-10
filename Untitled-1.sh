#!/bin/bash

# Create directories
mkdir -p backend/{db,handlers/{authHandlers,comment-Handlers,post-Handlers,user-Hnadlers},helpers,httpx,middleware,models,routes,websocket}
mkdir -p database
mkdir -p frontend/assets/css
mkdir -p frontend/js/{api,components,pages,services,templates,utils,websocket}

# Backend
touch backend/db/db.go

touch backend/handlers/authHandlers/loginHandler.go
touch backend/handlers/authHandlers/registerHandler.go

touch backend/handlers/comment-Handlers/creatCommentHandler.go
touch backend/handlers/comment-Handlers/deleteCommentHandler.go
touch backend/handlers/comment-Handlers/getCommentHandler.go

touch backend/handlers/post-Handlers/creatPostHandler.go
touch backend/handlers/post-Handlers/deletePostHandler.go
touch backend/handlers/post-Handlers/getPostHandler.go

touch backend/handlers/user-Hnadlers/apdateUserHandler.go
touch backend/handlers/user-Hnadlers/getUserHandler.go

touch backend/helpers/response.go
touch backend/helpers/security.go
touch backend/helpers/session.go

touch backend/httpx/httpx.go
touch backend/main.go

touch backend/middleware/authMiddleware.go
touch backend/middleware/mainMiddleware.go
touch backend/middleware/methodMiddleware.go

touch backend/models/models.go
touch backend/routes/routes.go
touch backend/websocket/handler.go

# Database
touch database/initDb.go
touch database/schema.sql
touch DB.db

# Frontend CSS
touch frontend/assets/css/auth.css
touch frontend/assets/css/chat.css
touch frontend/assets/css/comments.css
touch frontend/assets/css/feed.css
touch frontend/assets/css/main.css
touch frontend/assets/css/modal.css
touch frontend/assets/css/navbar.css
touch frontend/assets/css/post.css
touch frontend/assets/css/responsive.css

# Frontend root
touch frontend/index.html
touch frontend/README.md
touch frontend/server.js

# API
touch frontend/js/api/auth.js
touch frontend/js/api/chat.js
touch frontend/js/api/comments.js
touch frontend/js/api/posts.js
touch frontend/js/api/users.js

# JS root
touch frontend/js/app.js
touch frontend/js/router.js
touch frontend/js/state.js

# Components
touch frontend/js/components/chatList.js
touch frontend/js/components/chatWindow.js
touch frontend/js/components/commentCard.js
touch frontend/js/components/commentForm.js
touch frontend/js/components/loader.js
touch frontend/js/components/messageBubble.js
touch frontend/js/components/modal.js
touch frontend/js/components/navbar.js
touch frontend/js/components/notification.js
touch frontend/js/components/onlineUsers.js
touch frontend/js/components/postCard.js
touch frontend/js/components/postForm.js
touch frontend/js/components/sidebar.js

# Pages
touch frontend/js/pages/error.js
touch frontend/js/pages/home.js
touch frontend/js/pages/login.js
touch frontend/js/pages/profile.js
touch frontend/js/pages/register.js

# Services
touch frontend/js/services/authService.js
touch frontend/js/services/chatService.js
touch frontend/js/services/commentService.js
touch frontend/js/services/postService.js
touch frontend/js/services/websocketService.js

# Templates
touch frontend/js/templates/chatTemplate.js
touch frontend/js/templates/homeTemplate.js
touch frontend/js/templates/loginTemplate.js
touch frontend/js/templates/postTemplate.js
touch frontend/js/templates/registerTemplate.js

# Utils
touch frontend/js/utils/constants.js
touch frontend/js/utils/date.js
touch frontend/js/utils/debounce.js
touch frontend/js/utils/fetch.js
touch frontend/js/utils/render.js
touch frontend/js/utils/storage.js
touch frontend/js/utils/throttle.js
touch frontend/js/utils/validator.js

# WebSocket
touch frontend/js/websocket/messageEvents.js
touch frontend/js/websocket/notificationEvents.js
touch frontend/js/websocket/onlineEvents.js
touch frontend/js/websocket/socket.js

# Root files
touch go.mod
touch go.sum
touch message.txt

echo "✅ Project structure created successfully!"
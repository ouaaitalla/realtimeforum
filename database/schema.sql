PRAGMA foreign_keys = ON;

-------------------------------------------------
-- USERS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    nickname TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,

    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    age INTEGER,
    gender TEXT,

    avatar TEXT DEFAULT '',

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-------------------------------------------------
-- SESSIONS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY,

    user_id INTEGER NOT NULL UNIQUE,

    expires_at DATETIME NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- POSTS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    user_id INTEGER NOT NULL,

    title TEXT NOT NULL,

    content TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    updated_at DATETIME,

    FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- CATEGORIES
-------------------------------------------------
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    name TEXT NOT NULL UNIQUE
);

-------------------------------------------------
-- POST CATEGORIES
-------------------------------------------------
CREATE TABLE IF NOT EXISTS post_categories (
    post_id INTEGER NOT NULL,

    category_id INTEGER NOT NULL,

    PRIMARY KEY(post_id, category_id),

    FOREIGN KEY(post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,

    FOREIGN KEY(category_id)
        REFERENCES categories(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- COMMENTS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    post_id INTEGER NOT NULL,

    user_id INTEGER NOT NULL,

    content TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,

    FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- POST REACTIONS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS post_reactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    post_id INTEGER NOT NULL,

    user_id INTEGER NOT NULL,

    reaction INTEGER NOT NULL CHECK(reaction IN (1,-1)),

    UNIQUE(post_id, user_id),

    FOREIGN KEY(post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,

    FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- COMMENT REACTIONS
-------------------------------------------------
CREATE TABLE IF NOT EXISTS comment_reactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    comment_id INTEGER NOT NULL,

    user_id INTEGER NOT NULL,

    reaction INTEGER NOT NULL CHECK(reaction IN (1,-1)),

    UNIQUE(comment_id, user_id),

    FOREIGN KEY(comment_id)
        REFERENCES comments(id)
        ON DELETE CASCADE,

    FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- PRIVATE MESSAGES
-------------------------------------------------
CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    sender_id INTEGER NOT NULL,

    receiver_id INTEGER NOT NULL,

    content TEXT NOT NULL,

    is_read INTEGER DEFAULT 0,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(sender_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    FOREIGN KEY(receiver_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-------------------------------------------------
-- INDEXES
-------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_posts_user
ON posts(user_id);

CREATE INDEX IF NOT EXISTS idx_comments_post
ON comments(post_id);

CREATE INDEX IF NOT EXISTS idx_comments_user
ON comments(user_id);

CREATE INDEX IF NOT EXISTS idx_messages_sender
ON messages(sender_id);

CREATE INDEX IF NOT EXISTS idx_messages_receiver
ON messages(receiver_id);

CREATE INDEX IF NOT EXISTS idx_posts_created
ON posts(created_at DESC);

CREATE INDEX IF NOT EXISTS idx_messages_created
ON messages(created_at DESC);


INSERT OR IGNORE INTO categories (id, name) VALUES
(1, 'General'),
(2, 'Programming'),
(3, 'Sports'),
(4, 'Technology'),
(5, 'News');


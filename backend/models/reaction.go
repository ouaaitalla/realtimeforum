package models




type CommentReaction struct {
	ID        int `json:"id"`
	CommentID int `json:"comment_id"`
	UserID    int `json:"user_id"`
	Reaction  int `json:"reaction"`
}

type PostReaction struct {
	ID       int `json:"id"`
	PostID   int `json:"post_id"`
	UserID   int `json:"user_id"`
	Reaction int `json:"reaction"`
}

type ReactionResponse struct {
    Likes        int `json:"likes"`
    Dislikes     int `json:"dislikes"`
    UserReaction int `json:"user_reaction"`
}


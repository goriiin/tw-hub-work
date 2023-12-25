package domain

import "time"

type Post struct {
	PostID        int       `json:"postId"`
	UserID        int       `json:"userId"`
	Username      string    `json:"username"`
	Text          string    `json:"text"`
	Date          time.Time `json:"date"`
	IsLiked       bool      `json:"isLiked"`
	LikesCount    int       `json:"likesCount"`
	IsDisliked    bool      `json:"isDisliked"`
	DislikesCount int       `json:"dislikesCount"`
}

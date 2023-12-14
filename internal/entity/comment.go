package entity

import (
	"time"
)

// pure business logic
type Comment struct {
	Author      Author    `json:"author" bson:"author"`
	Body        string    `json:"body" bson:"body"`
	CreatedTime time.Time `json:"created" bson:"created"`
}

// business logic with implementation logic
type CommentExtend struct {
	Comment
	ID string `json:"id"`
}

func NewCommentExtend(comment Comment, id string) *CommentExtend {
	return &CommentExtend{
		Comment: comment,
		ID:      id,
	}
}

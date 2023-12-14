package entity

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrAlreadyUpvote   = errors.New("already upvote")
	ErrAlreadyDownvote = errors.New("already downvote")
	ErrAlreadyUnvote   = errors.New("already unvote")
)

type Vote struct {
	UserID string `json:"user" bson:"user"`
	Vote   int8   `json:"vote" bson:"vote"`
}

// pure enterprise business logic
type Post struct {
	Score            int64           `json:"score" bson:"score"`
	Views            uint64          `json:"views" bson:"views"`
	Type             string          `json:"type" bson:"type"`
	Title            string          `json:"title" bson:"title"`
	Author           Author          `json:"author" bson:"author"`
	Category         string          `json:"category" bson:"category"`
	Text             string          `json:"text,omitempty" bson:"text"`
	URL              string          `json:"url,omitempty" bson:"url"`
	Votes            []Vote          `json:"votes" bson:"votes"`
	Comments         []CommentExtend `json:"comments" bson:"comments"`
	CreatedTime      time.Time       `json:"created" bson:"created"`
	UpvotePercentage uint8           `json:"upvotePercentage" bson:"upvotePercentage"`
}

func (post *Post) View() Post {
	post.Views++
	return *post
}

func (post *Post) Upvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == 1 {
				return fmt.Errorf("[Post.Upvote]: %w", ErrAlreadyUpvote)
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{
				UserID: userID,
				Vote:   +1,
			}
			post.Score += 2
			post.updateScore()
			return nil
		}
	}

	post.Votes = append(post.Votes, Vote{
		UserID: userID,
		Vote:   1,
	})
	post.Score++
	post.updateScore()

	return nil

}

func (post *Post) Downvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == -1 {
				return fmt.Errorf("[Post.Downvote]: %w", ErrAlreadyDownvote)
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{
				UserID: userID,
				Vote:   -1,
			}
			post.Score -= 2
			post.updateScore()

			return nil
		}
	}

	post.Votes = append(post.Votes, Vote{
		UserID: userID,
		Vote:   -1,
	})
	post.Score--
	post.updateScore()

	return nil
}

func (post *Post) Unvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == 1 {
				post.Score--
			} else {
				post.Score++
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{}
			post.Votes = post.Votes[:len(post.Votes)-1]
			post.updateScore()

			return nil
		}
	}

	return fmt.Errorf("[Post.Unvote]: %w", ErrAlreadyUnvote)
}

func (post *Post) updateScore() {
	if post.Score > 0 {
		post.UpvotePercentage = uint8(float64((post.Score+int64(len(post.Votes)))/2) * 100 / float64(len(post.Votes)))
	} else {
		post.UpvotePercentage = 0
	}
}

// business logic with implementation logic
type PostExtend struct {
	Post `bson:"post"`
	ID   string `json:"id" bson:"id"`
}

func NewPostExtend(post Post, id string) *PostExtend {
	return &PostExtend{
		Post: post,
		ID:   id,
	}
}

func (post *PostExtend) GetPost() Post {
	return post.Post
}

func (post *PostExtend) SetPost(newPost Post) {
	post.Post = newPost
}

func (post *PostExtend) GetID() string {
	return post.ID
}

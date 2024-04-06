package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes" gorm:"default=0"`
	PostId    int       `json:"post_id"`
}

func (comment *Comment) Save() (*Comment, error) {
	err := Database.Create(comment).Error
	if err != nil {
		return &Comment{}, err
	}

	return comment, nil
}

func (comment *Comment) Delete() (*Comment, error) {
	err := Database.Delete(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// below function is different from posts as fetching all comments for a single post
// unsure what data type argument should be, was thinking int8 as per table structure
// but Users FindUserById function uses string instead
func FetchAllCommentsByPostId(post_id int) (*[]Comment, error) {
	var comments []Comment
	err := Database.Where("post_id = ?", post_id).Find(&comments).Error

	fmt.Println(comments)

	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func FetchSpecificComment(post_id int, comment_id int) (*Comment, error) {
	var comment Comment
	err := Database.Where("post_id = ? AND id = ?", post_id, comment_id).First(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

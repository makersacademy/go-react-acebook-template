package models

import (
	"time"

    "gorm.io/gorm"
)

type Post struct {
	gorm.Model // gorm.Model creates the following common fields automatically; ID (unit / gorm:"primaryKey"), CreatedAt (time.Time), UpdatedAt(time.Time), DeletedAt (gorm.DeletedAt / gorm:"index")
	UserID    string    `json:"user_id"`
  // ** WE CAN PROBABLY ADD 'CreatedAt' TO JSONPost TO BE ABLE TO USE IN FRONTEND **
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
    Likes     int       `json:"likes"`
//     Filename  *string `json:"image_filename,omitempty"`
//     FileSize  *int64  `json:"image_filesize,omitempty"`
//     FileType  *string `json:"image_filetype,omitempty"`
//     FileData  *[]byte `json:"image_filedata,omitempty"`
}

// This function creates a new record in the database
func (post *Post) Save() (*Post, error) {
    err := Database.Create(post).Error // Database relates to the database connection variable created in the db.go file. Create and Error are part of the functions imported from gorm
    if err != nil {
        return &Post{}, err
    }

    return post, nil
}

func (post *Post) Delete() (*Post, error) {
    err := Database.Delete(post).Error
    if err != nil {
        return &Post{}, err
    }
    return post, nil
}

func (post *Post) SaveLike() (*Post, error) {
    post.Likes++
    err := Database.Save(post).Error
    if err != nil {
        return nil, err
    }
    return post, nil
}

// This function retrieves all posts from the database and returns them as a slice
func FetchAllPosts() (*[]Post, error) {
    var posts []Post
    err := Database.Find(&posts).Error // Database.Find(&posts) fetches all records from the database and adds them to the 'posts' slice


	// fmt.Println(posts) // Prints out the slice (likely used for debugging)

    if err != nil {
        return &[]Post{}, err // Returns a pointer to an empty slice of Post structs and the error (if error)
    }

    return &posts, nil // Returns a pointer to the 'posts' slice (if successful)
}

func FetchSpecificPost(postID uint64) (*Post, error) {
    var post Post
    err := Database.First(&post, postID).Error

    if err != nil {
        return nil, err
    }

    return &post, nil
}


package controllers

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

type JSONPost struct {
	ID        uint    `json:"_id"` // Relates to the primary key ID for the posts table
	Message   string  `json:"message"`
	CreatedAt string  `json:"created_at"`
	Likes     int     `json:"likes"`
	Filename  *string `json:"image_filename,omitempty"`
	FileSize  *int64  `json:"image_filesize,omitempty"`
	FileType  *string `json:"image_filetype,omitempty"`
	FileData  *[]byte `json:"image_filedata,omitempty"`
  UserID    string `json:"user_id"`
}

func GetAllPosts(ctx *gin.Context) {
	posts, err := models.FetchAllPosts()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "UserID not found in context"})
		return
	}
	userID := val.(string)
	token, err := auth.GenerateToken(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	var jsonPosts []JSONPost
	for _, post := range *posts {
		var Filename *string
		var FileSize *int64
		var FileType *string
		var FileData *[]byte

		if post.Filename != nil {
			Filename = post.Filename
			FileSize = post.FileSize
			FileType = post.FileType
			FileData = post.FileData
		}

		jsonPost := JSONPost{
			Message:   post.Message,
			ID:        post.ID,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			Likes:     post.Likes,
      UserID:    post.UserID,
			Filename:  Filename,
			FileSize:  FileSize,
			FileType:  FileType,
			FileData:  FileData,
		}

		jsonPosts = append(jsonPosts, jsonPost)
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": jsonPosts, "token": token})
}

func GetSpecificPost(ctx *gin.Context) {
	postIDStr := ctx.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := models.FetchSpecificPost(postID)
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	jsonPost := JSONPost{
		Message:   post.Message,
		ID:        post.ID,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		Likes:     post.Likes,
		Filename:  post.Filename,
		FileSize:  post.FileSize,
		FileType:  post.FileType,
		FileData:  post.FileData,
		UserID:    post.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"post": jsonPost})
}

func CreatePost(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := ctx.Request.FormValue("message")
	if message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Post message empty"})
		return
	}

	file, fileHeader, err := ctx.Request.FormFile("image")
	//defer file.Close()
	if err != nil {
		postTime := time.Now()
		likeCount := 0
		newPost := models.Post{
			Message:   message,
			CreatedAt: postTime,
			Likes:     likeCount,
		}

		_, err = newPost.Save()
		if err != nil {
			SendInternalError(ctx, err)
			return
		}
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	fileName := fileHeader.Filename
	fileSize := fileHeader.Size
	fileType := fileHeader.Header.Get("Content-Type")

	postTime := time.Now()
	likeCount := 0
	newPost := models.Post{
    UserID:    userID.(string), // cast the user ID to a string
		Message:   message,
		CreatedAt: postTime,
		Likes:     likeCount,
		Filename:  &fileName,
		FileSize:  &fileSize,
		FileType:  &fileType,
		FileData:  &fileData,

	PostTime := time.Now()
	// formattedTime := PostTime.Format("2006-01-02 15:04:05")
	LikeCount := 0
	// getting the user id from the gin context and passing an error
	// if there is none
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ERROR": "USER ID NOT FOUND IN CONTEXT"})
		return
	}
	newPost := models.Post{
		UserID:    userID.(string), // cast the user ID to a string
		Message:   requestBody.Message,
		CreatedAt: PostTime,
		Likes:     LikeCount,
	}

	_, err = newPost.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	// val, _ := ctx.Get("userID")
	// userID := val.(string)
	// token, _ := auth.GenerateToken(userID)


	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created"}) //sends confirmation message back if successfully saved
}

// }
//  // if err != nil {
//  //  ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
//  //  return
//  // }

//  // if len(requestBody.Message) == 0 {
//  //  ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
//  //  return
//  // }

//  PostTime := time.Now()
//  // formattedTime := PostTime.Format("2006-01-02 15:04:05")
//  LikeCount := 0

//  newPost := models.Post{
//      Message:   requestBody.Message,
//      CreatedAt: PostTime,
//      Likes:     LikeCount,
//      Filename: newPost.Filename,
//      FileSize: newPost.FileSize,
//      FileType: newPost.FileType,
//      FileData: FileData,
//  }

//  _, err = newPost.Save() // Adds newPost to database
//  if err != nil {
//      SendInternalError(ctx, err)
//      return
//  }

//  val, _ := ctx.Get("userID")
//  userID := val.(string)
//  token, _ := auth.GenerateToken(userID)

//  ctx.JSON(http.StatusCreated, gin.H{"message": "Post created", "token": token}) //sends confirmation message back if successfully saved
// }

func DeletePost(ctx *gin.Context) {
	// Get the post ID from the URL path parameter
	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Fetch the post from the database
	post, err := models.FetchSpecificPost(uint64(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the post is nil
	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Delete post from database
	DeletedPost, err := post.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Post deleted successfully", "deleted post": DeletedPost})
}

func UpdatePostLikes(ctx *gin.Context) {
	// Get the post ID from the URL path parameter
	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Fetch the post from the database
	post, err := models.FetchSpecificPost(uint64(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the post is nil
	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Increment the likes count
	likedPost, err := post.SaveLike()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save like"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Like added successfully", "liked_post": likedPost})
}

// func (postID uint64) DeletePost(ctx *gin.Context) {
//  // var requestBody createPostRequestBody
//  // err := ctx.BindJSON(&requestBody)
//  postToDelete, err := GetSpecificPost(postID)
//      if err != nil {
//      ctx.JSON(http.StatusBadRequest, gin.H{"deletion error": err})
//      return
//  }
//  if err := Database.Delete(postToDelete).Error; err != nil {
//     return err
//  }
//  return nil
// }

//  if err != nil {
//      ctx.JSON(http.StatusBadRequest, gin.H{"deletion error": err})
//      return
//  }

//  if err := Database.Delete(postToDelete).Error; err != nil {
//      return err
//  }
//  return nil

// if len(requestBody.Message) == 0 {
//  ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
//  return
// }

// PostTime := time.Now()
// // formattedTime := PostTime.Format("2006-01-02 15:04:05")
// LikeCount := 0
// newPost := models.Post{
//  Message:   requestBody.Message,
//  CreatedAt: PostTime,
//  Likes:     LikeCount,
// }

// _, err = newPost.Save()
// if err != nil {
//  SendInternalError(ctx, err)
//  return
// }

// val, _ := ctx.Get("userID")
// userID := val.(string)
// token, _ := auth.GenerateToken(userID)

// ctx.JSON(http.StatusCreated, gin.H{"message": "Post deleted", "token": token})
// }

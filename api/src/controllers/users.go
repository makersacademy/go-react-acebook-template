package controllers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

func CreateUser(ctx *gin.Context) {
	var newUser models.User // Creates a variable called newUser with the User struct type User{gorm.Model(id,...), email, password}
	// err := ctx.ShouldBindJSON(&newUser) // Parses the JSON from the request and attempts to match the fields to the newUser fields

	// ERROR HANDLING for ShouldBindJSON below

	// if err != nil {
	// 	if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
	// 		fieldName := jsonErr.Field
	// 		errorMsg := fmt.Sprintf("Invalid value for field '%s': %v", fieldName, jsonErr.Error())
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
	// 	} else {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	}
	// 	return
	// }

	// The below block reads the image data from the request where
	// the content-type is set to multipart/form-data (in Headers)

	file, header, err := ctx.Request.FormFile("image")
	// image is the key in the Postman form-data POST request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	// Read file data
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser = models.User{
		// Update user fields with file information
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
		Username: ctx.PostForm("username"),
		Filename: header.Filename,
		FileSize: header.Size,
		FileType: header.Header.Get("Content-Type"),
		FileData: fileBytes,
	}

	if newUser.Email == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Must supply username and password"}) // Returns error if email and password are blank
		return
	}

	if len(newUser.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters"})
		return
	}

	var specialCharacters = []string{
		"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "'", "\"", "<", ">", ",", ".", "?", "/",
	}

	var containsSpecialCharacter = false
	for _, char := range newUser.Password {
		for _, specialChar := range specialCharacters {
			if string(char) == specialChar {
				containsSpecialCharacter = true
			}
		}
	}

	if containsSpecialCharacter != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password must have at least one special character"})
		return
	}

	if newUser.Email[0] == '@' {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	if !strings.Contains(newUser.Email, "@") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	if strings.Count(newUser.Email, "@") > 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	if strings.Contains(newUser.Email, " ") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	existingUser, err := models.FindUserByEmail(newUser.Email)
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	if existingUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// existingUser, err := models.FindUserByEmail(newUser.Email)
	// if err != nil {
	// 	SendInternalError(ctx, err)
	// 	return
	// }

	// if existingUser != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
	// 	return
	// }

	_, err = newUser.Save() // Adds newUser to database

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	userID := string(newUser.ID)
	token, _ := auth.GenerateToken(userID)

	ctx.JSON(http.StatusCreated, gin.H{"message": "OK", "token": token}) //sends confirmation message back if successfully saved
}

func GetUser(ctx *gin.Context) {
	// userID := ctx.Param("id") // This is to check response in postman

	// The below two lines of code are to extract userID from token when that functionality becomes possible
	// val, _ := ctx.Get("userID")
	// userID := val.(string)

	userID := "18" // hardcoded for frontend testing until userID can be extracted from token
	token, _ := auth.GenerateToken(userID)
	user, err := models.FindUser(userID)
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

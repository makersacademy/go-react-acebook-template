package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

type CreateTokenRequestBody struct {
	Email    string
	Password string
}

func CreateToken(ctx *gin.Context) {
	// This function essentially checks if the user's login details are in the database, and issues a token for their login if they are.
	// Is a handler function for a specific route in Gin.
	// gin.Context (ctx) is a request-level data object that carries request details, path parameters, headers, and other metadata.
	var input CreateTokenRequestBody
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fmt.Println(input)

	// Using the makers package, a model called FindUserByEmail
	// (which is called from the user.go in models)
	// checks if the given email from ctx is ok (perhaps if it's not already in the database?)
	// and sends an internal error, but one that is not broadcast

	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		SendInternalError(ctx, err)
	}
	// The next if section starts given there is no error from the previous and checks if the password matches their email in the database.
	// It will give an error if not
	if user.Password != input.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"})
		return
	}

	// If these two conditions are met, then the function will generate a token that allows them to login, and sends an internal error if not.
	token, err := auth.GenerateToken(string(user.ID))
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	ctx.Set("userID", string(user.ID))
	// The last line communicates this to the frontend and verifies that it's ok
	ctx.JSON(http.StatusCreated, gin.H{"token": token, "message": "OK"})
}

package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
)

// * Only one function again. Is taking the given token and authenticating it.
// * Designed to be executed before the actual route handler.

func AuthenticationMiddleware(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")[7:]
	// Get token string from Authorization header
	fmt.Println(tokenString)

	// Decode and validate token
	token, err := auth.DecodeToken(tokenString)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "auth error"})
		return
	}

	// Store user ID in context for later use
	ctx.Set("userID", token.UserID)
	// Pass control to the next middleware or route handler
	ctx.Next()
}

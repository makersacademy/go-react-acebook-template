package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/controllers"
)

func setupAuthenticationRoutes(baseRouter *gin.RouterGroup) {
	tokensRouter := baseRouter.Group("/tokens")

	tokensRouter.POST("", controllers.CreateToken)
}

// This code sets up a route group "/tokens" using the Gin web framework.
// It registers a POST route "/tokens" mapped to the "controllers.CreateToken" handler.
// When a POST request is sent to "/tokens":
//   - "controllers.CreateToken" verifies the provided email and password.
//   - If valid, it generates a JSON Web Token (JWT) using "auth.GenerateToken".
//   - The JWT is returned to the client.
// JWT generation in "auth.GenerateToken":
//   - Creates a new JWT with user ID, issued time, and expiration time claims.
//   - Signs the JWT with a secret key.
// Before reaching "controllers.CreateToken":
//   - "AuthenticationMiddleware" extracts the JWT from the "Authorization" header.
//   - It decodes and validates the JWT using "auth.DecodeToken".
//   - If valid, the user ID is stored in the Gin context.
//   - If invalid, an unauthorized error is returned.
// "auth.DecodeToken":
//   - Parses and verifies the JWT using the secret key.
//   - Extracts claims (user ID, issued time, expiration time) from the JWT.
//   - Returns an "AuthToken" struct with claims or an error.
// Flow:
//   1. Client sends POST "/tokens" with email and password.
//   2. "controllers.CreateToken" is invoked, generates JWT, and returns it.
//   3. Client includes JWT in "Authorization" header for subsequent requests.
//   4. "AuthenticationMiddleware" validates JWT and stores user ID in context.
//   5. Route handlers can access the authenticated user ID from the context.

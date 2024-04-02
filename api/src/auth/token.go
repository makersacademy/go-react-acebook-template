package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthToken struct {
	ExpiresAt time.Time
	IssuedAt  time.Time
	UserID    string
}

func GenerateToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	now := time.Now()
	/*
	 * Creates a token variable that creates a new webtoken, signs it, and maps its three values with a userID,
	 * a time and date which will be now, and an expiration time with now plus 10 minutes
	 */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": jwt.NewNumericDate(now),
		"exp": jwt.NewNumericDate(now.Add(time.Minute * 1)),
	})
	// This is then returned and signed into a string of bytes, and into our secret environment variable
	return token.SignedString([]byte(secret))
}

func DecodeToken(tokenString string) (AuthToken, error) {
	secret := os.Getenv("JWT_SECRET")
	/*
	 * First two lines take the code and parse it using the jwt package. Within it,
	 * there is a pointer to the jwt.Token, presumably to pull this function and apply it to the tokenString.
	 */
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return AuthToken{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AuthToken{}, fmt.Errorf("something is wrong")
	}

	return createAuthTokenFromClaims(claims)
}

func (authToken AuthToken) IsValid() bool {
	now := time.Now()
	// Basically checks if the AuthToken is valid by verifying the time it was
	// created is within the parameters set earlier,
	// that the time it was issued is not after now, and before the expiry time
	return now.After(authToken.IssuedAt) && now.Before(authToken.ExpiresAt)
}

func createAuthTokenFromClaims(claims jwt.MapClaims) (AuthToken, error) {
	/*
	 * From the previously issued AuthToken from claims in the decode function,
	 * this function is what powers the ending of the Decode Token and so is only callable from within this package
	 */
	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return AuthToken{}, err
	}

	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return AuthToken{}, err
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return AuthToken{}, err
	}

	// Function checks the expiration time, when it was issued, the user id,
	// using jwt.MapClaims to do so, and returning errors if there are any.
	// If there are no errors, it will return the AuthToken in the format of the
	// AuthToken struct
	return AuthToken{
		ExpiresAt: expiresAt.Time,
		IssuedAt:  issuedAt.Time,
		UserID:    userID,
	}, nil
}

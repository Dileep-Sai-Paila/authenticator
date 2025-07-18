package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("naa_key_naa_istam") // jwtSecretKey is the secret key used to sign the JWT tokens.

// in this struct, this will be encoded to a JWT,I will add my own custom claims here, as well as embedding the standard claims.
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// generates a new JWT for a given email.
func GenerateJWT(email string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour) //setting the expiration time for the token.

	// Create the JWT claims, which includes the user's email and the expiration time.
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //creating a new token object, specifying the signing method and the claims.

	// signing the token with our secret key to get the complete, signed token string.
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// to parse and validates a JWT string.
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		// just to handle errors like an expired token or a token that is not yet valid.
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	// to check if the token is valid.
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil // Return the claims if the token is valid.
}

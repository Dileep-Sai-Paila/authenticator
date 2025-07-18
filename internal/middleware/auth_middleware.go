package middleware

import (
	"authenticator/pkg/utils"
	"context"
	"net/http"
)

// my custom type for the context key inorderr to avoid collisions
type contextKey string

// the key that I used to store user's email in the request context
const userEmailKey = contextKey("userEmail")

// It checks for a valid JWT in the cookies, so that it protects routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1.)reading the "token" cookie from the request.
		cookie, err := r.Cookie("token")
		if err != nil {
			// if the cookie is not found, the user is not authenticated
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
				return
			}
			//for any other error
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// 2.)validating the JWT from the cookie.
		tokenStr := cookie.Value
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			// if the token is invalid, then access is denied
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// 3.)If the token is valid, it should pass the user's email to the next handler
		// to do this I added the email to the request's context
		ctx := context.WithValue(r.Context(), userEmailKey, claims.Email)

		// inorder to call the next handler in the chain, passing the updated request with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// function for handlers to safely retrieve the email that the middleware stored in the context
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userEmailKey).(string)
	return email, ok
}

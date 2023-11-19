package security

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/togglhire/backend-homework/types"
)

type contextKey string

const userIDContextKey contextKey = "USER_ID"

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		tokenString := extractTokenFromHeader(request)
		if tokenString != "" {
			auth, err := validateToken(tokenString)
			if err == nil && auth != nil {
				ctx = context.WithValue(ctx, userIDContextKey, &auth.UserID)
				*request = *request.WithContext(ctx)
			}
		}

		// Pass the user ID to the context for later use
		next.ServeHTTP(w, request.WithContext(ctx))
	})
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

func validateToken(token string) (*types.JWTClaims, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	claims := types.JWTClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey != "" {
		return secretKey
	}

	return "superSecretKey"
}

// GetIdentity get authentication info according to context
func GetUserID(ctx context.Context) *int {
	if ctx != nil {
		if userID := ctx.Value(userIDContextKey); userID != nil {
			return userID.(*int)
		}
	}

	return nil
}

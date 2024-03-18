package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// TODO: move to environment
			return []byte("BknHBOeY3j7lxUKYK8TUHDsx5J0KUqnIt81TeHllIrY="), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if time.Now().Before(claims["exp"].(time.Time)) {
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}

		ctx := context.WithValue(r.Context(), "role", claims["role"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

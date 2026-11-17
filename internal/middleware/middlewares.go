package middleware

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/auth"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.GetBearerToken(&r.Header)
		if err != nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}

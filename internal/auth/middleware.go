package auth

import (
	"assets/internal/api"
	"context"
	"net/http"
	"strings"
)

func Middleware(next http.HandlerFunc, serviceAuth ServiceAuth) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := strings.TrimSpace(
			strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1),
		)

		session, err := serviceAuth.GetActiveSessionByToken(r.Context(), token)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		next(w, r.WithContext(context.WithValue(r.Context(), "uid", session.UId)))
	})
}

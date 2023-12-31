package api

import (
	"context"
	"finger-print-voting-backend/internal/api/auth"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (srv *Server) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Autherization")

		log.Println("autherisation header: ", token)

		token = strings.TrimSpace(token)

		if !strings.HasPrefix(token, "Bearer {") {
			WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid autherization header"))
			return
		}

		token = strings.TrimPrefix(token, "Bearer {")

		if !strings.HasSuffix(token, "}") {
			WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid autherization header"))
			return
		}

		token = strings.TrimSuffix(token, "}")

		claims, err := auth.GetClaims(srv.passwordSecret, token)
		if err != nil {
			WriteError(w, http.StatusUnauthorized, fmt.Errorf("failed to get claims"))
			return
		}

		user, err := srv.db.GetUser(claims.Username)
		if err != nil {
			WriteError(w, http.StatusUnauthorized, fmt.Errorf("failed to get user details"))
			return
		}

		ctx := context.WithValue(r.Context(), types.UserContext, types.UserDetails{}.FromUser(user))
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})

}

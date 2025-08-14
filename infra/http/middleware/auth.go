package httpapi

import (
	"context"
	"net/http"
	"strings"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type ctxKey string

const actorKey ctxKey = "actorID"

func RequireAuth(tokens contract.TokenIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			raw := r.Header.Get("Authorization")
			if !strings.HasPrefix(raw, "Bearer ") {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}
			tok := strings.TrimPrefix(raw, "Bearer ")
			claims, err := tokens.ParseAndVerify(r.Context(), tok)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			sub, _ := claims["sub"].(string)
			if sub == "" {
				http.Error(w, "invalid subject", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), actorKey, domain.AccountID(sub))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func actorID(r *http.Request) domain.AccountID {
	if v := r.Context().Value(actorKey); v != nil {
		return v.(domain.AccountID)
	}
	return ""
}

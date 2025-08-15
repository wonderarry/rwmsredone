package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wonderarry/rwmsredone/infra/http/httputils"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

func RequireAuth(tokens contract.TokenIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			raw := r.Header.Get("Authorization")
			if !strings.HasPrefix(raw, "Bearer ") {
				httputils.ErrorJSON(w, http.StatusUnauthorized, ErrMissingToken)
				return
			}
			claims, err := tokens.ParseAndVerify(r.Context(), strings.TrimPrefix(raw, "Bearer "))
			if err != nil {
				httputils.ErrorJSON(w, http.StatusUnauthorized, err)
				return
			}
			sub, _ := claims["sub"].(string)
			if sub == "" {
				httputils.ErrorJSON(w, http.StatusUnauthorized, ErrInvalidSubject)
				return
			}
			ctx := httputils.WithActor(r.Context(), domain.AccountID(sub))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var (
	ErrMissingToken   = fmt.Errorf("missing bearer token")
	ErrInvalidSubject = fmt.Errorf("invalid subject")
)

package v1

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prok05/gophermart/internal/entity"
	jwt2 "github.com/prok05/gophermart/pkg/jwt"
	"net/http"
	"strings"
)

func (h *V1) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			h.unauthorizedResponse(w, r, fmt.Errorf("authorization header is missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			h.unauthorizedResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}

		token := parts[1]

		jwtToken, err := jwt2.ValidateToken(
			token,
			h.cfg.App.Name,
			h.cfg.App.Name,
			h.cfg.JWT.Secret,
		)
		if err != nil {
			h.unauthorizedResponse(w, r, err)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)

		userID := fmt.Sprint(claims["sub"])

		ctx := r.Context()

		user, err := h.u.GetByID(ctx, userID)
		if err != nil {
			h.unauthorizedResponse(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, entity.ContextUserID, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

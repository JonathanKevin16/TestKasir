package user

import (
	"TestKasir/internal/domain/auth/service"
	"github.com/go-chi/chi"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func ProvideAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Router(r chi.Router) {

	r.Route("/cashier", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.RegisterCashier)
			r.Post("/login", h.LoginCashier)
		})
	})
}

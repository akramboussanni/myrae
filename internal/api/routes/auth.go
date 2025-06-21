package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handleRegister)
	r.Post("/login", handleLogin)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/me", handleProfile)
		r.Post("/logout", handleLogout)
	})

	return r
}

func handleRegister(w http.ResponseWriter, r *http.Request) {}
func handleLogin(w http.ResponseWriter, r *http.Request)    {}
func handleProfile(w http.ResponseWriter, r *http.Request)  {}
func handleLogout(w http.ResponseWriter, r *http.Request)   {}

func authMiddleware(next http.Handler) http.Handler { return next }

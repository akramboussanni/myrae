package routes

import (
	"net/http"

	"github.com/akramboussanni/myrae/internal/api/routes/auth"
	"github.com/akramboussanni/myrae/internal/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(repos *repo.Repos) http.Handler {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("myrae.io"))
	})

	r.Mount("/api/auth", auth.NewAuthRouter(repos.User, repos.Token))

	return r
}

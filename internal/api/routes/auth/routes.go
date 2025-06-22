package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/akramboussanni/myrae/internal/model"
	"github.com/akramboussanni/myrae/internal/repo"
	"github.com/go-chi/chi/v5"
)

type AuthRouter struct {
	userRepo *repo.UserRepo
}

func NewAuthRouter(userRepo *repo.UserRepo) http.Handler {
	ar := &AuthRouter{userRepo: userRepo}
	r := chi.NewRouter()

	r.Post("/register", ar.handleRegister)
	r.Post("/login", ar.handleLogin)

	r.Group(func(r chi.Router) {
		r.Use(ar.authMiddleware)
		r.Get("/me", ar.handleProfile)
		r.Post("/logout", ar.handleLogout)
	})

	return r
}

func (ar *AuthRouter) handleRegister(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if !IsValidEmail(cred.Email) || !IsValidPassword(cred.Password) {
		http.Error(w, "invalid email or password (password must be 8 characters + 1 numeric)", http.StatusBadRequest)
		return
	}

	duplicate, err := ar.userRepo.DuplicateName(r.Context(), cred.Username)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if duplicate {
		http.Error(w, "that username is taken", http.StatusBadRequest)
		return
	}

	hash, err := HashPassword(cred.Password)
	user := &model.User{Username: cred.Username, PasswordHash: hash, Email: cred.Email, CreatedAt: time.Now()}
	if err := ar.userRepo.Create(r.Context(), user); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ar *AuthRouter) handleLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse login, verify credentials, issue JWT/session
}

func (ar *AuthRouter) handleProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Return user info from context/session
}

func (ar *AuthRouter) handleLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Invalidate session/JWT
}

func (ar *AuthRouter) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Verify auth token, load user, put user in context
		next.ServeHTTP(w, r)
	})
}

package repo

import (
	"context"
	"database/sql"

	"github.com/akramboussanni/myrae/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, u *model.User) error {
	// TODO: insert user into DB
	return nil
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (*model.User, error) {
	// TODO: fetch user by id
	return nil, nil
}

func (r *UserRepo) DuplicateName(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	return exists, err
}

func (r *UserRepo) DuplicateEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	return exists, err
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: fetch user by email
	return nil, nil
}

func (r *UserRepo) Update(ctx context.Context, u *model.User) error {
	// TODO: update user in DB
	return nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	// TODO: delete user by id
	return nil
}

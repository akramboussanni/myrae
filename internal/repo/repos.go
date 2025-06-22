package repo

import "database/sql"

type Repos struct {
	User  *UserRepo
	Roles *RoleRepo
	// Add other repos here
}

func NewRepos(db *sql.DB) *Repos {
	return &Repos{
		User:  NewUserRepo(db),
		Roles: NewRoleRepo(db),
	}
}

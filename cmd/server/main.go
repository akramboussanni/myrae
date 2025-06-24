package main

import (
	"log"
	"net/http"

	"github.com/akramboussanni/myrae/config"
	"github.com/akramboussanni/myrae/internal/api/routes"
	"github.com/akramboussanni/myrae/internal/api/routes/auth"
	"github.com/akramboussanni/myrae/internal/db"
	"github.com/akramboussanni/myrae/internal/repo"
)

func main() {
	config.Init()
	err := auth.InitSnowflake(1)
	if err != nil {
		panic(err)
	}

	db.Init("todo")
	db.RunMigrations("./migrations")

	repos := repo.NewRepos(db.DB)
	r := routes.SetupRouter(repos)

	log.Println("server will run @ localhost:9520")
	http.ListenAndServe(":9520", r)
}

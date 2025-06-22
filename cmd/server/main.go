package main

import (
	"log"
	"net/http"

	"github.com/akramboussanni/myrae/config"
	"github.com/akramboussanni/myrae/internal/api/routes"
	"github.com/akramboussanni/myrae/internal/db"
	"github.com/akramboussanni/myrae/internal/repo"
)

var jwtSecret []byte

func main() {
	config.Init()

	db.Init("todo")
	repos := repo.NewRepos(db.DB)
	r := routes.SetupRouter(repos)

	log.Println("server will run @ localhost:9520")
	http.ListenAndServe(":9520", r)
}

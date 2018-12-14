package main

import (
	"github.com/IvanAlekseevichPopov/movieApi/config"
	"github.com/IvanAlekseevichPopov/movieApi/config/db"
	"github.com/IvanAlekseevichPopov/movieApi/controller"
	"github.com/IvanAlekseevichPopov/movieApi/entity"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db.Conn.AutoMigrate(entity.User{}, entity.Staff{}) //TODO only separate migrations
	router := mux.NewRouter()

	controller.NewApi(router, "/api")
	controller.NewAdmin(router, "/admin")

	log.Printf("Listening on %s", config.Config.Port)
	http.ListenAndServe(config.Config.Port, router)
}

package controller

import (
	"github.com/IvanAlekseevichPopov/movieApi/config/db"
	"github.com/IvanAlekseevichPopov/movieApi/entity"
	"github.com/gorilla/mux"
	"github.com/qor/admin"
	"net/http"
)

func NewAdmin(router *mux.Router, route string) {
	Admin := admin.New(&admin.AdminConfig{DB: db.Conn})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&entity.User{})
	Admin.AddResource(&entity.Staff{})
	Admin.AddResource(&entity.Movie{})

	server := http.NewServeMux()
	router.PathPrefix(route).Handler(server)

	// Mount admin interface to mux
	Admin.MountTo(route, server)
}

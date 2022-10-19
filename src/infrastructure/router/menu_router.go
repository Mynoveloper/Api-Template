package router

import (
	"github.com/gorilla/mux"

	"github.com/Mynor2397/Api-Template/src/app/controller/menu"
	"github.com/Mynor2397/Api-Template/src/domain/service"
	"github.com/Mynor2397/Api-Template/src/infrastructure/repository"
)

func (r *router) menuRoutes(router *mux.Router) *mux.Router {

	menuRepo := repository.NewMenuRepository(r.database)
	menuService := service.NewMenuService(menuRepo)
	menuController := menu.NewMenuController(r.logger, r.httpResponse, menuService)

	menus := router.PathPrefix("/menus").Subrouter()

	menus.HandleFunc("", menuController.NewMenu).Methods("POST")
	menus.HandleFunc("", menuController.GetAll).Methods("GET")

	return router
}

package router

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynoveloper/logger"

	"github.com/Mynor2397/Mongo-Quick-Start/internal/app/controller"
	"github.com/Mynor2397/Mongo-Quick-Start/internal/app/controller/menu"
	"github.com/Mynor2397/Mongo-Quick-Start/internal/domain/service"
	"github.com/Mynor2397/Mongo-Quick-Start/internal/infrastructure/repository"
)

func menuRoutes(logger logger.ILogger,
	httpResponse controller.IHttpResponse,
	database *mongo.Database,
	router *mux.Router,
) *mux.Router {

	menuRepo := repository.NewMenuRepository(database)
	menuService := service.NewMenuService(menuRepo)
	menuController := menu.NewMenuController(logger, httpResponse, menuService)

	menus := router.PathPrefix("/menus").Subrouter()

	menus.HandleFunc("", menuController.NewMenu).Methods("POST")
	menus.HandleFunc("", menuController.GetAll).Methods("GET")

	return router
}

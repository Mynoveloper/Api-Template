package router

import (
	"github.com/Mynoveloper/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynor2397/Api-Template/src/app/controller"
	"github.com/Mynor2397/Api-Template/src/app/middleware"
)

type router struct {
	logger       logger.ILogger
	database     *mongo.Database
	httpResponse controller.HttpResponse
}

type Router interface {
	InitRoutes() *mux.Router
	menuRoutes(router *mux.Router) *mux.Router
}

func NewRouter(logger logger.ILogger, database *mongo.Database) Router {
	return &router{
		logger:       logger,
		database:     database,
		httpResponse: controller.NewResponse(logger),
	}
}

func (r *router) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1/").Subrouter()
	api = r.menuRoutes(api)

	router.Use(middleware.HttpJson)
	return router
}

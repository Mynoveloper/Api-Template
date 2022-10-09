package router

import (
	"github.com/Mynoveloper/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynor2397/Mongo-Quick-Start/internal/app/controller"
	"github.com/Mynor2397/Mongo-Quick-Start/internal/app/middleware"
)

func InitRoutes(logger logger.ILogger, database *mongo.Database) *mux.Router {
	router := mux.NewRouter()
	httpResponse := controller.NewResponse(logger)

	api := router.PathPrefix("/api/v1/").Subrouter()
	api = menuRoutes(logger, httpResponse, database, api)

	router.Use(middleware.HttpJson)
	return router
}

package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Mynoveloper/logger"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynor2397/Api-Template/internal/app/router"
	"github.com/Mynor2397/Api-Template/internal/infrastructure/utils"
)

type httpServe struct {
	logger   logger.ILogger
	database *mongo.Database
}
type IHttpServe interface {
	ListenAndServe()
	GetPort() string
}

func NewServer(logger logger.ILogger, database *mongo.Database) IHttpServe {
	return &httpServe{
		logger:   logger,
		database: database,
	}
}
func (s *httpServe) GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	return port
}

func (s *httpServe) OpenBrowser() {
	if os.Getenv("APP_ENVIRONMENT") == "development" {
		utils.OpenBrowser("http://localhost:" + s.GetPort())
	}
}

func (s *httpServe) ListenAndServe() {

	router := router.InitRoutes(s.logger, s.database)

	s.logger.Info("Listen And serve on port: ", s.GetPort())

	s.OpenBrowser()
	http.ListenAndServe(fmt.Sprintf(":%s", s.GetPort()), router)
}

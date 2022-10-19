package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Mynoveloper/logger"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynor2397/Api-Template/src/infrastructure/router"
	"github.com/Mynor2397/Api-Template/src/infrastructure/utils"
)

type httpServe struct {
	logger   logger.ILogger
	database *mongo.Database
}

type HttpServe interface {
	ListenAndServe()
	getPort() string
}

func NewServer(logger logger.ILogger, database *mongo.Database) HttpServe {
	return &httpServe{
		logger:   logger,
		database: database,
	}
}
func (s *httpServe) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	return port
}

func (s *httpServe) OpenBrowser() {
	if os.Getenv("APP_ENVIRONMENT") == "development" {
		utils.OpenBrowser("http://localhost:" + s.getPort())
	}
}

func (s *httpServe) ListenAndServe() {

	router := router.NewRouter(s.logger, s.database)

	s.logger.Info("Listen And serve on port: ", s.getPort())

	s.OpenBrowser()

	err := http.ListenAndServe(fmt.Sprintf(":%s", s.getPort()), router.InitRoutes())
	if err != nil {
		s.logger.Fatal(err)
	}
}

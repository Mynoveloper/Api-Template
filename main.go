package main

import (
	"os"

	"github.com/Mynoveloper/logger"
	"github.com/joho/godotenv"

	"github.com/Mynor2397/Api-Template/internal/cmd/server"
	"github.com/Mynor2397/Api-Template/internal/infrastructure/database"
)

func main() {
	logger := logger.NewLogger(logger.LoggerInfo{
		ProgramName: "ApiTemplate",
		Level:       "debug",
		JSONOutput:  true,
	})

	if os.Getenv("APP_ENVIRONMENT") == "development" {
		if err := godotenv.Load(); err != nil {
			logger.Fatal("No .env file found")
		}
	}

	database := database.Connection(logger)
	http := server.NewServer(logger, database)

	http.ListenAndServe()
}

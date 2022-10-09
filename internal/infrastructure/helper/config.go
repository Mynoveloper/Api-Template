package helper

import (
	"log"
	"os"
	"strconv"

	"github.com/Mynor2397/Mongo-Quick-Start/internal/domain/model"
)

// Config del servidor
func Config() model.Config {

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	log.Println(dbPort)
	return model.Config{
		ApplicationPort: os.Getenv("APPLICATION_PORT"),
		Hostdb:          os.Getenv("DB_HOST"),
		Portdb:          dbPort,
		Database:        os.Getenv("DATABASE"),
	}
}

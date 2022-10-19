package helper

import (
	"os"
	"strconv"

	"github.com/Mynor2397/Api-Template/src/domain/model"
)

// Config: load the server configuration
func Config() model.Config {

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return model.Config{
		ApplicationPort: os.Getenv("APPLICATION_PORT"),
		Hostdb:          os.Getenv("DB_HOST"),
		Portdb:          dbPort,
		Database:        os.Getenv("DATABASE"),
	}
}

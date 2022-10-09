package database

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Mynor2397/Mongo-Quick-Start/internal/infrastructure/helper"
	"github.com/Mynoveloper/logger"
)

var (
	once     sync.Once
	database *mongo.Database
)

// Connection return a conection to mongodb
func Connection(logger logger.ILogger) *mongo.Database {
	c := helper.Config()
	logger.Info(c)
	once.Do(func() {

		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", c.Hostdb, c.Portdb))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("Connected mongodb successfull at ", c.Hostdb, c.Portdb)
		database = client.Database(c.Database)
	})

	return database
}

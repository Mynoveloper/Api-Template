package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/Mynoveloper/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Mynor2397/Api-Template/src/infrastructure/helper"
)

var (
	once     sync.Once
	database *mongo.Database
)

// Connection return a new conection to specific mongodb database
func Connection(logger logger.ILogger) *mongo.Database {
	c := helper.Config()
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

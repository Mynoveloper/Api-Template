package repository

import (
	"context"

	"github.com/Mynoveloper/mongopaginator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Mynor2397/Api-Template/internal/domain/model"
)

type menuRepository struct {
	database *mongo.Database
}

type IMenuRepository interface {
	NewMenu() error
	GetAll(contex context.Context) (*model.Menus, error)
}

func NewMenuRepository(database *mongo.Database) IMenuRepository {
	return &menuRepository{
		database: database,
	}
}

func (r *menuRepository) NewMenu() error {

	return nil
}

func (r *menuRepository) GetAll(contex context.Context) (*model.Menus, error) {
	menus := model.Menus{}
	menu := model.Menu{}

	collection := r.database.Collection("menu")

	result, err := mongopaginator.New(collection).Limit(15).Page(0).Aggregate()

	if err != nil {
		return &menus, err
	}

	for _, row := range result.Data {
		if err := bson.Unmarshal(row, &menu); err != nil {
			return &menus, err
		}

		menus = append(menus, menu)
	}

	return &menus, nil
}

package service

import (
	"context"

	"github.com/Mynor2397/Api-Template/internal/domain/model"
	"github.com/Mynor2397/Api-Template/internal/infrastructure/repository"
)

type menuService struct{}

type IMenuService interface {
	NewMenu(context context.Context, menu *model.Menu) (string, error)
	GetAll(contex context.Context) (*model.Menus, error)
}

var _menuRepository repository.IMenuRepository

func NewMenuService(menuRepository repository.IMenuRepository) IMenuService {
	_menuRepository = menuRepository

	return &menuService{}
}

func (*menuService) NewMenu(context context.Context, menu *model.Menu) (string, error) {

	return "", nil
}

func (*menuService) GetAll(contex context.Context) (*model.Menus, error) {

	return _menuRepository.GetAll(contex)
}

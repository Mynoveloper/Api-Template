package service

import (
	"context"

	"github.com/Mynor2397/Api-Template/src/domain/model"
	"github.com/Mynor2397/Api-Template/src/infrastructure/repository"
)

type menuService struct {
	menuRepository repository.MenuRepository
}

type MenuService interface {
	NewMenu(context context.Context, menu *model.Menu) (string, error)
	GetAll(contex context.Context) (*model.Menus, error)
}

func NewMenuService(menuRepository repository.MenuRepository) MenuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) NewMenu(context context.Context, menu *model.Menu) (string, error) {
	// TODO:

	return "", nil
}

func (s *menuService) GetAll(contex context.Context) (*model.Menus, error) {

	return s.menuRepository.GetAll(contex)
}

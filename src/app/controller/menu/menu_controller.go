package menu

import (
	"net/http"

	"github.com/Mynor2397/Api-Template/src/app/controller"
	"github.com/Mynor2397/Api-Template/src/domain/model"
	"github.com/Mynor2397/Api-Template/src/domain/service"
	"github.com/Mynoveloper/logger"
)

type menuController struct {
	logger      logger.ILogger
	response    controller.HttpResponse
	menuService service.MenuService
}
type IMenuController interface {
	NewMenu(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewMenuController(logger logger.ILogger, response controller.HttpResponse, menuService service.MenuService) IMenuController {
	return &menuController{
		logger:      logger,
		response:    response,
		menuService: menuService,
	}
}

func (c *menuController) NewMenu(w http.ResponseWriter, r *http.Request) {
	c.response.Respond(w, model.Response{
		Ok:   true,
		Data: 12,
	}, http.StatusOK)

}

func (c *menuController) GetAll(w http.ResponseWriter, r *http.Request) {
	menus, err := c.menuService.GetAll(r.Context())

	if err != nil {
		c.response.RespondError(w, err)
		return
	}

	c.response.Respond(w, model.Response{
		Ok:   true,
		Data: menus,
	}, http.StatusOK)
}

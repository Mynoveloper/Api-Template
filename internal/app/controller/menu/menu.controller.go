package menu

import (
	"net/http"

	"github.com/Mynor2397/Api-Template/internal/app/controller"
	"github.com/Mynor2397/Api-Template/internal/domain/model"
	"github.com/Mynor2397/Api-Template/internal/domain/service"
	"github.com/Mynoveloper/logger"
)

var _logger logger.ILogger
var _response controller.IHttpResponse
var _menuService service.IMenuService

type menuController struct{}
type IMenuController interface {
	NewMenu(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewMenuController(logger logger.ILogger, response controller.IHttpResponse, menuService service.IMenuService) IMenuController {
	_logger = logger
	_response = response
	_menuService = menuService

	return &menuController{}
}

func (c *menuController) NewMenu(w http.ResponseWriter, r *http.Request) {
	_response.Respond(w, model.Response{
		Ok:   true,
		Data: 12,
	}, http.StatusOK)

}

func (c *menuController) GetAll(w http.ResponseWriter, r *http.Request) {
	menus, err := _menuService.GetAll(r.Context())

	if err != nil {
		_response.RespondError(w, err)
		return
	}

	_response.Respond(w, model.Response{
		Ok:   true,
		Data: menus,
	}, http.StatusOK)
}

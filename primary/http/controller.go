package http

import (
	"github.com/labstack/echo/v4"
	"github.com/legoulart/poc-go/internal/domain/model"
	"github.com/legoulart/poc-go/internal/domain/port/primary"
	"github.com/legoulart/poc-go/primary/http/request"
	"net/http"
)

const Path = "/v1/user"

type Controller struct {
	router echo.Group
	port   primary.Port
}

func NewController(router echo.Group, port primary.Port) Controller {
	return Controller{router: router, port: port}
}

func (c Controller) Route() {
	c.router.POST("", func(ctx echo.Context) error {
		var req request.Request
		_ = ctx.Bind(req)

		c.port.Create(*model.NewUser(req.Name, req.Email))

		return ctx.JSON(http.StatusOK, nil)
	})
}

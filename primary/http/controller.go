package http

import (
	"github.com/labstack/echo/v4"
	user "github.com/legoulart/poc-go/internal/domain"
	"github.com/legoulart/poc-go/internal/port"
	"net/http"
)

const Path = "/v1/subscribe"

type Controller struct {
	router echo.Group
	port   port.Port
}

func NewController(router echo.Group, port port.Port) Controller {
	return Controller{router: router, port: port}
}

func (c Controller) Route() {
	c.router.POST("", func(ctx echo.Context) error {
		var req Request
		_ = ctx.Bind(req)

		c.port.Create(*user.NewSubscription(req.Email))

		return ctx.JSON(http.StatusOK, nil)
	})
}

package order

import (
	"github.com/adryanchiko/x-order/service/order-app/pkg/registry"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/labstack/echo/v4"
)

type service struct {
	config *settings.Settings
}

func (ss *service) RegisterRoutes(router *echo.Group) {
	router.GET("/order", ss.fetch)
}

func New(config *settings.Settings) registry.Router {
	return &service{
		config: config,
	}
}

func init() {
	registry.RegisterServiceFactory(New)
}

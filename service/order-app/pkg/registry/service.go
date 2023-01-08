package registry

import (
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"

	"github.com/labstack/echo/v4"
)

// Router is an interface to register router handlers to base router
type Router interface {
	RegisterRoutes(base *echo.Group)
}

type ServiceFactory func(*settings.Settings) Router

var _services []ServiceFactory

func RegisterServiceFactory(factory ServiceFactory) {
	_services = append(_services, factory)
}

func ServiceFactories() []ServiceFactory {
	return _services
}

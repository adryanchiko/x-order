package order

import (
	"net/http"

	"github.com/adryanchiko/x-order/service/order-app/lib/helper"
	"github.com/adryanchiko/x-order/service/order-app/pkg/orderitem"
	"github.com/labstack/echo/v4"
)

func (ss *service) fetch(c echo.Context) error {
	store := orderitem.NewOrderItemEnt()
	res, err := store.Find(c.Request().Context(), orderitem.Criteria{
		Find: helper.Find{
			Keyword: "",
			Skip:    0,
			Limit:   5,
		},
	})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

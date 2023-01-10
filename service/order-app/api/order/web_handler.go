package order

import (
	"net/http"
	"time"

	"github.com/adryanchiko/x-order/service/order-app/lib/helper"
	"github.com/adryanchiko/x-order/service/order-app/pkg/orderitem"
	"github.com/labstack/echo/v4"
)

type (
	fetchQuery struct {
		Keyword   string `query:"keyword"`
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
		Skip      int    `query:"skip"`
	}
)

func (ss *service) fetch(c echo.Context) error {
	var q fetchQuery
	if err := c.Bind(&q); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	criteria := orderitem.Criteria{
		Find: helper.Find{
			Keyword: q.Keyword,
			Skip:    q.Skip,
			Limit:   5,
		},
	}

	if q.StartDate != "" {
		layout := "2006-01-02T15:04:05.000Z"
		startDate, err := time.Parse(layout, q.StartDate)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		criteria.Find.From = &startDate
	}

	if q.EndDate != "" {
		layout := "2006-01-02T15:04:05.000Z"
		endDate, err := time.Parse(layout, q.EndDate)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		criteria.Find.To = &endDate
	}

	store := orderitem.NewOrderItemEnt()
	res, err := store.Find(c.Request().Context(), criteria)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (ss *service) fetchAmount(c echo.Context) error {
	var q fetchQuery
	if err := c.Bind(&q); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	criteria := orderitem.Criteria{
		Find: helper.Find{
			Keyword: q.Keyword,
		},
	}

	if q.StartDate != "" {
		layout := "2006-01-02T15:04:05.000Z"
		startDate, err := time.Parse(layout, q.StartDate)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		criteria.Find.From = &startDate
	}

	if q.EndDate != "" {
		layout := "2006-01-02T15:04:05.000Z"
		endDate, err := time.Parse(layout, q.EndDate)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		criteria.Find.To = &endDate
	}

	store := orderitem.NewOrderItemEnt()
	res, err := store.TotalAmount(c.Request().Context(), criteria)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

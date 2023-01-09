package orderitem

import (
	"context"
	"time"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/lib/helper"
)

type (
	NewOrderItem struct {
		ID           int    `json:"id,omitempty"`
		PricePerUnit int    `json:"price_per_unit"`
		Quantity     int    `json:"quantity"`
		Product      string `json:"product"`
		OrderID      int    `json:"order_id,omitempty"`
	}

	RecordOrderItem struct {
		ID              int       `json:"id,omitempty"`
		PricePerUnit    int       `json:"price_per_unit"`
		Quantity        int       `json:"quantity"`
		Product         string    `json:"product"`
		CreatedAt       time.Time `json:"created_at"`
		OrderName       string    `json:"order_name"`
		CustomerName    string    `json:"customer_name"`
		CustomerCompany string    `json:"customer_company"`
		DeliveredAmount int       `json:"delivered_amount"`
		TotalAmount     int       `json:"total_amount"`
	}

	Criteria struct {
		helper.Find
		IsOrder bool `json:"is_order"`
	}

	SearchResult struct {
		Records []*RecordOrderItem `json:"records"`
		helper.Pagination
	}

	Store interface {
		Create(context.Context, *NewOrderItem) (int, error)
		BulkCreate(ctx context.Context, companies []NewOrderItem) error
		Find(ctx context.Context, criteria Criteria) (*SearchResult, error)
		TotalAmount(ctx context.Context, criteria Criteria) (int, error)
	}
)

func (r *RecordOrderItem) FromEnt(data *ent.OrderItem) *RecordOrderItem {
	// Order Item
	r.ID = data.ID
	r.PricePerUnit = data.PricePerUnit
	r.Quantity = data.Quantity
	r.Product = data.Product

	// Order
	r.CreatedAt = data.Edges.Order.CreatedAt
	r.OrderName = data.Edges.Order.OrderName

	// Customer
	r.CustomerName = data.Edges.Order.Edges.Customer.Name
	r.CustomerCompany = data.Edges.Order.Edges.Customer.Edges.Company.CompanyName

	deliveredQuantity := 0
	for _, d := range data.Edges.Deliveries {
		deliveredQuantity += d.DeliveredQuantity
	}

	// Amount
	r.DeliveredAmount = r.PricePerUnit * deliveredQuantity
	r.TotalAmount = r.PricePerUnit * r.Quantity

	return r
}

package orderitem

import (
	"context"
)

type (
	NewOrderItem struct {
		ID           int    `json:"id,omitempty"`
		PricePerUnit int    `json:"price_per_unit"`
		Quantity     int    `json:"quantity"`
		Product      string `json:"product"`
		OrderID      int    `json:"order_id,omitempty"`
	}

	Store interface {
		Create(context.Context, *NewOrderItem) (int, error)
		BulkCreate(ctx context.Context, companies []NewOrderItem) error
	}
)

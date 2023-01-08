package delivery

import (
	"context"
)

type (
	NewDelivery struct {
		ID                int `json:"id,omitempty"`
		DeliveredQuantity int `json:"delivered_quantity"`
		OrderItemID       int `json:"order_item_id,omitempty"`
	}

	Store interface {
		Create(context.Context, *NewDelivery) (int, error)
		BulkCreate(ctx context.Context, companies []NewDelivery) error
	}
)

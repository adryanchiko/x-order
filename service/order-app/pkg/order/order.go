package order

import (
	"context"
	"time"
)

type (
	NewOrder struct {
		ID         int       `json:"id,omitempty"`
		CreatedAt  time.Time `json:"created_at"`
		OrderName  string    `json:"order_name"`
		CustomerID string    `json:"customer_id,omitempty"`
	}

	Store interface {
		Create(context.Context, *NewOrder) (int, error)
		BulkCreate(ctx context.Context, companies []NewOrder) error
	}
)

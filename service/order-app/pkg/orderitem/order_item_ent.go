package orderitem

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

type orderItemEnt struct{}

func (c *orderItemEnt) Create(ctx context.Context, data *NewOrderItem) (int, error) {
	prepared := entsql.DB().
		OrderItem.
		Create().
		SetPricePerUnit(data.PricePerUnit).
		SetQuantity(data.Quantity).
		SetProduct(data.Product)

	if data.ID != -1 {
		prepared.SetID(data.ID)
	}

	if data.OrderID != -1 {
		prepared.SetOrderID(data.OrderID)
	}

	result, err := prepared.
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return result.ID, nil
}

func (c *orderItemEnt) BulkCreate(ctx context.Context, companies []NewOrderItem) error {
	err := entsql.WithTx(ctx, entsql.DB(), func(tx *ent.Tx) error {
		for _, data := range companies {
			prepared := tx.
				OrderItem.
				Create().
				SetPricePerUnit(data.PricePerUnit).
				SetQuantity(data.Quantity).
				SetProduct(data.Product)

			if data.ID != -1 {
				prepared.SetID(data.ID)
			}

			if data.OrderID != -1 {
				prepared.SetOrderID(data.OrderID)
			}

			if _, err := prepared.Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func NewOrderItemEnt() Store {
	return &orderItemEnt{}
}

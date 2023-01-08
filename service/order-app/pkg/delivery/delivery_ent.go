package delivery

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

type deliveryEnt struct{}

func (c *deliveryEnt) Create(ctx context.Context, data *NewDelivery) (int, error) {
	prepared := entsql.DB().
		Delivery.
		Create().
		SetDeliveredQuantity(data.DeliveredQuantity)

	if data.ID != -1 {
		prepared.SetID(data.ID)
	}

	if data.OrderItemID != -1 {
		prepared.SetOrderItemID(data.OrderItemID)
	}

	result, err := prepared.
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return result.ID, nil
}

func (c *deliveryEnt) BulkCreate(ctx context.Context, companies []NewDelivery) error {
	err := entsql.WithTx(ctx, entsql.DB(), func(tx *ent.Tx) error {
		for _, data := range companies {
			prepared := tx.
				Delivery.
				Create().
				SetDeliveredQuantity(data.DeliveredQuantity)

			if data.ID != -1 {
				prepared.SetID(data.ID)
			}

			if data.OrderItemID != -1 {
				prepared.SetOrderItemID(data.OrderItemID)
			}

			if _, err := prepared.Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func NewDeliveryEnt() Store {
	return &deliveryEnt{}
}

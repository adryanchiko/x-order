package order

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

type orderEnt struct{}

func (c *orderEnt) Create(ctx context.Context, data *NewOrder) (int, error) {
	prepared := entsql.DB().
		Order.
		Create().
		SetCreatedAt(data.CreatedAt).
		SetOrderName(data.OrderName)

	if data.ID != -1 {
		prepared.SetID(data.ID)
	}

	if data.CustomerID != "" {
		prepared.SetCustomerID(data.CustomerID)
	}

	result, err := prepared.
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return result.ID, nil
}

func (c *orderEnt) BulkCreate(ctx context.Context, companies []NewOrder) error {
	err := entsql.WithTx(ctx, entsql.DB(), func(tx *ent.Tx) error {
		for _, data := range companies {
			prepared := tx.
				Order.
				Create().
				SetCreatedAt(data.CreatedAt).
				SetOrderName(data.OrderName)

			if data.ID != -1 {
				prepared.SetID(data.ID)
			}

			if data.CustomerID != "" {
				prepared.SetCustomerID(data.CustomerID)
			}

			if _, err := prepared.Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func NewOrderEnt() Store {
	return &orderEnt{}
}

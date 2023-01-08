package customer

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

type customerEnt struct{}

func (c *customerEnt) Create(ctx context.Context, data *NewCustomer) (string, error) {
	prepared := entsql.DB().
		Customer.
		Create().
		SetLogin(data.Login).
		SetPassword(data.Password).
		SetName(data.Name).
		SetCreditCards(data.CreditCards)

	if data.ID != "" {
		prepared.SetID(data.ID)
	}

	if data.CompanyID != -1 {
		prepared.SetCompanyID(data.CompanyID)
	}

	result, err := prepared.
		Save(ctx)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func (c *customerEnt) BulkCreate(ctx context.Context, companies []NewCustomer) error {
	err := entsql.WithTx(ctx, entsql.DB(), func(tx *ent.Tx) error {
		for _, data := range companies {
			prepared := tx.
				Customer.
				Create().
				SetLogin(data.Login).
				SetPassword(data.Password).
				SetName(data.Name).
				SetCreditCards(data.CreditCards)

			if data.ID != "" {
				prepared.SetID(data.ID)
			}

			if data.CompanyID != -1 {
				prepared.SetCompanyID(data.CompanyID)
			}

			if _, err := prepared.Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func NewCustomerEnt() Store {
	return &customerEnt{}
}

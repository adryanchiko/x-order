package company

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

type companyEnt struct{}

func (c *companyEnt) Create(ctx context.Context, data *NewCompany) (int, error) {
	prepared := entsql.DB().
		Company.
		Create().
		SetCompanyName(data.CompanyName)

	if data.ID != -1 {
		prepared.SetID(data.ID)
	}

	result, err := prepared.
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return result.ID, nil
}

func (c *companyEnt) BulkCreate(ctx context.Context, companies []NewCompany) error {
	err := entsql.WithTx(ctx, entsql.DB(), func(tx *ent.Tx) error {
		for _, data := range companies {
			prepared := tx.
				Company.
				Create().
				SetCompanyName(data.CompanyName)

			if data.ID != -1 {
				prepared.SetID(data.ID)
			}

			if _, err := prepared.Save(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func NewCompanyEnt() Store {
	return &companyEnt{}
}

package company

import "context"

type (
	NewCompany struct {
		ID          int    `json:"id,omitempty"`
		CompanyName string `json:"company_name"`
	}

	Store interface {
		Create(context.Context, *NewCompany) (int, error)
		BulkCreate(context.Context, []NewCompany) error
	}
)

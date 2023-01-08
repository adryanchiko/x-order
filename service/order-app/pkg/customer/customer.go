package customer

import "context"

type (
	NewCustomer struct {
		ID          string   `json:"id,omitempty"`
		Login       string   `json:"login"`
		Password    string   `json:"password"`
		Name        string   `json:"name"`
		CreditCards []string `json:"credit_cards"`
		CompanyID   int      `json:"company_id,omitempty"`
	}

	Store interface {
		Create(context.Context, *NewCustomer) (string, error)
		BulkCreate(ctx context.Context, companies []NewCustomer) error
	}
)

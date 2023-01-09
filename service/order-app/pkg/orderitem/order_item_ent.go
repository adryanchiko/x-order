package orderitem

import (
	"context"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/ent/company"
	"github.com/adryanchiko/x-order/service/order-app/ent/customer"
	"github.com/adryanchiko/x-order/service/order-app/ent/order"
	"github.com/adryanchiko/x-order/service/order-app/ent/orderitem"
	"github.com/adryanchiko/x-order/service/order-app/ent/predicate"
	"github.com/adryanchiko/x-order/service/order-app/lib/helper"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
)

const defaultLimit = 5

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

func (c *orderItemEnt) Find(ctx context.Context, criteria Criteria) (*SearchResult, error) {
	var result SearchResult
	predicates := []predicate.OrderItem{}

	if criteria.Keyword != "" {
		predicates = append(predicates,
			orderitem.Or(
				orderitem.ProductContainsFold(criteria.Keyword),
				orderitem.HasOrderWith(
					predicate.Order(
						order.Or(
							order.OrderNameContainsFold(criteria.Keyword),
							order.HasCustomerWith(
								predicate.Customer(
									customer.Or(
										customer.NameContainsFold(criteria.Keyword),
										customer.HasCompanyWith(
											predicate.Company(company.CompanyNameContainsFold(criteria.Keyword)),
										),
									),
								),
							),
						),
					),
				),
			),
		)
	}

	if criteria.From != nil {
		predicates = append(predicates, predicate.OrderItem(order.CreatedAtGT(*criteria.From)))
	}
	if criteria.To != nil {
		predicates = append(predicates, predicate.OrderItem(order.CreatedAtLT(*criteria.To)))
	}

	if criteria.Limit <= 0 {
		criteria.Limit = defaultLimit
	}

	filterQuery := entsql.DB().
		OrderItem.
		Query().
		Where(predicates...).
		WithOrder(func(q *ent.OrderQuery) {
			q.WithCustomer(func(q *ent.CustomerQuery) {
				q.WithCompany()
			})
		}).
		WithDeliveries()

	if err := helper.WithPaginationData(
		ctx, &result.Pagination, filterQuery, criteria.Find,
	); err != nil {
		return nil, err
	}

	rows, err := filterQuery.
		Offset(criteria.Skip).
		Limit(criteria.Limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range rows {
		result.Records = append(result.Records, new(RecordOrderItem).FromEnt(r))
	}

	return &result, nil
}

func (c *orderItemEnt) TotalAmount(ctx context.Context, criteria Criteria) (int, error) {
	predicates := []predicate.OrderItem{}

	if criteria.Keyword != "" {
		predicates = append(predicates,
			orderitem.Or(
				orderitem.ProductContainsFold(criteria.Keyword),
				orderitem.HasOrderWith(
					predicate.Order(
						order.Or(
							order.OrderNameContainsFold(criteria.Keyword),
							order.HasCustomerWith(
								predicate.Customer(
									customer.Or(
										customer.NameContainsFold(criteria.Keyword),
										customer.HasCompanyWith(
											predicate.Company(company.CompanyNameContainsFold(criteria.Keyword)),
										),
									),
								),
							),
						),
					),
				),
			),
		)
	}

	if criteria.From != nil {
		predicates = append(predicates, predicate.OrderItem(order.CreatedAtGT(*criteria.From)))
	}
	if criteria.To != nil {
		predicates = append(predicates, predicate.OrderItem(order.CreatedAtLT(*criteria.To)))
	}

	if criteria.Limit <= 0 {
		criteria.Limit = defaultLimit
	}

	rows, err := entsql.DB().
		OrderItem.
		Query().
		Where(predicates...).
		WithOrder().
		All(ctx)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, r := range rows {
		total += r.PricePerUnit * r.Quantity
	}

	return total, nil
}

func NewOrderItemEnt() Store {
	return &orderItemEnt{}
}

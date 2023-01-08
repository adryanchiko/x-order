// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/adryanchiko/x-order/service/order-app/ent/order"
	"github.com/adryanchiko/x-order/service/order-app/ent/orderitem"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PricePerUnit holds the value of the "price_per_unit" field.
	PricePerUnit int `json:"price_per_unit,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Product holds the value of the "product" field.
	Product string `json:"product,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges             OrderItemEdges `json:"edges"`
	order_order_items *int
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// Deliveries holds the value of the deliveries edge.
	Deliveries []*Delivery `json:"deliveries,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// DeliveriesOrErr returns the Deliveries value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) DeliveriesOrErr() ([]*Delivery, error) {
	if e.loadedTypes[1] {
		return e.Deliveries, nil
	}
	return nil, &NotLoadedError{edge: "deliveries"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID, orderitem.FieldPricePerUnit, orderitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldProduct:
			values[i] = new(sql.NullString)
		case orderitem.ForeignKeys[0]: // order_order_items
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderitem.FieldPricePerUnit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price_per_unit", values[i])
			} else if value.Valid {
				oi.PricePerUnit = int(value.Int64)
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = int(value.Int64)
			}
		case orderitem.FieldProduct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product", values[i])
			} else if value.Valid {
				oi.Product = value.String
			}
		case orderitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_order_items", value)
			} else if value.Valid {
				oi.order_order_items = new(int)
				*oi.order_order_items = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrder queries the "order" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrder() *OrderQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrder(oi)
}

// QueryDeliveries queries the "deliveries" edge of the OrderItem entity.
func (oi *OrderItem) QueryDeliveries() *DeliveryQuery {
	return (&OrderItemClient{config: oi.config}).QueryDeliveries(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return (&OrderItemClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("price_per_unit=")
	builder.WriteString(fmt.Sprintf("%v", oi.PricePerUnit))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteString(", ")
	builder.WriteString("product=")
	builder.WriteString(oi.Product)
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem

func (oi OrderItems) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}

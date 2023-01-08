// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/adryanchiko/x-order/service/order-app/ent/delivery"
	"github.com/adryanchiko/x-order/service/order-app/ent/orderitem"
)

// Delivery is the model entity for the Delivery schema.
type Delivery struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DeliveredQuantity holds the value of the "delivered_quantity" field.
	DeliveredQuantity int `json:"delivered_quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeliveryQuery when eager-loading is set.
	Edges                 DeliveryEdges `json:"edges"`
	order_item_deliveries *int
}

// DeliveryEdges holds the relations/edges for other nodes in the graph.
type DeliveryEdges struct {
	// OrderItem holds the value of the order_item edge.
	OrderItem *OrderItem `json:"order_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderItemOrErr returns the OrderItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryEdges) OrderItemOrErr() (*OrderItem, error) {
	if e.loadedTypes[0] {
		if e.OrderItem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orderitem.Label}
		}
		return e.OrderItem, nil
	}
	return nil, &NotLoadedError{edge: "order_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Delivery) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case delivery.FieldID, delivery.FieldDeliveredQuantity:
			values[i] = new(sql.NullInt64)
		case delivery.ForeignKeys[0]: // order_item_deliveries
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Delivery", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Delivery fields.
func (d *Delivery) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case delivery.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case delivery.FieldDeliveredQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delivered_quantity", values[i])
			} else if value.Valid {
				d.DeliveredQuantity = int(value.Int64)
			}
		case delivery.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_item_deliveries", value)
			} else if value.Valid {
				d.order_item_deliveries = new(int)
				*d.order_item_deliveries = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrderItem queries the "order_item" edge of the Delivery entity.
func (d *Delivery) QueryOrderItem() *OrderItemQuery {
	return (&DeliveryClient{config: d.config}).QueryOrderItem(d)
}

// Update returns a builder for updating this Delivery.
// Note that you need to call Delivery.Unwrap() before calling this method if this Delivery
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Delivery) Update() *DeliveryUpdateOne {
	return (&DeliveryClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Delivery entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Delivery) Unwrap() *Delivery {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Delivery is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Delivery) String() string {
	var builder strings.Builder
	builder.WriteString("Delivery(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("delivered_quantity=")
	builder.WriteString(fmt.Sprintf("%v", d.DeliveredQuantity))
	builder.WriteByte(')')
	return builder.String()
}

// Deliveries is a parsable slice of Delivery.
type Deliveries []*Delivery

func (d Deliveries) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}

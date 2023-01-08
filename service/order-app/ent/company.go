// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/adryanchiko/x-order/service/order-app/ent/company"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CompanyName holds the value of the "company_name" field.
	CompanyName string `json:"company_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges CompanyEdges `json:"edges"`
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// Customers holds the value of the customers edge.
	Customers []*Customer `json:"customers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CustomersOrErr returns the Customers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) CustomersOrErr() ([]*Customer, error) {
	if e.loadedTypes[0] {
		return e.Customers, nil
	}
	return nil, &NotLoadedError{edge: "customers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			values[i] = new(sql.NullInt64)
		case company.FieldCompanyName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Company", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case company.FieldCompanyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company_name", values[i])
			} else if value.Valid {
				c.CompanyName = value.String
			}
		}
	}
	return nil
}

// QueryCustomers queries the "customers" edge of the Company entity.
func (c *Company) QueryCustomers() *CustomerQuery {
	return (&CompanyClient{config: c.config}).QueryCustomers(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return (&CompanyClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("company_name=")
	builder.WriteString(c.CompanyName)
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company

func (c Companies) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
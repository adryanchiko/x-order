// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/adryanchiko/x-order/service/order-app/ent/company"
	"github.com/adryanchiko/x-order/service/order-app/ent/customer"
	"github.com/adryanchiko/x-order/service/order-app/ent/order"
	"github.com/adryanchiko/x-order/service/order-app/ent/predicate"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks     []Hook
	mutation  *CustomerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetLogin sets the "login" field.
func (cu *CustomerUpdate) SetLogin(s string) *CustomerUpdate {
	cu.mutation.SetLogin(s)
	return cu
}

// SetPassword sets the "password" field.
func (cu *CustomerUpdate) SetPassword(s string) *CustomerUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetCreditCards sets the "credit_cards" field.
func (cu *CustomerUpdate) SetCreditCards(s []string) *CustomerUpdate {
	cu.mutation.SetCreditCards(s)
	return cu
}

// AppendCreditCards appends s to the "credit_cards" field.
func (cu *CustomerUpdate) AppendCreditCards(s []string) *CustomerUpdate {
	cu.mutation.AppendCreditCards(s)
	return cu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cu *CustomerUpdate) SetCompanyID(id int) *CustomerUpdate {
	cu.mutation.SetCompanyID(id)
	return cu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCompanyID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetCompanyID(*id)
	}
	return cu
}

// SetCompany sets the "company" edge to the Company entity.
func (cu *CustomerUpdate) SetCompany(c *Company) *CustomerUpdate {
	return cu.SetCompanyID(c.ID)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (cu *CustomerUpdate) AddOrderIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddOrderIDs(ids...)
	return cu
}

// AddOrders adds the "orders" edges to the Order entity.
func (cu *CustomerUpdate) AddOrders(o ...*Order) *CustomerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.AddOrderIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (cu *CustomerUpdate) ClearCompany() *CustomerUpdate {
	cu.mutation.ClearCompany()
	return cu
}

// ClearOrders clears all "orders" edges to the Order entity.
func (cu *CustomerUpdate) ClearOrders() *CustomerUpdate {
	cu.mutation.ClearOrders()
	return cu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (cu *CustomerUpdate) RemoveOrderIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveOrderIDs(ids...)
	return cu
}

// RemoveOrders removes "orders" edges to Order entities.
func (cu *CustomerUpdate) RemoveOrders(o ...*Order) *CustomerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.RemoveOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CustomerMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CustomerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomerUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Login(); ok {
		_spec.SetField(customer.FieldLogin, field.TypeString, value)
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.CreditCards(); ok {
		_spec.SetField(customer.FieldCreditCards, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedCreditCards(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, customer.FieldCreditCards, value)
		})
	}
	if cu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !cu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CustomerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetLogin sets the "login" field.
func (cuo *CustomerUpdateOne) SetLogin(s string) *CustomerUpdateOne {
	cuo.mutation.SetLogin(s)
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CustomerUpdateOne) SetPassword(s string) *CustomerUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetCreditCards sets the "credit_cards" field.
func (cuo *CustomerUpdateOne) SetCreditCards(s []string) *CustomerUpdateOne {
	cuo.mutation.SetCreditCards(s)
	return cuo
}

// AppendCreditCards appends s to the "credit_cards" field.
func (cuo *CustomerUpdateOne) AppendCreditCards(s []string) *CustomerUpdateOne {
	cuo.mutation.AppendCreditCards(s)
	return cuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cuo *CustomerUpdateOne) SetCompanyID(id int) *CustomerUpdateOne {
	cuo.mutation.SetCompanyID(id)
	return cuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCompanyID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetCompanyID(*id)
	}
	return cuo
}

// SetCompany sets the "company" edge to the Company entity.
func (cuo *CustomerUpdateOne) SetCompany(c *Company) *CustomerUpdateOne {
	return cuo.SetCompanyID(c.ID)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (cuo *CustomerUpdateOne) AddOrderIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddOrderIDs(ids...)
	return cuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (cuo *CustomerUpdateOne) AddOrders(o ...*Order) *CustomerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.AddOrderIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (cuo *CustomerUpdateOne) ClearCompany() *CustomerUpdateOne {
	cuo.mutation.ClearCompany()
	return cuo
}

// ClearOrders clears all "orders" edges to the Order entity.
func (cuo *CustomerUpdateOne) ClearOrders() *CustomerUpdateOne {
	cuo.mutation.ClearOrders()
	return cuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (cuo *CustomerUpdateOne) RemoveOrderIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveOrderIDs(ids...)
	return cuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (cuo *CustomerUpdateOne) RemoveOrders(o ...*Order) *CustomerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.RemoveOrderIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	return withHooks[*Customer, CustomerMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CustomerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomerUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Login(); ok {
		_spec.SetField(customer.FieldLogin, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CreditCards(); ok {
		_spec.SetField(customer.FieldCreditCards, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedCreditCards(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, customer.FieldCreditCards, value)
		})
	}
	if cuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !cuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}

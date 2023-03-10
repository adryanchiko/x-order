// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adryanchiko/x-order/service/order-app/ent/company"
	"github.com/adryanchiko/x-order/service/order-app/ent/customer"
	"github.com/adryanchiko/x-order/service/order-app/ent/predicate"
)

// CompanyUpdate is the builder for updating Company entities.
type CompanyUpdate struct {
	config
	hooks     []Hook
	mutation  *CompanyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cu *CompanyUpdate) Where(ps ...predicate.Company) *CompanyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCompanyName sets the "company_name" field.
func (cu *CompanyUpdate) SetCompanyName(s string) *CompanyUpdate {
	cu.mutation.SetCompanyName(s)
	return cu
}

// AddCustomerIDs adds the "customers" edge to the Customer entity by IDs.
func (cu *CompanyUpdate) AddCustomerIDs(ids ...string) *CompanyUpdate {
	cu.mutation.AddCustomerIDs(ids...)
	return cu
}

// AddCustomers adds the "customers" edges to the Customer entity.
func (cu *CompanyUpdate) AddCustomers(c ...*Customer) *CompanyUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCustomerIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cu *CompanyUpdate) Mutation() *CompanyMutation {
	return cu.mutation
}

// ClearCustomers clears all "customers" edges to the Customer entity.
func (cu *CompanyUpdate) ClearCustomers() *CompanyUpdate {
	cu.mutation.ClearCustomers()
	return cu
}

// RemoveCustomerIDs removes the "customers" edge to Customer entities by IDs.
func (cu *CompanyUpdate) RemoveCustomerIDs(ids ...string) *CompanyUpdate {
	cu.mutation.RemoveCustomerIDs(ids...)
	return cu
}

// RemoveCustomers removes "customers" edges to Customer entities.
func (cu *CompanyUpdate) RemoveCustomers(c ...*Customer) *CompanyUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCustomerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompanyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CompanyMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompanyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompanyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CompanyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompanyUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
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
	if value, ok := cu.mutation.CompanyName(); ok {
		_spec.SetField(company.FieldCompanyName, field.TypeString, value)
	}
	if cu.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCustomersIDs(); len(nodes) > 0 && !cu.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
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
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CompanyUpdateOne is the builder for updating a single Company entity.
type CompanyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CompanyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCompanyName sets the "company_name" field.
func (cuo *CompanyUpdateOne) SetCompanyName(s string) *CompanyUpdateOne {
	cuo.mutation.SetCompanyName(s)
	return cuo
}

// AddCustomerIDs adds the "customers" edge to the Customer entity by IDs.
func (cuo *CompanyUpdateOne) AddCustomerIDs(ids ...string) *CompanyUpdateOne {
	cuo.mutation.AddCustomerIDs(ids...)
	return cuo
}

// AddCustomers adds the "customers" edges to the Customer entity.
func (cuo *CompanyUpdateOne) AddCustomers(c ...*Customer) *CompanyUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCustomerIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cuo *CompanyUpdateOne) Mutation() *CompanyMutation {
	return cuo.mutation
}

// ClearCustomers clears all "customers" edges to the Customer entity.
func (cuo *CompanyUpdateOne) ClearCustomers() *CompanyUpdateOne {
	cuo.mutation.ClearCustomers()
	return cuo
}

// RemoveCustomerIDs removes the "customers" edge to Customer entities by IDs.
func (cuo *CompanyUpdateOne) RemoveCustomerIDs(ids ...string) *CompanyUpdateOne {
	cuo.mutation.RemoveCustomerIDs(ids...)
	return cuo
}

// RemoveCustomers removes "customers" edges to Customer entities.
func (cuo *CompanyUpdateOne) RemoveCustomers(c ...*Customer) *CompanyUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCustomerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompanyUpdateOne) Select(field string, fields ...string) *CompanyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Company entity.
func (cuo *CompanyUpdateOne) Save(ctx context.Context) (*Company, error) {
	return withHooks[*Company, CompanyMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompanyUpdateOne) SaveX(ctx context.Context) *Company {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompanyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CompanyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompanyUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CompanyUpdateOne) sqlSave(ctx context.Context) (_node *Company, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Company.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for _, f := range fields {
			if !company.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != company.FieldID {
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
	if value, ok := cuo.mutation.CompanyName(); ok {
		_spec.SetField(company.FieldCompanyName, field.TypeString, value)
	}
	if cuo.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCustomersIDs(); len(nodes) > 0 && !cuo.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CustomersTable,
			Columns: []string{company.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Company{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}

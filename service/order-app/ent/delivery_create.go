// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adryanchiko/x-order/service/order-app/ent/delivery"
	"github.com/adryanchiko/x-order/service/order-app/ent/orderitem"
)

// DeliveryCreate is the builder for creating a Delivery entity.
type DeliveryCreate struct {
	config
	mutation *DeliveryMutation
	hooks    []Hook
}

// SetDeliveredQuantity sets the "delivered_quantity" field.
func (dc *DeliveryCreate) SetDeliveredQuantity(i int) *DeliveryCreate {
	dc.mutation.SetDeliveredQuantity(i)
	return dc
}

// SetID sets the "id" field.
func (dc *DeliveryCreate) SetID(i int) *DeliveryCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (dc *DeliveryCreate) SetOrderItemID(id int) *DeliveryCreate {
	dc.mutation.SetOrderItemID(id)
	return dc
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (dc *DeliveryCreate) SetNillableOrderItemID(id *int) *DeliveryCreate {
	if id != nil {
		dc = dc.SetOrderItemID(*id)
	}
	return dc
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (dc *DeliveryCreate) SetOrderItem(o *OrderItem) *DeliveryCreate {
	return dc.SetOrderItemID(o.ID)
}

// Mutation returns the DeliveryMutation object of the builder.
func (dc *DeliveryCreate) Mutation() *DeliveryMutation {
	return dc.mutation
}

// Save creates the Delivery in the database.
func (dc *DeliveryCreate) Save(ctx context.Context) (*Delivery, error) {
	return withHooks[*Delivery, DeliveryMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeliveryCreate) SaveX(ctx context.Context) *Delivery {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeliveryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeliveryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeliveryCreate) check() error {
	if _, ok := dc.mutation.DeliveredQuantity(); !ok {
		return &ValidationError{Name: "delivered_quantity", err: errors.New(`ent: missing required field "Delivery.delivered_quantity"`)}
	}
	return nil
}

func (dc *DeliveryCreate) sqlSave(ctx context.Context) (*Delivery, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeliveryCreate) createSpec() (*Delivery, *sqlgraph.CreateSpec) {
	var (
		_node = &Delivery{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: delivery.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: delivery.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.DeliveredQuantity(); ok {
		_spec.SetField(delivery.FieldDeliveredQuantity, field.TypeInt, value)
		_node.DeliveredQuantity = value
	}
	if nodes := dc.mutation.OrderItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   delivery.OrderItemTable,
			Columns: []string{delivery.OrderItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_deliveries = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeliveryCreateBulk is the builder for creating many Delivery entities in bulk.
type DeliveryCreateBulk struct {
	config
	builders []*DeliveryCreate
}

// Save creates the Delivery entities in the database.
func (dcb *DeliveryCreateBulk) Save(ctx context.Context) ([]*Delivery, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Delivery, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeliveryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeliveryCreateBulk) SaveX(ctx context.Context) []*Delivery {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeliveryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeliveryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

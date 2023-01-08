package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Delivery holds the schema definition for the Delivery entity.
type Delivery struct {
	ent.Schema
}

// Fields of the Delivery.
func (Delivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("delivered_quantity"),
	}
}

// Edges of the Delivery.
func (Delivery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_item", OrderItem.Type).
			Ref("deliveries").
			Unique(),
	}
}

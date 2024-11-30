package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Policies struct {
	ent.Schema
}

func (Policies) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now()),
	}
}

func (Policies) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type), // A Policies can target roles
		edge.To("users", User.Type), // A Policies can target users
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now()),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type),    // Role has many permissions
		edge.From("users", User.Type).Ref("roles"), // Many users can have this role
	}
}

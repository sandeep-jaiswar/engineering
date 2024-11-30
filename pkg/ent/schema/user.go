package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("hashed_password"),
		field.String("status").Default("active"), // active, disabled
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now()),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type), // One user can have many roles
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("name"),
		field.String("email").Unique(),
		field.String("password"),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

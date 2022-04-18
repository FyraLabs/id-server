package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TOTPMethod holds the schema definition for the TOTPMethod entity.
type TOTPMethod struct {
	ent.Schema
}

// Fields of the TOTPMethod.
func (TOTPMethod) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("secret"),
		field.Time("createdAt").Default(time.Now),
		field.Time("lastUsedAt").Optional().Nillable(),
		field.String("name"),
	}
}

// Edges of the TOTPMethod.
func (TOTPMethod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("totpMethods").
			Unique(),
	}

}

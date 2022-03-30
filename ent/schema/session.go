package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("ip"),
		field.String("userAgent"),
		field.Time("createdAt").Default(time.Now),
		field.Time("lastUsedAt").Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}

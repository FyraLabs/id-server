// Code generated by entc, DO NOT EDIT.

package totpmethod

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fyralabs/id-server/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// LastUsedAt applies equality check predicate on the "lastUsedAt" field. It's identical to LastUsedAtEQ.
func LastUsedAt(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUsedAt), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecret), v))
	})
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecret), v))
	})
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecret), v))
	})
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecret), v))
	})
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecret), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// LastUsedAtEQ applies the EQ predicate on the "lastUsedAt" field.
func LastUsedAtEQ(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtNEQ applies the NEQ predicate on the "lastUsedAt" field.
func LastUsedAtNEQ(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtIn applies the In predicate on the "lastUsedAt" field.
func LastUsedAtIn(vs ...time.Time) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastUsedAt), v...))
	})
}

// LastUsedAtNotIn applies the NotIn predicate on the "lastUsedAt" field.
func LastUsedAtNotIn(vs ...time.Time) predicate.TOTPMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TOTPMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastUsedAt), v...))
	})
}

// LastUsedAtGT applies the GT predicate on the "lastUsedAt" field.
func LastUsedAtGT(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtGTE applies the GTE predicate on the "lastUsedAt" field.
func LastUsedAtGTE(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtLT applies the LT predicate on the "lastUsedAt" field.
func LastUsedAtLT(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtLTE applies the LTE predicate on the "lastUsedAt" field.
func LastUsedAtLTE(v time.Time) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUsedAt), v))
	})
}

// LastUsedAtIsNil applies the IsNil predicate on the "lastUsedAt" field.
func LastUsedAtIsNil() predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastUsedAt)))
	})
}

// LastUsedAtNotNil applies the NotNil predicate on the "lastUsedAt" field.
func LastUsedAtNotNil() predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastUsedAt)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TOTPMethod) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TOTPMethod) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TOTPMethod) predicate.TOTPMethod {
	return predicate.TOTPMethod(func(s *sql.Selector) {
		p(s.Not())
	})
}

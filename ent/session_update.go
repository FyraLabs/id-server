// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fyralabs/id-server/ent/predicate"
	"github.com/fyralabs/id-server/ent/session"
	"github.com/fyralabs/id-server/ent/user"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIP sets the "ip" field.
func (su *SessionUpdate) SetIP(s string) *SessionUpdate {
	su.mutation.SetIP(s)
	return su
}

// SetUserAgent sets the "userAgent" field.
func (su *SessionUpdate) SetUserAgent(s string) *SessionUpdate {
	su.mutation.SetUserAgent(s)
	return su
}

// SetCreatedAt sets the "createdAt" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetLastUsedAt sets the "lastUsedAt" field.
func (su *SessionUpdate) SetLastUsedAt(t time.Time) *SessionUpdate {
	su.mutation.SetLastUsedAt(t)
	return su
}

// SetNillableLastUsedAt sets the "lastUsedAt" field if the given value is not nil.
func (su *SessionUpdate) SetNillableLastUsedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetLastUsedAt(*t)
	}
	return su
}

// ClearLastUsedAt clears the value of the "lastUsedAt" field.
func (su *SessionUpdate) ClearLastUsedAt() *SessionUpdate {
	su.mutation.ClearLastUsedAt()
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SessionUpdate) SetUserID(id uuid.UUID) *SessionUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(id *uuid.UUID) *SessionUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldIP,
		})
	}
	if value, ok := su.mutation.UserAgent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldUserAgent,
		})
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.LastUsedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldLastUsedAt,
		})
	}
	if su.mutation.LastUsedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldLastUsedAt,
		})
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetIP sets the "ip" field.
func (suo *SessionUpdateOne) SetIP(s string) *SessionUpdateOne {
	suo.mutation.SetIP(s)
	return suo
}

// SetUserAgent sets the "userAgent" field.
func (suo *SessionUpdateOne) SetUserAgent(s string) *SessionUpdateOne {
	suo.mutation.SetUserAgent(s)
	return suo
}

// SetCreatedAt sets the "createdAt" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetLastUsedAt sets the "lastUsedAt" field.
func (suo *SessionUpdateOne) SetLastUsedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetLastUsedAt(t)
	return suo
}

// SetNillableLastUsedAt sets the "lastUsedAt" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableLastUsedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetLastUsedAt(*t)
	}
	return suo
}

// ClearLastUsedAt clears the value of the "lastUsedAt" field.
func (suo *SessionUpdateOne) ClearLastUsedAt() *SessionUpdateOne {
	suo.mutation.ClearLastUsedAt()
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SessionUpdateOne) SetUserID(id uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(id *uuid.UUID) *SessionUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldIP,
		})
	}
	if value, ok := suo.mutation.UserAgent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldUserAgent,
		})
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.LastUsedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldLastUsedAt,
		})
	}
	if suo.mutation.LastUsedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldLastUsedAt,
		})
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
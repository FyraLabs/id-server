// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fyralabs/id-server/ent/totpmethod"
	"github.com/fyralabs/id-server/ent/user"
	"github.com/google/uuid"
)

// TOTPMethodCreate is the builder for creating a TOTPMethod entity.
type TOTPMethodCreate struct {
	config
	mutation *TOTPMethodMutation
	hooks    []Hook
}

// SetSecret sets the "secret" field.
func (tmc *TOTPMethodCreate) SetSecret(s string) *TOTPMethodCreate {
	tmc.mutation.SetSecret(s)
	return tmc
}

// SetCreatedAt sets the "createdAt" field.
func (tmc *TOTPMethodCreate) SetCreatedAt(t time.Time) *TOTPMethodCreate {
	tmc.mutation.SetCreatedAt(t)
	return tmc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tmc *TOTPMethodCreate) SetNillableCreatedAt(t *time.Time) *TOTPMethodCreate {
	if t != nil {
		tmc.SetCreatedAt(*t)
	}
	return tmc
}

// SetLastUsedAt sets the "lastUsedAt" field.
func (tmc *TOTPMethodCreate) SetLastUsedAt(t time.Time) *TOTPMethodCreate {
	tmc.mutation.SetLastUsedAt(t)
	return tmc
}

// SetNillableLastUsedAt sets the "lastUsedAt" field if the given value is not nil.
func (tmc *TOTPMethodCreate) SetNillableLastUsedAt(t *time.Time) *TOTPMethodCreate {
	if t != nil {
		tmc.SetLastUsedAt(*t)
	}
	return tmc
}

// SetName sets the "name" field.
func (tmc *TOTPMethodCreate) SetName(s string) *TOTPMethodCreate {
	tmc.mutation.SetName(s)
	return tmc
}

// SetID sets the "id" field.
func (tmc *TOTPMethodCreate) SetID(u uuid.UUID) *TOTPMethodCreate {
	tmc.mutation.SetID(u)
	return tmc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tmc *TOTPMethodCreate) SetUserID(id uuid.UUID) *TOTPMethodCreate {
	tmc.mutation.SetUserID(id)
	return tmc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tmc *TOTPMethodCreate) SetNillableUserID(id *uuid.UUID) *TOTPMethodCreate {
	if id != nil {
		tmc = tmc.SetUserID(*id)
	}
	return tmc
}

// SetUser sets the "user" edge to the User entity.
func (tmc *TOTPMethodCreate) SetUser(u *User) *TOTPMethodCreate {
	return tmc.SetUserID(u.ID)
}

// Mutation returns the TOTPMethodMutation object of the builder.
func (tmc *TOTPMethodCreate) Mutation() *TOTPMethodMutation {
	return tmc.mutation
}

// Save creates the TOTPMethod in the database.
func (tmc *TOTPMethodCreate) Save(ctx context.Context) (*TOTPMethod, error) {
	var (
		err  error
		node *TOTPMethod
	)
	tmc.defaults()
	if len(tmc.hooks) == 0 {
		if err = tmc.check(); err != nil {
			return nil, err
		}
		node, err = tmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TOTPMethodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tmc.check(); err != nil {
				return nil, err
			}
			tmc.mutation = mutation
			if node, err = tmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tmc.hooks) - 1; i >= 0; i-- {
			if tmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tmc *TOTPMethodCreate) SaveX(ctx context.Context) *TOTPMethod {
	v, err := tmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmc *TOTPMethodCreate) Exec(ctx context.Context) error {
	_, err := tmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmc *TOTPMethodCreate) ExecX(ctx context.Context) {
	if err := tmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmc *TOTPMethodCreate) defaults() {
	if _, ok := tmc.mutation.CreatedAt(); !ok {
		v := totpmethod.DefaultCreatedAt()
		tmc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tmc *TOTPMethodCreate) check() error {
	if _, ok := tmc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "TOTPMethod.secret"`)}
	}
	if _, ok := tmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "TOTPMethod.createdAt"`)}
	}
	if _, ok := tmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TOTPMethod.name"`)}
	}
	return nil
}

func (tmc *TOTPMethodCreate) sqlSave(ctx context.Context) (*TOTPMethod, error) {
	_node, _spec := tmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tmc *TOTPMethodCreate) createSpec() (*TOTPMethod, *sqlgraph.CreateSpec) {
	var (
		_node = &TOTPMethod{config: tmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: totpmethod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: totpmethod.FieldID,
			},
		}
	)
	if id, ok := tmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tmc.mutation.Secret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: totpmethod.FieldSecret,
		})
		_node.Secret = value
	}
	if value, ok := tmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: totpmethod.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tmc.mutation.LastUsedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: totpmethod.FieldLastUsedAt,
		})
		_node.LastUsedAt = &value
	}
	if value, ok := tmc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: totpmethod.FieldName,
		})
		_node.Name = value
	}
	if nodes := tmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   totpmethod.UserTable,
			Columns: []string{totpmethod.UserColumn},
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
		_node.user_totp_methods = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TOTPMethodCreateBulk is the builder for creating many TOTPMethod entities in bulk.
type TOTPMethodCreateBulk struct {
	config
	builders []*TOTPMethodCreate
}

// Save creates the TOTPMethod entities in the database.
func (tmcb *TOTPMethodCreateBulk) Save(ctx context.Context) ([]*TOTPMethod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmcb.builders))
	nodes := make([]*TOTPMethod, len(tmcb.builders))
	mutators := make([]Mutator, len(tmcb.builders))
	for i := range tmcb.builders {
		func(i int, root context.Context) {
			builder := tmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TOTPMethodMutation)
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
					_, err = mutators[i+1].Mutate(root, tmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tmcb *TOTPMethodCreateBulk) SaveX(ctx context.Context) []*TOTPMethod {
	v, err := tmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmcb *TOTPMethodCreateBulk) Exec(ctx context.Context) error {
	_, err := tmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcb *TOTPMethodCreateBulk) ExecX(ctx context.Context) {
	if err := tmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"graph/ent/forum"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ForumCreate is the builder for creating a Forum entity.
type ForumCreate struct {
	config
	mutation *ForumMutation
	hooks    []Hook
}

// Mutation returns the ForumMutation object of the builder.
func (fc *ForumCreate) Mutation() *ForumMutation {
	return fc.mutation
}

// Save creates the Forum in the database.
func (fc *ForumCreate) Save(ctx context.Context) (*Forum, error) {
	var (
		err  error
		node *Forum
	)
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ForumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *ForumCreate) SaveX(ctx context.Context) *Forum {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (fc *ForumCreate) check() error {
	return nil
}

func (fc *ForumCreate) sqlSave(ctx context.Context) (*Forum, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *ForumCreate) createSpec() (*Forum, *sqlgraph.CreateSpec) {
	var (
		_node = &Forum{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: forum.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: forum.FieldID,
			},
		}
	)
	return _node, _spec
}

// ForumCreateBulk is the builder for creating a bulk of Forum entities.
type ForumCreateBulk struct {
	config
	builders []*ForumCreate
}

// Save creates the Forum entities in the database.
func (fcb *ForumCreateBulk) Save(ctx context.Context) ([]*Forum, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Forum, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ForumMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (fcb *ForumCreateBulk) SaveX(ctx context.Context) []*Forum {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

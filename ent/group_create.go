// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"ent_example/ent/group"
	"ent_example/ent/user"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetName sets the name field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetNickname sets the nickname field.
func (gc *GroupCreate) SetNickname(s string) *GroupCreate {
	gc.mutation.SetNickname(s)
	return gc
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (gc *GroupCreate) SetNillableNickname(s *string) *GroupCreate {
	if s != nil {
		gc.SetNickname(*s)
	}
	return gc
}

// SetCount sets the count field.
func (gc *GroupCreate) SetCount(i int) *GroupCreate {
	gc.mutation.SetCount(i)
	return gc
}

// SetNillableCount sets the count field if the given value is not nil.
func (gc *GroupCreate) SetNillableCount(i *int) *GroupCreate {
	if i != nil {
		gc.SetCount(*i)
	}
	return gc
}

// SetCode sets the code field.
func (gc *GroupCreate) SetCode(i int) *GroupCreate {
	gc.mutation.SetCode(i)
	return gc
}

// SetNillableCode sets the code field if the given value is not nil.
func (gc *GroupCreate) SetNillableCode(i *int) *GroupCreate {
	if i != nil {
		gc.SetCode(*i)
	}
	return gc
}

// SetIndex sets the index field.
func (gc *GroupCreate) SetIndex(i int) *GroupCreate {
	gc.mutation.SetIndex(i)
	return gc
}

// SetNillableIndex sets the index field if the given value is not nil.
func (gc *GroupCreate) SetNillableIndex(i *int) *GroupCreate {
	if i != nil {
		gc.SetIndex(*i)
	}
	return gc
}

// SetMin sets the min field.
func (gc *GroupCreate) SetMin(i int) *GroupCreate {
	gc.mutation.SetMin(i)
	return gc
}

// SetNillableMin sets the min field if the given value is not nil.
func (gc *GroupCreate) SetNillableMin(i *int) *GroupCreate {
	if i != nil {
		gc.SetMin(*i)
	}
	return gc
}

// SetMax sets the max field.
func (gc *GroupCreate) SetMax(i int) *GroupCreate {
	gc.mutation.SetMax(i)
	return gc
}

// SetNillableMax sets the max field if the given value is not nil.
func (gc *GroupCreate) SetNillableMax(i *int) *GroupCreate {
	if i != nil {
		gc.SetMax(*i)
	}
	return gc
}

// SetRange sets the range field.
func (gc *GroupCreate) SetRange(i int) *GroupCreate {
	gc.mutation.SetRange(i)
	return gc
}

// SetNillableRange sets the range field if the given value is not nil.
func (gc *GroupCreate) SetNillableRange(i *int) *GroupCreate {
	if i != nil {
		gc.SetRange(*i)
	}
	return gc
}

// SetNote sets the note field.
func (gc *GroupCreate) SetNote(s string) *GroupCreate {
	gc.mutation.SetNote(s)
	return gc
}

// SetNillableNote sets the note field if the given value is not nil.
func (gc *GroupCreate) SetNillableNote(s *string) *GroupCreate {
	if s != nil {
		gc.SetNote(*s)
	}
	return gc
}

// SetLog sets the log field.
func (gc *GroupCreate) SetLog(s string) *GroupCreate {
	gc.mutation.SetLog(s)
	return gc
}

// SetNillableLog sets the log field if the given value is not nil.
func (gc *GroupCreate) SetNillableLog(s *string) *GroupCreate {
	if s != nil {
		gc.SetLog(*s)
	}
	return gc
}

// SetUsername sets the username field.
func (gc *GroupCreate) SetUsername(s string) *GroupCreate {
	gc.mutation.SetUsername(s)
	return gc
}

// SetNillableUsername sets the username field if the given value is not nil.
func (gc *GroupCreate) SetNillableUsername(s *string) *GroupCreate {
	if s != nil {
		gc.SetUsername(*s)
	}
	return gc
}

// SetID sets the id field.
func (gc *GroupCreate) SetID(i int) *GroupCreate {
	gc.mutation.SetID(i)
	return gc
}

// AddUserIDs adds the users edge to User by ids.
func (gc *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the users edges to User.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Nickname(); ok {
		if err := group.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("ent: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Count(); ok {
		if err := group.CountValidator(v); err != nil {
			return &ValidationError{Name: "count", err: fmt.Errorf("ent: validator failed for field \"count\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Code(); ok {
		if err := group.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Index(); ok {
		if err := group.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf("ent: validator failed for field \"index\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Min(); ok {
		if err := group.MinValidator(v); err != nil {
			return &ValidationError{Name: "min", err: fmt.Errorf("ent: validator failed for field \"min\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Max(); ok {
		if err := group.MaxValidator(v); err != nil {
			return &ValidationError{Name: "max", err: fmt.Errorf("ent: validator failed for field \"max\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Range(); ok {
		if err := group.RangeValidator(v); err != nil {
			return &ValidationError{Name: "range", err: fmt.Errorf("ent: validator failed for field \"range\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Note(); ok {
		if err := group.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Log(); ok {
		if err := group.LogValidator(v); err != nil {
			return &ValidationError{Name: "log", err: fmt.Errorf("ent: validator failed for field \"log\": %w", err)}
		}
	}
	if v, ok := gc.mutation.Username(); ok {
		if err := group.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := gc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := gc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := gc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldIndex,
		})
		_node.Index = value
	}
	if value, ok := gc.mutation.Min(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMin,
		})
		_node.Min = value
	}
	if value, ok := gc.mutation.Max(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMax,
		})
		_node.Max = value
	}
	if value, ok := gc.mutation.Range(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldRange,
		})
		_node.Range = value
	}
	if value, ok := gc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNote,
		})
		_node.Note = value
	}
	if value, ok := gc.mutation.Log(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldLog,
		})
		_node.Log = value
	}
	if value, ok := gc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldUsername,
		})
		_node.Username = value
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupCreateBulk is the builder for creating a bulk of Group entities.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

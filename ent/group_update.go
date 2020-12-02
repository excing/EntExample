// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"ent_example/ent/group"
	"ent_example/ent/predicate"
	"ent_example/ent/user"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.predicates = append(gu.mutation.predicates, ps...)
	return gu
}

// SetName sets the name field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNickname sets the nickname field.
func (gu *GroupUpdate) SetNickname(s string) *GroupUpdate {
	gu.mutation.SetNickname(s)
	return gu
}

// SetCount sets the count field.
func (gu *GroupUpdate) SetCount(i int) *GroupUpdate {
	gu.mutation.ResetCount()
	gu.mutation.SetCount(i)
	return gu
}

// AddCount adds i to count.
func (gu *GroupUpdate) AddCount(i int) *GroupUpdate {
	gu.mutation.AddCount(i)
	return gu
}

// SetCode sets the code field.
func (gu *GroupUpdate) SetCode(i int) *GroupUpdate {
	gu.mutation.ResetCode()
	gu.mutation.SetCode(i)
	return gu
}

// AddCode adds i to code.
func (gu *GroupUpdate) AddCode(i int) *GroupUpdate {
	gu.mutation.AddCode(i)
	return gu
}

// SetIndex sets the index field.
func (gu *GroupUpdate) SetIndex(i int) *GroupUpdate {
	gu.mutation.ResetIndex()
	gu.mutation.SetIndex(i)
	return gu
}

// AddIndex adds i to index.
func (gu *GroupUpdate) AddIndex(i int) *GroupUpdate {
	gu.mutation.AddIndex(i)
	return gu
}

// SetMin sets the min field.
func (gu *GroupUpdate) SetMin(i int) *GroupUpdate {
	gu.mutation.ResetMin()
	gu.mutation.SetMin(i)
	return gu
}

// AddMin adds i to min.
func (gu *GroupUpdate) AddMin(i int) *GroupUpdate {
	gu.mutation.AddMin(i)
	return gu
}

// SetMax sets the max field.
func (gu *GroupUpdate) SetMax(i int) *GroupUpdate {
	gu.mutation.ResetMax()
	gu.mutation.SetMax(i)
	return gu
}

// AddMax adds i to max.
func (gu *GroupUpdate) AddMax(i int) *GroupUpdate {
	gu.mutation.AddMax(i)
	return gu
}

// SetRange sets the range field.
func (gu *GroupUpdate) SetRange(i int) *GroupUpdate {
	gu.mutation.ResetRange()
	gu.mutation.SetRange(i)
	return gu
}

// AddRange adds i to range.
func (gu *GroupUpdate) AddRange(i int) *GroupUpdate {
	gu.mutation.AddRange(i)
	return gu
}

// SetNote sets the note field.
func (gu *GroupUpdate) SetNote(s string) *GroupUpdate {
	gu.mutation.SetNote(s)
	return gu
}

// SetLog sets the log field.
func (gu *GroupUpdate) SetLog(s string) *GroupUpdate {
	gu.mutation.SetLog(s)
	return gu
}

// SetUsername sets the username field.
func (gu *GroupUpdate) SetUsername(s string) *GroupUpdate {
	gu.mutation.SetUsername(s)
	return gu
}

// AddUserIDs adds the users edge to User by ids.
func (gu *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the users edges to User.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearUsers clears all "users" edges to type User.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the users edge to User by ids.
func (gu *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes users edges to User.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Nickname(); ok {
		if err := group.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("ent: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Count(); ok {
		if err := group.CountValidator(v); err != nil {
			return &ValidationError{Name: "count", err: fmt.Errorf("ent: validator failed for field \"count\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Code(); ok {
		if err := group.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Index(); ok {
		if err := group.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf("ent: validator failed for field \"index\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Min(); ok {
		if err := group.MinValidator(v); err != nil {
			return &ValidationError{Name: "min", err: fmt.Errorf("ent: validator failed for field \"min\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Max(); ok {
		if err := group.MaxValidator(v); err != nil {
			return &ValidationError{Name: "max", err: fmt.Errorf("ent: validator failed for field \"max\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Range(); ok {
		if err := group.RangeValidator(v); err != nil {
			return &ValidationError{Name: "range", err: fmt.Errorf("ent: validator failed for field \"range\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Note(); ok {
		if err := group.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Log(); ok {
		if err := group.LogValidator(v); err != nil {
			return &ValidationError{Name: "log", err: fmt.Errorf("ent: validator failed for field \"log\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Username(); ok {
		if err := group.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := gu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNickname,
		})
	}
	if value, ok := gu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCount,
		})
	}
	if value, ok := gu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCount,
		})
	}
	if value, ok := gu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCode,
		})
	}
	if value, ok := gu.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCode,
		})
	}
	if value, ok := gu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldIndex,
		})
	}
	if value, ok := gu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldIndex,
		})
	}
	if value, ok := gu.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMin,
		})
	}
	if value, ok := gu.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMin,
		})
	}
	if value, ok := gu.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMax,
		})
	}
	if value, ok := gu.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMax,
		})
	}
	if value, ok := gu.mutation.Range(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldRange,
		})
	}
	if value, ok := gu.mutation.AddedRange(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldRange,
		})
	}
	if value, ok := gu.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNote,
		})
	}
	if value, ok := gu.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldLog,
		})
	}
	if value, ok := gu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldUsername,
		})
	}
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the name field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNickname sets the nickname field.
func (guo *GroupUpdateOne) SetNickname(s string) *GroupUpdateOne {
	guo.mutation.SetNickname(s)
	return guo
}

// SetCount sets the count field.
func (guo *GroupUpdateOne) SetCount(i int) *GroupUpdateOne {
	guo.mutation.ResetCount()
	guo.mutation.SetCount(i)
	return guo
}

// AddCount adds i to count.
func (guo *GroupUpdateOne) AddCount(i int) *GroupUpdateOne {
	guo.mutation.AddCount(i)
	return guo
}

// SetCode sets the code field.
func (guo *GroupUpdateOne) SetCode(i int) *GroupUpdateOne {
	guo.mutation.ResetCode()
	guo.mutation.SetCode(i)
	return guo
}

// AddCode adds i to code.
func (guo *GroupUpdateOne) AddCode(i int) *GroupUpdateOne {
	guo.mutation.AddCode(i)
	return guo
}

// SetIndex sets the index field.
func (guo *GroupUpdateOne) SetIndex(i int) *GroupUpdateOne {
	guo.mutation.ResetIndex()
	guo.mutation.SetIndex(i)
	return guo
}

// AddIndex adds i to index.
func (guo *GroupUpdateOne) AddIndex(i int) *GroupUpdateOne {
	guo.mutation.AddIndex(i)
	return guo
}

// SetMin sets the min field.
func (guo *GroupUpdateOne) SetMin(i int) *GroupUpdateOne {
	guo.mutation.ResetMin()
	guo.mutation.SetMin(i)
	return guo
}

// AddMin adds i to min.
func (guo *GroupUpdateOne) AddMin(i int) *GroupUpdateOne {
	guo.mutation.AddMin(i)
	return guo
}

// SetMax sets the max field.
func (guo *GroupUpdateOne) SetMax(i int) *GroupUpdateOne {
	guo.mutation.ResetMax()
	guo.mutation.SetMax(i)
	return guo
}

// AddMax adds i to max.
func (guo *GroupUpdateOne) AddMax(i int) *GroupUpdateOne {
	guo.mutation.AddMax(i)
	return guo
}

// SetRange sets the range field.
func (guo *GroupUpdateOne) SetRange(i int) *GroupUpdateOne {
	guo.mutation.ResetRange()
	guo.mutation.SetRange(i)
	return guo
}

// AddRange adds i to range.
func (guo *GroupUpdateOne) AddRange(i int) *GroupUpdateOne {
	guo.mutation.AddRange(i)
	return guo
}

// SetNote sets the note field.
func (guo *GroupUpdateOne) SetNote(s string) *GroupUpdateOne {
	guo.mutation.SetNote(s)
	return guo
}

// SetLog sets the log field.
func (guo *GroupUpdateOne) SetLog(s string) *GroupUpdateOne {
	guo.mutation.SetLog(s)
	return guo
}

// SetUsername sets the username field.
func (guo *GroupUpdateOne) SetUsername(s string) *GroupUpdateOne {
	guo.mutation.SetUsername(s)
	return guo
}

// AddUserIDs adds the users edge to User by ids.
func (guo *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the users edges to User.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearUsers clears all "users" edges to type User.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the users edge to User by ids.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes users edges to User.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Nickname(); ok {
		if err := group.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("ent: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Count(); ok {
		if err := group.CountValidator(v); err != nil {
			return &ValidationError{Name: "count", err: fmt.Errorf("ent: validator failed for field \"count\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Code(); ok {
		if err := group.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Index(); ok {
		if err := group.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf("ent: validator failed for field \"index\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Min(); ok {
		if err := group.MinValidator(v); err != nil {
			return &ValidationError{Name: "min", err: fmt.Errorf("ent: validator failed for field \"min\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Max(); ok {
		if err := group.MaxValidator(v); err != nil {
			return &ValidationError{Name: "max", err: fmt.Errorf("ent: validator failed for field \"max\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Range(); ok {
		if err := group.RangeValidator(v); err != nil {
			return &ValidationError{Name: "range", err: fmt.Errorf("ent: validator failed for field \"range\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Note(); ok {
		if err := group.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Log(); ok {
		if err := group.LogValidator(v); err != nil {
			return &ValidationError{Name: "log", err: fmt.Errorf("ent: validator failed for field \"log\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Username(); ok {
		if err := group.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Group.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := guo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNickname,
		})
	}
	if value, ok := guo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCount,
		})
	}
	if value, ok := guo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCount,
		})
	}
	if value, ok := guo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCode,
		})
	}
	if value, ok := guo.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldCode,
		})
	}
	if value, ok := guo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldIndex,
		})
	}
	if value, ok := guo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldIndex,
		})
	}
	if value, ok := guo.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMin,
		})
	}
	if value, ok := guo.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMin,
		})
	}
	if value, ok := guo.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMax,
		})
	}
	if value, ok := guo.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMax,
		})
	}
	if value, ok := guo.mutation.Range(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldRange,
		})
	}
	if value, ok := guo.mutation.AddedRange(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldRange,
		})
	}
	if value, ok := guo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldNote,
		})
	}
	if value, ok := guo.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldLog,
		})
	}
	if value, ok := guo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldUsername,
		})
	}
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

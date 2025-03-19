// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/accesstoken"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/user"
	"github.com/google/uuid"
)

// AccessTokenUpdate is the builder for updating AccessToken entities.
type AccessTokenUpdate struct {
	config
	hooks     []Hook
	mutation  *AccessTokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atu *AccessTokenUpdate) Where(ps ...predicate.AccessToken) *AccessTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetName sets the "name" field.
func (atu *AccessTokenUpdate) SetName(s string) *AccessTokenUpdate {
	atu.mutation.SetName(s)
	return atu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableName(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetName(*s)
	}
	return atu
}

// SetUserID sets the "user_id" field.
func (atu *AccessTokenUpdate) SetUserID(u uuid.UUID) *AccessTokenUpdate {
	atu.mutation.SetUserID(u)
	return atu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableUserID(u *uuid.UUID) *AccessTokenUpdate {
	if u != nil {
		atu.SetUserID(*u)
	}
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *AccessTokenUpdate) SetUser(u *User) *AccessTokenUpdate {
	return atu.SetUserID(u.ID)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atu *AccessTokenUpdate) Mutation() *AccessTokenMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *AccessTokenUpdate) ClearUser() *AccessTokenUpdate {
	atu.mutation.ClearUser()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AccessTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AccessTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AccessTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AccessTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AccessTokenUpdate) check() error {
	if _, ok := atu.mutation.UserID(); atu.mutation.UserCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "AccessToken.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (atu *AccessTokenUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccessTokenUpdate {
	atu.modifiers = append(atu.modifiers, modifiers...)
	return atu
}

func (atu *AccessTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.SetField(accesstoken.FieldName, field.TypeString, value)
	}
	if atu.mutation.CreatedAtCleared() {
		_spec.ClearField(accesstoken.FieldCreatedAt, field.TypeTime)
	}
	if atu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = atu.schemaConfig.AccessToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = atu.schemaConfig.AccessToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = atu.schemaConfig.AccessToken
	ctx = internal.NewSchemaConfigContext(ctx, atu.schemaConfig)
	_spec.AddModifiers(atu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AccessTokenUpdateOne is the builder for updating a single AccessToken entity.
type AccessTokenUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AccessTokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (atuo *AccessTokenUpdateOne) SetName(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableName(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetName(*s)
	}
	return atuo
}

// SetUserID sets the "user_id" field.
func (atuo *AccessTokenUpdateOne) SetUserID(u uuid.UUID) *AccessTokenUpdateOne {
	atuo.mutation.SetUserID(u)
	return atuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableUserID(u *uuid.UUID) *AccessTokenUpdateOne {
	if u != nil {
		atuo.SetUserID(*u)
	}
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *AccessTokenUpdateOne) SetUser(u *User) *AccessTokenUpdateOne {
	return atuo.SetUserID(u.ID)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atuo *AccessTokenUpdateOne) Mutation() *AccessTokenMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *AccessTokenUpdateOne) ClearUser() *AccessTokenUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atuo *AccessTokenUpdateOne) Where(ps ...predicate.AccessToken) *AccessTokenUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AccessTokenUpdateOne) Select(field string, fields ...string) *AccessTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AccessToken entity.
func (atuo *AccessTokenUpdateOne) Save(ctx context.Context) (*AccessToken, error) {
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) SaveX(ctx context.Context) *AccessToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AccessTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AccessTokenUpdateOne) check() error {
	if _, ok := atuo.mutation.UserID(); atuo.mutation.UserCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "AccessToken.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (atuo *AccessTokenUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccessTokenUpdateOne {
	atuo.modifiers = append(atuo.modifiers, modifiers...)
	return atuo
}

func (atuo *AccessTokenUpdateOne) sqlSave(ctx context.Context) (_node *AccessToken, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "AccessToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesstoken.FieldID)
		for _, f := range fields {
			if !accesstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != accesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.Name(); ok {
		_spec.SetField(accesstoken.FieldName, field.TypeString, value)
	}
	if atuo.mutation.CreatedAtCleared() {
		_spec.ClearField(accesstoken.FieldCreatedAt, field.TypeTime)
	}
	if atuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = atuo.schemaConfig.AccessToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = atuo.schemaConfig.AccessToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = atuo.schemaConfig.AccessToken
	ctx = internal.NewSchemaConfigContext(ctx, atuo.schemaConfig)
	_spec.AddModifiers(atuo.modifiers...)
	_node = &AccessToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/env"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/envalias"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/team"
	"github.com/google/uuid"
)

// EnvUpdate is the builder for updating Env entities.
type EnvUpdate struct {
	config
	hooks    []Hook
	mutation *EnvMutation
}

// Where appends a list predicates to the EnvUpdate builder.
func (eu *EnvUpdate) Where(ps ...predicate.Env) *EnvUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EnvUpdate) SetUpdatedAt(t time.Time) *EnvUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eu *EnvUpdate) SetNillableUpdatedAt(t *time.Time) *EnvUpdate {
	if t != nil {
		eu.SetUpdatedAt(*t)
	}
	return eu
}

// SetTeamID sets the "team_id" field.
func (eu *EnvUpdate) SetTeamID(u uuid.UUID) *EnvUpdate {
	eu.mutation.SetTeamID(u)
	return eu
}

// SetDockerfile sets the "dockerfile" field.
func (eu *EnvUpdate) SetDockerfile(s string) *EnvUpdate {
	eu.mutation.SetDockerfile(s)
	return eu
}

// SetPublic sets the "public" field.
func (eu *EnvUpdate) SetPublic(b bool) *EnvUpdate {
	eu.mutation.SetPublic(b)
	return eu
}

// SetBuildID sets the "build_id" field.
func (eu *EnvUpdate) SetBuildID(u uuid.UUID) *EnvUpdate {
	eu.mutation.SetBuildID(u)
	return eu
}

// SetBuildCount sets the "build_count" field.
func (eu *EnvUpdate) SetBuildCount(i int) *EnvUpdate {
	eu.mutation.ResetBuildCount()
	eu.mutation.SetBuildCount(i)
	return eu
}

// SetNillableBuildCount sets the "build_count" field if the given value is not nil.
func (eu *EnvUpdate) SetNillableBuildCount(i *int) *EnvUpdate {
	if i != nil {
		eu.SetBuildCount(*i)
	}
	return eu
}

// AddBuildCount adds i to the "build_count" field.
func (eu *EnvUpdate) AddBuildCount(i int) *EnvUpdate {
	eu.mutation.AddBuildCount(i)
	return eu
}

// SetTeam sets the "team" edge to the Team entity.
func (eu *EnvUpdate) SetTeam(t *Team) *EnvUpdate {
	return eu.SetTeamID(t.ID)
}

// AddEnvAliasIDs adds the "env_aliases" edge to the EnvAlias entity by IDs.
func (eu *EnvUpdate) AddEnvAliasIDs(ids ...int) *EnvUpdate {
	eu.mutation.AddEnvAliasIDs(ids...)
	return eu
}

// AddEnvAliases adds the "env_aliases" edges to the EnvAlias entity.
func (eu *EnvUpdate) AddEnvAliases(e ...*EnvAlias) *EnvUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.AddEnvAliasIDs(ids...)
}

// Mutation returns the EnvMutation object of the builder.
func (eu *EnvUpdate) Mutation() *EnvMutation {
	return eu.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (eu *EnvUpdate) ClearTeam() *EnvUpdate {
	eu.mutation.ClearTeam()
	return eu
}

// ClearEnvAliases clears all "env_aliases" edges to the EnvAlias entity.
func (eu *EnvUpdate) ClearEnvAliases() *EnvUpdate {
	eu.mutation.ClearEnvAliases()
	return eu
}

// RemoveEnvAliasIDs removes the "env_aliases" edge to EnvAlias entities by IDs.
func (eu *EnvUpdate) RemoveEnvAliasIDs(ids ...int) *EnvUpdate {
	eu.mutation.RemoveEnvAliasIDs(ids...)
	return eu
}

// RemoveEnvAliases removes "env_aliases" edges to EnvAlias entities.
func (eu *EnvUpdate) RemoveEnvAliases(e ...*EnvAlias) *EnvUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.RemoveEnvAliasIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EnvUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EnvUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EnvUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EnvUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EnvUpdate) check() error {
	if _, ok := eu.mutation.TeamID(); eu.mutation.TeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Env.team"`)
	}
	return nil
}

func (eu *EnvUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(env.Table, env.Columns, sqlgraph.NewFieldSpec(env.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(env.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.Dockerfile(); ok {
		_spec.SetField(env.FieldDockerfile, field.TypeString, value)
	}
	if value, ok := eu.mutation.Public(); ok {
		_spec.SetField(env.FieldPublic, field.TypeBool, value)
	}
	if value, ok := eu.mutation.BuildID(); ok {
		_spec.SetField(env.FieldBuildID, field.TypeUUID, value)
	}
	if value, ok := eu.mutation.BuildCount(); ok {
		_spec.SetField(env.FieldBuildCount, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedBuildCount(); ok {
		_spec.AddField(env.FieldBuildCount, field.TypeInt, value)
	}
	if eu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   env.TeamTable,
			Columns: []string{env.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = eu.schemaConfig.Env
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   env.TeamTable,
			Columns: []string{env.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = eu.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.EnvAliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = eu.schemaConfig.EnvAlias
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedEnvAliasesIDs(); len(nodes) > 0 && !eu.mutation.EnvAliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = eu.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.EnvAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = eu.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = eu.schemaConfig.Env
	ctx = internal.NewSchemaConfigContext(ctx, eu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{env.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EnvUpdateOne is the builder for updating a single Env entity.
type EnvUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EnvMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EnvUpdateOne) SetUpdatedAt(t time.Time) *EnvUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (euo *EnvUpdateOne) SetNillableUpdatedAt(t *time.Time) *EnvUpdateOne {
	if t != nil {
		euo.SetUpdatedAt(*t)
	}
	return euo
}

// SetTeamID sets the "team_id" field.
func (euo *EnvUpdateOne) SetTeamID(u uuid.UUID) *EnvUpdateOne {
	euo.mutation.SetTeamID(u)
	return euo
}

// SetDockerfile sets the "dockerfile" field.
func (euo *EnvUpdateOne) SetDockerfile(s string) *EnvUpdateOne {
	euo.mutation.SetDockerfile(s)
	return euo
}

// SetPublic sets the "public" field.
func (euo *EnvUpdateOne) SetPublic(b bool) *EnvUpdateOne {
	euo.mutation.SetPublic(b)
	return euo
}

// SetBuildID sets the "build_id" field.
func (euo *EnvUpdateOne) SetBuildID(u uuid.UUID) *EnvUpdateOne {
	euo.mutation.SetBuildID(u)
	return euo
}

// SetBuildCount sets the "build_count" field.
func (euo *EnvUpdateOne) SetBuildCount(i int) *EnvUpdateOne {
	euo.mutation.ResetBuildCount()
	euo.mutation.SetBuildCount(i)
	return euo
}

// SetNillableBuildCount sets the "build_count" field if the given value is not nil.
func (euo *EnvUpdateOne) SetNillableBuildCount(i *int) *EnvUpdateOne {
	if i != nil {
		euo.SetBuildCount(*i)
	}
	return euo
}

// AddBuildCount adds i to the "build_count" field.
func (euo *EnvUpdateOne) AddBuildCount(i int) *EnvUpdateOne {
	euo.mutation.AddBuildCount(i)
	return euo
}

// SetTeam sets the "team" edge to the Team entity.
func (euo *EnvUpdateOne) SetTeam(t *Team) *EnvUpdateOne {
	return euo.SetTeamID(t.ID)
}

// AddEnvAliasIDs adds the "env_aliases" edge to the EnvAlias entity by IDs.
func (euo *EnvUpdateOne) AddEnvAliasIDs(ids ...int) *EnvUpdateOne {
	euo.mutation.AddEnvAliasIDs(ids...)
	return euo
}

// AddEnvAliases adds the "env_aliases" edges to the EnvAlias entity.
func (euo *EnvUpdateOne) AddEnvAliases(e ...*EnvAlias) *EnvUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.AddEnvAliasIDs(ids...)
}

// Mutation returns the EnvMutation object of the builder.
func (euo *EnvUpdateOne) Mutation() *EnvMutation {
	return euo.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (euo *EnvUpdateOne) ClearTeam() *EnvUpdateOne {
	euo.mutation.ClearTeam()
	return euo
}

// ClearEnvAliases clears all "env_aliases" edges to the EnvAlias entity.
func (euo *EnvUpdateOne) ClearEnvAliases() *EnvUpdateOne {
	euo.mutation.ClearEnvAliases()
	return euo
}

// RemoveEnvAliasIDs removes the "env_aliases" edge to EnvAlias entities by IDs.
func (euo *EnvUpdateOne) RemoveEnvAliasIDs(ids ...int) *EnvUpdateOne {
	euo.mutation.RemoveEnvAliasIDs(ids...)
	return euo
}

// RemoveEnvAliases removes "env_aliases" edges to EnvAlias entities.
func (euo *EnvUpdateOne) RemoveEnvAliases(e ...*EnvAlias) *EnvUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.RemoveEnvAliasIDs(ids...)
}

// Where appends a list predicates to the EnvUpdate builder.
func (euo *EnvUpdateOne) Where(ps ...predicate.Env) *EnvUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EnvUpdateOne) Select(field string, fields ...string) *EnvUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Env entity.
func (euo *EnvUpdateOne) Save(ctx context.Context) (*Env, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EnvUpdateOne) SaveX(ctx context.Context) *Env {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EnvUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EnvUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EnvUpdateOne) check() error {
	if _, ok := euo.mutation.TeamID(); euo.mutation.TeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Env.team"`)
	}
	return nil
}

func (euo *EnvUpdateOne) sqlSave(ctx context.Context) (_node *Env, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(env.Table, env.Columns, sqlgraph.NewFieldSpec(env.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Env.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, env.FieldID)
		for _, f := range fields {
			if !env.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != env.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(env.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.Dockerfile(); ok {
		_spec.SetField(env.FieldDockerfile, field.TypeString, value)
	}
	if value, ok := euo.mutation.Public(); ok {
		_spec.SetField(env.FieldPublic, field.TypeBool, value)
	}
	if value, ok := euo.mutation.BuildID(); ok {
		_spec.SetField(env.FieldBuildID, field.TypeUUID, value)
	}
	if value, ok := euo.mutation.BuildCount(); ok {
		_spec.SetField(env.FieldBuildCount, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedBuildCount(); ok {
		_spec.AddField(env.FieldBuildCount, field.TypeInt, value)
	}
	if euo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   env.TeamTable,
			Columns: []string{env.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = euo.schemaConfig.Env
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   env.TeamTable,
			Columns: []string{env.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = euo.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.EnvAliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = euo.schemaConfig.EnvAlias
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedEnvAliasesIDs(); len(nodes) > 0 && !euo.mutation.EnvAliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = euo.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.EnvAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   env.EnvAliasesTable,
			Columns: []string{env.EnvAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeInt),
			},
		}
		edge.Schema = euo.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = euo.schemaConfig.Env
	ctx = internal.NewSchemaConfigContext(ctx, euo.schemaConfig)
	_node = &Env{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{env.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}

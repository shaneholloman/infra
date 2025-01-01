// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/snapshot"
	"github.com/google/uuid"
)

// SnapshotCreate is the builder for creating a Snapshot entity.
type SnapshotCreate struct {
	config
	mutation *SnapshotMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SnapshotCreate) SetCreatedAt(t time.Time) *SnapshotCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SnapshotCreate) SetNillableCreatedAt(t *time.Time) *SnapshotCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetBaseEnvID sets the "base_env_id" field.
func (sc *SnapshotCreate) SetBaseEnvID(s string) *SnapshotCreate {
	sc.mutation.SetBaseEnvID(s)
	return sc
}

// SetEnvID sets the "env_id" field.
func (sc *SnapshotCreate) SetEnvID(s string) *SnapshotCreate {
	sc.mutation.SetEnvID(s)
	return sc
}

// SetSandboxID sets the "sandbox_id" field.
func (sc *SnapshotCreate) SetSandboxID(s string) *SnapshotCreate {
	sc.mutation.SetSandboxID(s)
	return sc
}

// SetMetadata sets the "metadata" field.
func (sc *SnapshotCreate) SetMetadata(m map[string]string) *SnapshotCreate {
	sc.mutation.SetMetadata(m)
	return sc
}

// SetID sets the "id" field.
func (sc *SnapshotCreate) SetID(u uuid.UUID) *SnapshotCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetEnv sets the "env" edge to the Env entity.
func (sc *SnapshotCreate) SetEnv(e *Env) *SnapshotCreate {
	return sc.SetEnvID(e.ID)
}

// Mutation returns the SnapshotMutation object of the builder.
func (sc *SnapshotCreate) Mutation() *SnapshotMutation {
	return sc.mutation
}

// Save creates the Snapshot in the database.
func (sc *SnapshotCreate) Save(ctx context.Context) (*Snapshot, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SnapshotCreate) SaveX(ctx context.Context) *Snapshot {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SnapshotCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SnapshotCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SnapshotCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := snapshot.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SnapshotCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "Snapshot.created_at"`)}
	}
	if _, ok := sc.mutation.BaseEnvID(); !ok {
		return &ValidationError{Name: "base_env_id", err: errors.New(`models: missing required field "Snapshot.base_env_id"`)}
	}
	if _, ok := sc.mutation.EnvID(); !ok {
		return &ValidationError{Name: "env_id", err: errors.New(`models: missing required field "Snapshot.env_id"`)}
	}
	if _, ok := sc.mutation.SandboxID(); !ok {
		return &ValidationError{Name: "sandbox_id", err: errors.New(`models: missing required field "Snapshot.sandbox_id"`)}
	}
	if _, ok := sc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New(`models: missing required field "Snapshot.metadata"`)}
	}
	if _, ok := sc.mutation.EnvID(); !ok {
		return &ValidationError{Name: "env", err: errors.New(`models: missing required edge "Snapshot.env"`)}
	}
	return nil
}

func (sc *SnapshotCreate) sqlSave(ctx context.Context) (*Snapshot, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SnapshotCreate) createSpec() (*Snapshot, *sqlgraph.CreateSpec) {
	var (
		_node = &Snapshot{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(snapshot.Table, sqlgraph.NewFieldSpec(snapshot.FieldID, field.TypeUUID))
	)
	_spec.Schema = sc.schemaConfig.Snapshot
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(snapshot.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.BaseEnvID(); ok {
		_spec.SetField(snapshot.FieldBaseEnvID, field.TypeString, value)
		_node.BaseEnvID = value
	}
	if value, ok := sc.mutation.SandboxID(); ok {
		_spec.SetField(snapshot.FieldSandboxID, field.TypeString, value)
		_node.SandboxID = value
	}
	if value, ok := sc.mutation.Metadata(); ok {
		_spec.SetField(snapshot.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := sc.mutation.EnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   snapshot.EnvTable,
			Columns: []string{snapshot.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.Snapshot
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EnvID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Snapshot.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SnapshotUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *SnapshotCreate) OnConflict(opts ...sql.ConflictOption) *SnapshotUpsertOne {
	sc.conflict = opts
	return &SnapshotUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SnapshotCreate) OnConflictColumns(columns ...string) *SnapshotUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SnapshotUpsertOne{
		create: sc,
	}
}

type (
	// SnapshotUpsertOne is the builder for "upsert"-ing
	//  one Snapshot node.
	SnapshotUpsertOne struct {
		create *SnapshotCreate
	}

	// SnapshotUpsert is the "OnConflict" setter.
	SnapshotUpsert struct {
		*sql.UpdateSet
	}
)

// SetBaseEnvID sets the "base_env_id" field.
func (u *SnapshotUpsert) SetBaseEnvID(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldBaseEnvID, v)
	return u
}

// UpdateBaseEnvID sets the "base_env_id" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateBaseEnvID() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldBaseEnvID)
	return u
}

// SetEnvID sets the "env_id" field.
func (u *SnapshotUpsert) SetEnvID(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldEnvID, v)
	return u
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateEnvID() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldEnvID)
	return u
}

// SetSandboxID sets the "sandbox_id" field.
func (u *SnapshotUpsert) SetSandboxID(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldSandboxID, v)
	return u
}

// UpdateSandboxID sets the "sandbox_id" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateSandboxID() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldSandboxID)
	return u
}

// SetMetadata sets the "metadata" field.
func (u *SnapshotUpsert) SetMetadata(v map[string]string) *SnapshotUpsert {
	u.Set(snapshot.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateMetadata() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldMetadata)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(snapshot.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SnapshotUpsertOne) UpdateNewValues() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(snapshot.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(snapshot.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SnapshotUpsertOne) Ignore() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SnapshotUpsertOne) DoNothing() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SnapshotCreate.OnConflict
// documentation for more info.
func (u *SnapshotUpsertOne) Update(set func(*SnapshotUpsert)) *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SnapshotUpsert{UpdateSet: update})
	}))
	return u
}

// SetBaseEnvID sets the "base_env_id" field.
func (u *SnapshotUpsertOne) SetBaseEnvID(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetBaseEnvID(v)
	})
}

// UpdateBaseEnvID sets the "base_env_id" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateBaseEnvID() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateBaseEnvID()
	})
}

// SetEnvID sets the "env_id" field.
func (u *SnapshotUpsertOne) SetEnvID(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetEnvID(v)
	})
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateEnvID() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateEnvID()
	})
}

// SetSandboxID sets the "sandbox_id" field.
func (u *SnapshotUpsertOne) SetSandboxID(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSandboxID(v)
	})
}

// UpdateSandboxID sets the "sandbox_id" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateSandboxID() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSandboxID()
	})
}

// SetMetadata sets the "metadata" field.
func (u *SnapshotUpsertOne) SetMetadata(v map[string]string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateMetadata() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateMetadata()
	})
}

// Exec executes the query.
func (u *SnapshotUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for SnapshotCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SnapshotUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SnapshotUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("models: SnapshotUpsertOne.ID is not supported by MySQL driver. Use SnapshotUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SnapshotUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SnapshotCreateBulk is the builder for creating many Snapshot entities in bulk.
type SnapshotCreateBulk struct {
	config
	err      error
	builders []*SnapshotCreate
	conflict []sql.ConflictOption
}

// Save creates the Snapshot entities in the database.
func (scb *SnapshotCreateBulk) Save(ctx context.Context) ([]*Snapshot, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Snapshot, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SnapshotMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SnapshotCreateBulk) SaveX(ctx context.Context) []*Snapshot {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SnapshotCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SnapshotCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Snapshot.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SnapshotUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *SnapshotCreateBulk) OnConflict(opts ...sql.ConflictOption) *SnapshotUpsertBulk {
	scb.conflict = opts
	return &SnapshotUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SnapshotCreateBulk) OnConflictColumns(columns ...string) *SnapshotUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SnapshotUpsertBulk{
		create: scb,
	}
}

// SnapshotUpsertBulk is the builder for "upsert"-ing
// a bulk of Snapshot nodes.
type SnapshotUpsertBulk struct {
	create *SnapshotCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(snapshot.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SnapshotUpsertBulk) UpdateNewValues() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(snapshot.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(snapshot.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SnapshotUpsertBulk) Ignore() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SnapshotUpsertBulk) DoNothing() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SnapshotCreateBulk.OnConflict
// documentation for more info.
func (u *SnapshotUpsertBulk) Update(set func(*SnapshotUpsert)) *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SnapshotUpsert{UpdateSet: update})
	}))
	return u
}

// SetBaseEnvID sets the "base_env_id" field.
func (u *SnapshotUpsertBulk) SetBaseEnvID(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetBaseEnvID(v)
	})
}

// UpdateBaseEnvID sets the "base_env_id" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateBaseEnvID() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateBaseEnvID()
	})
}

// SetEnvID sets the "env_id" field.
func (u *SnapshotUpsertBulk) SetEnvID(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetEnvID(v)
	})
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateEnvID() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateEnvID()
	})
}

// SetSandboxID sets the "sandbox_id" field.
func (u *SnapshotUpsertBulk) SetSandboxID(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSandboxID(v)
	})
}

// UpdateSandboxID sets the "sandbox_id" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateSandboxID() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSandboxID()
	})
}

// SetMetadata sets the "metadata" field.
func (u *SnapshotUpsertBulk) SetMetadata(v map[string]string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateMetadata() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateMetadata()
	})
}

// Exec executes the query.
func (u *SnapshotUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("models: OnConflict was set for builder %d. Set it on the SnapshotCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for SnapshotCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SnapshotUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

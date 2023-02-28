// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/module"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ModuleCreate is the builder for creating a Module entity.
type ModuleCreate struct {
	config
	mutation *ModuleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStatus sets the "status" field.
func (mc *ModuleCreate) SetStatus(s string) *ModuleCreate {
	mc.mutation.SetStatus(s)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableStatus(s *string) *ModuleCreate {
	if s != nil {
		mc.SetStatus(*s)
	}
	return mc
}

// SetStatusMessage sets the "statusMessage" field.
func (mc *ModuleCreate) SetStatusMessage(s string) *ModuleCreate {
	mc.mutation.SetStatusMessage(s)
	return mc
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableStatusMessage(s *string) *ModuleCreate {
	if s != nil {
		mc.SetStatusMessage(*s)
	}
	return mc
}

// SetCreateTime sets the "createTime" field.
func (mc *ModuleCreate) SetCreateTime(t time.Time) *ModuleCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableCreateTime(t *time.Time) *ModuleCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "updateTime" field.
func (mc *ModuleCreate) SetUpdateTime(t time.Time) *ModuleCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableUpdateTime(t *time.Time) *ModuleCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *ModuleCreate) SetDescription(s string) *ModuleCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableDescription(s *string) *ModuleCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetLabels sets the "labels" field.
func (mc *ModuleCreate) SetLabels(m map[string]string) *ModuleCreate {
	mc.mutation.SetLabels(m)
	return mc
}

// SetSource sets the "source" field.
func (mc *ModuleCreate) SetSource(s string) *ModuleCreate {
	mc.mutation.SetSource(s)
	return mc
}

// SetVersion sets the "version" field.
func (mc *ModuleCreate) SetVersion(s string) *ModuleCreate {
	mc.mutation.SetVersion(s)
	return mc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (mc *ModuleCreate) SetNillableVersion(s *string) *ModuleCreate {
	if s != nil {
		mc.SetVersion(*s)
	}
	return mc
}

// SetSchema sets the "schema" field.
func (mc *ModuleCreate) SetSchema(ts *types.ModuleSchema) *ModuleCreate {
	mc.mutation.SetSchema(ts)
	return mc
}

// SetID sets the "id" field.
func (mc *ModuleCreate) SetID(s string) *ModuleCreate {
	mc.mutation.SetID(s)
	return mc
}

// Mutation returns the ModuleMutation object of the builder.
func (mc *ModuleCreate) Mutation() *ModuleMutation {
	return mc.mutation
}

// Save creates the Module in the database.
func (mc *ModuleCreate) Save(ctx context.Context) (*Module, error) {
	mc.defaults()
	return withHooks[*Module, ModuleMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModuleCreate) SaveX(ctx context.Context) *Module {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModuleCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModuleCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *ModuleCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := module.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := module.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
	if _, ok := mc.mutation.Labels(); !ok {
		v := module.DefaultLabels
		mc.mutation.SetLabels(v)
	}
	if _, ok := mc.mutation.Schema(); !ok {
		v := module.DefaultSchema
		mc.mutation.SetSchema(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModuleCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "Module.createTime"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`model: missing required field "Module.updateTime"`)}
	}
	if _, ok := mc.mutation.Labels(); !ok {
		return &ValidationError{Name: "labels", err: errors.New(`model: missing required field "Module.labels"`)}
	}
	if _, ok := mc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`model: missing required field "Module.source"`)}
	}
	if v, ok := mc.mutation.Source(); ok {
		if err := module.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "Module.source": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Schema(); !ok {
		return &ValidationError{Name: "schema", err: errors.New(`model: missing required field "Module.schema"`)}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := module.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`model: validator failed for field "Module.id": %w`, err)}
		}
	}
	return nil
}

func (mc *ModuleCreate) sqlSave(ctx context.Context) (*Module, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Module.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *ModuleCreate) createSpec() (*Module, *sqlgraph.CreateSpec) {
	var (
		_node = &Module{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(module.Table, sqlgraph.NewFieldSpec(module.FieldID, field.TypeString))
	)
	_spec.Schema = mc.schemaConfig.Module
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(module.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.StatusMessage(); ok {
		_spec.SetField(module.FieldStatusMessage, field.TypeString, value)
		_node.StatusMessage = value
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(module.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(module.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(module.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.Labels(); ok {
		_spec.SetField(module.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	if value, ok := mc.mutation.Source(); ok {
		_spec.SetField(module.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := mc.mutation.Version(); ok {
		_spec.SetField(module.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := mc.mutation.Schema(); ok {
		_spec.SetField(module.FieldSchema, field.TypeJSON, value)
		_node.Schema = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Module.Create().
//		SetStatus(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ModuleUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (mc *ModuleCreate) OnConflict(opts ...sql.ConflictOption) *ModuleUpsertOne {
	mc.conflict = opts
	return &ModuleUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Module.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *ModuleCreate) OnConflictColumns(columns ...string) *ModuleUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &ModuleUpsertOne{
		create: mc,
	}
}

type (
	// ModuleUpsertOne is the builder for "upsert"-ing
	//  one Module node.
	ModuleUpsertOne struct {
		create *ModuleCreate
	}

	// ModuleUpsert is the "OnConflict" setter.
	ModuleUpsert struct {
		*sql.UpdateSet
	}
)

// SetStatus sets the "status" field.
func (u *ModuleUpsert) SetStatus(v string) *ModuleUpsert {
	u.Set(module.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateStatus() *ModuleUpsert {
	u.SetExcluded(module.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *ModuleUpsert) ClearStatus() *ModuleUpsert {
	u.SetNull(module.FieldStatus)
	return u
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ModuleUpsert) SetStatusMessage(v string) *ModuleUpsert {
	u.Set(module.FieldStatusMessage, v)
	return u
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateStatusMessage() *ModuleUpsert {
	u.SetExcluded(module.FieldStatusMessage)
	return u
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ModuleUpsert) ClearStatusMessage() *ModuleUpsert {
	u.SetNull(module.FieldStatusMessage)
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *ModuleUpsert) SetUpdateTime(v time.Time) *ModuleUpsert {
	u.Set(module.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateUpdateTime() *ModuleUpsert {
	u.SetExcluded(module.FieldUpdateTime)
	return u
}

// SetDescription sets the "description" field.
func (u *ModuleUpsert) SetDescription(v string) *ModuleUpsert {
	u.Set(module.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateDescription() *ModuleUpsert {
	u.SetExcluded(module.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *ModuleUpsert) ClearDescription() *ModuleUpsert {
	u.SetNull(module.FieldDescription)
	return u
}

// SetLabels sets the "labels" field.
func (u *ModuleUpsert) SetLabels(v map[string]string) *ModuleUpsert {
	u.Set(module.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateLabels() *ModuleUpsert {
	u.SetExcluded(module.FieldLabels)
	return u
}

// SetSource sets the "source" field.
func (u *ModuleUpsert) SetSource(v string) *ModuleUpsert {
	u.Set(module.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateSource() *ModuleUpsert {
	u.SetExcluded(module.FieldSource)
	return u
}

// SetVersion sets the "version" field.
func (u *ModuleUpsert) SetVersion(v string) *ModuleUpsert {
	u.Set(module.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateVersion() *ModuleUpsert {
	u.SetExcluded(module.FieldVersion)
	return u
}

// ClearVersion clears the value of the "version" field.
func (u *ModuleUpsert) ClearVersion() *ModuleUpsert {
	u.SetNull(module.FieldVersion)
	return u
}

// SetSchema sets the "schema" field.
func (u *ModuleUpsert) SetSchema(v *types.ModuleSchema) *ModuleUpsert {
	u.Set(module.FieldSchema, v)
	return u
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *ModuleUpsert) UpdateSchema() *ModuleUpsert {
	u.SetExcluded(module.FieldSchema)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Module.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(module.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ModuleUpsertOne) UpdateNewValues() *ModuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(module.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(module.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Module.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ModuleUpsertOne) Ignore() *ModuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ModuleUpsertOne) DoNothing() *ModuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ModuleCreate.OnConflict
// documentation for more info.
func (u *ModuleUpsertOne) Update(set func(*ModuleUpsert)) *ModuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ModuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ModuleUpsertOne) SetStatus(v string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateStatus() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ModuleUpsertOne) ClearStatus() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ModuleUpsertOne) SetStatusMessage(v string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateStatusMessage() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ModuleUpsertOne) ClearStatusMessage() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *ModuleUpsertOne) SetUpdateTime(v time.Time) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateUpdateTime() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *ModuleUpsertOne) SetDescription(v string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateDescription() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ModuleUpsertOne) ClearDescription() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *ModuleUpsertOne) SetLabels(v map[string]string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateLabels() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateLabels()
	})
}

// SetSource sets the "source" field.
func (u *ModuleUpsertOne) SetSource(v string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateSource() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateSource()
	})
}

// SetVersion sets the "version" field.
func (u *ModuleUpsertOne) SetVersion(v string) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateVersion() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *ModuleUpsertOne) ClearVersion() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearVersion()
	})
}

// SetSchema sets the "schema" field.
func (u *ModuleUpsertOne) SetSchema(v *types.ModuleSchema) *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.SetSchema(v)
	})
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *ModuleUpsertOne) UpdateSchema() *ModuleUpsertOne {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateSchema()
	})
}

// Exec executes the query.
func (u *ModuleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ModuleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ModuleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ModuleUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ModuleUpsertOne.ID is not supported by MySQL driver. Use ModuleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ModuleUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ModuleCreateBulk is the builder for creating many Module entities in bulk.
type ModuleCreateBulk struct {
	config
	builders []*ModuleCreate
	conflict []sql.ConflictOption
}

// Save creates the Module entities in the database.
func (mcb *ModuleCreateBulk) Save(ctx context.Context) ([]*Module, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Module, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModuleMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *ModuleCreateBulk) SaveX(ctx context.Context) []*Module {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModuleCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModuleCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Module.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ModuleUpsert) {
//			SetStatus(v+v).
//		}).
//		Exec(ctx)
func (mcb *ModuleCreateBulk) OnConflict(opts ...sql.ConflictOption) *ModuleUpsertBulk {
	mcb.conflict = opts
	return &ModuleUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Module.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *ModuleCreateBulk) OnConflictColumns(columns ...string) *ModuleUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &ModuleUpsertBulk{
		create: mcb,
	}
}

// ModuleUpsertBulk is the builder for "upsert"-ing
// a bulk of Module nodes.
type ModuleUpsertBulk struct {
	create *ModuleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Module.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(module.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ModuleUpsertBulk) UpdateNewValues() *ModuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(module.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(module.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Module.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ModuleUpsertBulk) Ignore() *ModuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ModuleUpsertBulk) DoNothing() *ModuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ModuleCreateBulk.OnConflict
// documentation for more info.
func (u *ModuleUpsertBulk) Update(set func(*ModuleUpsert)) *ModuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ModuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ModuleUpsertBulk) SetStatus(v string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateStatus() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *ModuleUpsertBulk) ClearStatus() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearStatus()
	})
}

// SetStatusMessage sets the "statusMessage" field.
func (u *ModuleUpsertBulk) SetStatusMessage(v string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetStatusMessage(v)
	})
}

// UpdateStatusMessage sets the "statusMessage" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateStatusMessage() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateStatusMessage()
	})
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (u *ModuleUpsertBulk) ClearStatusMessage() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearStatusMessage()
	})
}

// SetUpdateTime sets the "updateTime" field.
func (u *ModuleUpsertBulk) SetUpdateTime(v time.Time) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateUpdateTime() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *ModuleUpsertBulk) SetDescription(v string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateDescription() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ModuleUpsertBulk) ClearDescription() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *ModuleUpsertBulk) SetLabels(v map[string]string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateLabels() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateLabels()
	})
}

// SetSource sets the "source" field.
func (u *ModuleUpsertBulk) SetSource(v string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateSource() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateSource()
	})
}

// SetVersion sets the "version" field.
func (u *ModuleUpsertBulk) SetVersion(v string) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateVersion() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *ModuleUpsertBulk) ClearVersion() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.ClearVersion()
	})
}

// SetSchema sets the "schema" field.
func (u *ModuleUpsertBulk) SetSchema(v *types.ModuleSchema) *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.SetSchema(v)
	})
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *ModuleUpsertBulk) UpdateSchema() *ModuleUpsertBulk {
	return u.Update(func(s *ModuleUpsert) {
		s.UpdateSchema()
	})
}

// Exec executes the query.
func (u *ModuleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ModuleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ModuleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ModuleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

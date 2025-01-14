// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/token"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// SubjectUpdate is the builder for updating Subject entities.
type SubjectUpdate struct {
	config
	hooks     []Hook
	mutation  *SubjectMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Subject
}

// Where appends a list predicates to the SubjectUpdate builder.
func (su *SubjectUpdate) Where(ps ...predicate.Subject) *SubjectUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SubjectUpdate) SetUpdateTime(t time.Time) *SubjectUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetDescription sets the "description" field.
func (su *SubjectUpdate) SetDescription(s string) *SubjectUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SubjectUpdate) SetNillableDescription(s *string) *SubjectUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SubjectUpdate) ClearDescription() *SubjectUpdate {
	su.mutation.ClearDescription()
	return su
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (su *SubjectUpdate) AddTokenIDs(ids ...object.ID) *SubjectUpdate {
	su.mutation.AddTokenIDs(ids...)
	return su
}

// AddTokens adds the "tokens" edges to the Token entity.
func (su *SubjectUpdate) AddTokens(t ...*Token) *SubjectUpdate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTokenIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the SubjectRoleRelationship entity by IDs.
func (su *SubjectUpdate) AddRoleIDs(ids ...object.ID) *SubjectUpdate {
	su.mutation.AddRoleIDs(ids...)
	return su
}

// AddRoles adds the "roles" edges to the SubjectRoleRelationship entity.
func (su *SubjectUpdate) AddRoles(s ...*SubjectRoleRelationship) *SubjectUpdate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddRoleIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (su *SubjectUpdate) Mutation() *SubjectMutation {
	return su.mutation
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (su *SubjectUpdate) ClearTokens() *SubjectUpdate {
	su.mutation.ClearTokens()
	return su
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (su *SubjectUpdate) RemoveTokenIDs(ids ...object.ID) *SubjectUpdate {
	su.mutation.RemoveTokenIDs(ids...)
	return su
}

// RemoveTokens removes "tokens" edges to Token entities.
func (su *SubjectUpdate) RemoveTokens(t ...*Token) *SubjectUpdate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTokenIDs(ids...)
}

// ClearRoles clears all "roles" edges to the SubjectRoleRelationship entity.
func (su *SubjectUpdate) ClearRoles() *SubjectUpdate {
	su.mutation.ClearRoles()
	return su
}

// RemoveRoleIDs removes the "roles" edge to SubjectRoleRelationship entities by IDs.
func (su *SubjectUpdate) RemoveRoleIDs(ids ...object.ID) *SubjectUpdate {
	su.mutation.RemoveRoleIDs(ids...)
	return su
}

// RemoveRoles removes "roles" edges to SubjectRoleRelationship entities.
func (su *SubjectUpdate) RemoveRoles(s ...*SubjectRoleRelationship) *SubjectUpdate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubjectUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubjectUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubjectUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubjectUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SubjectUpdate) defaults() error {
	if _, ok := su.mutation.UpdateTime(); !ok {
		if subject.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized subject.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := subject.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (su *SubjectUpdate) Set(obj *Subject) *SubjectUpdate {
	// Without Default.
	if obj.Description != "" {
		su.SetDescription(obj.Description)
	} else {
		su.ClearDescription()
	}

	// With Default.
	if obj.UpdateTime != nil {
		su.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	su.object = obj

	return su
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SubjectUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubjectUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SubjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(subject.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(subject.FieldDescription, field.TypeString, value)
	}
	if su.mutation.DescriptionCleared() {
		_spec.ClearField(subject.FieldDescription, field.TypeString)
	}
	if su.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Token
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTokensIDs(); len(nodes) > 0 && !su.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.SubjectRoleRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedRolesIDs(); len(nodes) > 0 && !su.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = su.schemaConfig.Subject
	ctx = internal.NewSchemaConfigContext(ctx, su.schemaConfig)
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubjectUpdateOne is the builder for updating a single Subject entity.
type SubjectUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SubjectMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Subject
}

// SetUpdateTime sets the "update_time" field.
func (suo *SubjectUpdateOne) SetUpdateTime(t time.Time) *SubjectUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetDescription sets the "description" field.
func (suo *SubjectUpdateOne) SetDescription(s string) *SubjectUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableDescription(s *string) *SubjectUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SubjectUpdateOne) ClearDescription() *SubjectUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (suo *SubjectUpdateOne) AddTokenIDs(ids ...object.ID) *SubjectUpdateOne {
	suo.mutation.AddTokenIDs(ids...)
	return suo
}

// AddTokens adds the "tokens" edges to the Token entity.
func (suo *SubjectUpdateOne) AddTokens(t ...*Token) *SubjectUpdateOne {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTokenIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the SubjectRoleRelationship entity by IDs.
func (suo *SubjectUpdateOne) AddRoleIDs(ids ...object.ID) *SubjectUpdateOne {
	suo.mutation.AddRoleIDs(ids...)
	return suo
}

// AddRoles adds the "roles" edges to the SubjectRoleRelationship entity.
func (suo *SubjectUpdateOne) AddRoles(s ...*SubjectRoleRelationship) *SubjectUpdateOne {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddRoleIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (suo *SubjectUpdateOne) Mutation() *SubjectMutation {
	return suo.mutation
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (suo *SubjectUpdateOne) ClearTokens() *SubjectUpdateOne {
	suo.mutation.ClearTokens()
	return suo
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (suo *SubjectUpdateOne) RemoveTokenIDs(ids ...object.ID) *SubjectUpdateOne {
	suo.mutation.RemoveTokenIDs(ids...)
	return suo
}

// RemoveTokens removes "tokens" edges to Token entities.
func (suo *SubjectUpdateOne) RemoveTokens(t ...*Token) *SubjectUpdateOne {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTokenIDs(ids...)
}

// ClearRoles clears all "roles" edges to the SubjectRoleRelationship entity.
func (suo *SubjectUpdateOne) ClearRoles() *SubjectUpdateOne {
	suo.mutation.ClearRoles()
	return suo
}

// RemoveRoleIDs removes the "roles" edge to SubjectRoleRelationship entities by IDs.
func (suo *SubjectUpdateOne) RemoveRoleIDs(ids ...object.ID) *SubjectUpdateOne {
	suo.mutation.RemoveRoleIDs(ids...)
	return suo
}

// RemoveRoles removes "roles" edges to SubjectRoleRelationship entities.
func (suo *SubjectUpdateOne) RemoveRoles(s ...*SubjectRoleRelationship) *SubjectUpdateOne {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the SubjectUpdate builder.
func (suo *SubjectUpdateOne) Where(ps ...predicate.Subject) *SubjectUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubjectUpdateOne) Select(field string, fields ...string) *SubjectUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subject entity.
func (suo *SubjectUpdateOne) Save(ctx context.Context) (*Subject, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveX(ctx context.Context) *Subject {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubjectUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SubjectUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		if subject.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized subject.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := subject.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (suo *SubjectUpdateOne) Set(obj *Subject) *SubjectUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*SubjectMutation)
			db, err := mt.Client().Subject.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting Subject with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Description != "" {
				if db.Description != obj.Description {
					suo.SetDescription(obj.Description)
				}
			} else {
				suo.ClearDescription()
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				suo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			suo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	suo.hooks = append(suo.hooks, h)

	return suo
}

// getClientSet returns the ClientSet for the given builder.
func (suo *SubjectUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := suo.config.driver.(*txDriver); ok {
		tx := &Tx{config: suo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: suo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the Subject entity,
// which is always good for cascading update operations.
func (suo *SubjectUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) (*Subject, error) {
	obj, err := suo.Save(ctx)
	if err != nil &&
		(suo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := suo.getClientSet()

	if obj == nil {
		obj = suo.object
	} else if x := suo.object; x != nil {
		if _, set := suo.mutation.Field(subject.FieldDescription); set {
			obj.Description = x.Description
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) *Subject {
	obj, err := suo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (suo *SubjectUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) error {
	_, err := suo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) {
	if err := suo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SubjectUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubjectUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SubjectUpdateOne) sqlSave(ctx context.Context) (_node *Subject, err error) {
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Subject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for _, f := range fields {
			if !subject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != subject.FieldID {
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
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(subject.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(subject.FieldDescription, field.TypeString, value)
	}
	if suo.mutation.DescriptionCleared() {
		_spec.ClearField(subject.FieldDescription, field.TypeString)
	}
	if suo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Token
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !suo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.SubjectRoleRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !suo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = suo.schemaConfig.Subject
	ctx = internal.NewSchemaConfigContext(ctx, suo.schemaConfig)
	_spec.AddModifiers(suo.modifiers...)
	_node = &Subject{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

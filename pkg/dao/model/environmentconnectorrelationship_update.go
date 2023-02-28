// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// EnvironmentConnectorRelationshipUpdate is the builder for updating EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipUpdate struct {
	config
	hooks     []Hook
	mutation  *EnvironmentConnectorRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EnvironmentConnectorRelationshipUpdate builder.
func (ecru *EnvironmentConnectorRelationshipUpdate) Where(ps ...predicate.EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipUpdate {
	ecru.mutation.Where(ps...)
	return ecru
}

// Mutation returns the EnvironmentConnectorRelationshipMutation object of the builder.
func (ecru *EnvironmentConnectorRelationshipUpdate) Mutation() *EnvironmentConnectorRelationshipMutation {
	return ecru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecru *EnvironmentConnectorRelationshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EnvironmentConnectorRelationshipMutation](ctx, ecru.sqlSave, ecru.mutation, ecru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecru *EnvironmentConnectorRelationshipUpdate) SaveX(ctx context.Context) int {
	affected, err := ecru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecru *EnvironmentConnectorRelationshipUpdate) Exec(ctx context.Context) error {
	_, err := ecru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecru *EnvironmentConnectorRelationshipUpdate) ExecX(ctx context.Context) {
	if err := ecru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecru *EnvironmentConnectorRelationshipUpdate) check() error {
	if _, ok := ecru.mutation.EnvironmentID(); ecru.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "EnvironmentConnectorRelationship.environment"`)
	}
	if _, ok := ecru.mutation.ConnectorID(); ecru.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "EnvironmentConnectorRelationship.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ecru *EnvironmentConnectorRelationshipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvironmentConnectorRelationshipUpdate {
	ecru.modifiers = append(ecru.modifiers, modifiers...)
	return ecru
}

func (ecru *EnvironmentConnectorRelationshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ecru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(environmentconnectorrelationship.Table, environmentconnectorrelationship.Columns, sqlgraph.NewFieldSpec(environmentconnectorrelationship.FieldEnvironmentID, field.TypeString), sqlgraph.NewFieldSpec(environmentconnectorrelationship.FieldConnectorID, field.TypeString))
	if ps := ecru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = ecru.schemaConfig.EnvironmentConnectorRelationship
	ctx = internal.NewSchemaConfigContext(ctx, ecru.schemaConfig)
	_spec.AddModifiers(ecru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ecru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{environmentconnectorrelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecru.mutation.done = true
	return n, nil
}

// EnvironmentConnectorRelationshipUpdateOne is the builder for updating a single EnvironmentConnectorRelationship entity.
type EnvironmentConnectorRelationshipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EnvironmentConnectorRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Mutation returns the EnvironmentConnectorRelationshipMutation object of the builder.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Mutation() *EnvironmentConnectorRelationshipMutation {
	return ecruo.mutation
}

// Where appends a list predicates to the EnvironmentConnectorRelationshipUpdate builder.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Where(ps ...predicate.EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipUpdateOne {
	ecruo.mutation.Where(ps...)
	return ecruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Select(field string, fields ...string) *EnvironmentConnectorRelationshipUpdateOne {
	ecruo.fields = append([]string{field}, fields...)
	return ecruo
}

// Save executes the query and returns the updated EnvironmentConnectorRelationship entity.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Save(ctx context.Context) (*EnvironmentConnectorRelationship, error) {
	return withHooks[*EnvironmentConnectorRelationship, EnvironmentConnectorRelationshipMutation](ctx, ecruo.sqlSave, ecruo.mutation, ecruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) SaveX(ctx context.Context) *EnvironmentConnectorRelationship {
	node, err := ecruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Exec(ctx context.Context) error {
	_, err := ecruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) ExecX(ctx context.Context) {
	if err := ecruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) check() error {
	if _, ok := ecruo.mutation.EnvironmentID(); ecruo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "EnvironmentConnectorRelationship.environment"`)
	}
	if _, ok := ecruo.mutation.ConnectorID(); ecruo.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "EnvironmentConnectorRelationship.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ecruo *EnvironmentConnectorRelationshipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvironmentConnectorRelationshipUpdateOne {
	ecruo.modifiers = append(ecruo.modifiers, modifiers...)
	return ecruo
}

func (ecruo *EnvironmentConnectorRelationshipUpdateOne) sqlSave(ctx context.Context) (_node *EnvironmentConnectorRelationship, err error) {
	if err := ecruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(environmentconnectorrelationship.Table, environmentconnectorrelationship.Columns, sqlgraph.NewFieldSpec(environmentconnectorrelationship.FieldEnvironmentID, field.TypeString), sqlgraph.NewFieldSpec(environmentconnectorrelationship.FieldConnectorID, field.TypeString))
	if id, ok := ecruo.mutation.EnvironmentID(); !ok {
		return nil, &ValidationError{Name: "environment_id", err: errors.New(`model: missing "EnvironmentConnectorRelationship.environment_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ecruo.mutation.ConnectorID(); !ok {
		return nil, &ValidationError{Name: "connector_id", err: errors.New(`model: missing "EnvironmentConnectorRelationship.connector_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ecruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !environmentconnectorrelationship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ecruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = ecruo.schemaConfig.EnvironmentConnectorRelationship
	ctx = internal.NewSchemaConfigContext(ctx, ecruo.schemaConfig)
	_spec.AddModifiers(ecruo.modifiers...)
	_node = &EnvironmentConnectorRelationship{config: ecruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{environmentconnectorrelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecruo.mutation.done = true
	return _node, nil
}

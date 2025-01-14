// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponentrelationship"
)

// ResourceComponentRelationshipDelete is the builder for deleting a ResourceComponentRelationship entity.
type ResourceComponentRelationshipDelete struct {
	config
	hooks    []Hook
	mutation *ResourceComponentRelationshipMutation
}

// Where appends a list predicates to the ResourceComponentRelationshipDelete builder.
func (rcrd *ResourceComponentRelationshipDelete) Where(ps ...predicate.ResourceComponentRelationship) *ResourceComponentRelationshipDelete {
	rcrd.mutation.Where(ps...)
	return rcrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcrd *ResourceComponentRelationshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rcrd.sqlExec, rcrd.mutation, rcrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rcrd *ResourceComponentRelationshipDelete) ExecX(ctx context.Context) int {
	n, err := rcrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcrd *ResourceComponentRelationshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(resourcecomponentrelationship.Table, sqlgraph.NewFieldSpec(resourcecomponentrelationship.FieldID, field.TypeString))
	_spec.Node.Schema = rcrd.schemaConfig.ResourceComponentRelationship
	ctx = internal.NewSchemaConfigContext(ctx, rcrd.schemaConfig)
	if ps := rcrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rcrd.mutation.done = true
	return affected, err
}

// ResourceComponentRelationshipDeleteOne is the builder for deleting a single ResourceComponentRelationship entity.
type ResourceComponentRelationshipDeleteOne struct {
	rcrd *ResourceComponentRelationshipDelete
}

// Where appends a list predicates to the ResourceComponentRelationshipDelete builder.
func (rcrdo *ResourceComponentRelationshipDeleteOne) Where(ps ...predicate.ResourceComponentRelationship) *ResourceComponentRelationshipDeleteOne {
	rcrdo.rcrd.mutation.Where(ps...)
	return rcrdo
}

// Exec executes the deletion query.
func (rcrdo *ResourceComponentRelationshipDeleteOne) Exec(ctx context.Context) error {
	n, err := rcrdo.rcrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{resourcecomponentrelationship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcrdo *ResourceComponentRelationshipDeleteOne) ExecX(ctx context.Context) {
	if err := rcrdo.Exec(ctx); err != nil {
		panic(err)
	}
}

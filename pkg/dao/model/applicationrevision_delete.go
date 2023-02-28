// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/applicationrevision"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ApplicationRevisionDelete is the builder for deleting a ApplicationRevision entity.
type ApplicationRevisionDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationRevisionMutation
}

// Where appends a list predicates to the ApplicationRevisionDelete builder.
func (ard *ApplicationRevisionDelete) Where(ps ...predicate.ApplicationRevision) *ApplicationRevisionDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *ApplicationRevisionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ApplicationRevisionMutation](ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *ApplicationRevisionDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *ApplicationRevisionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(applicationrevision.Table, sqlgraph.NewFieldSpec(applicationrevision.FieldID, field.TypeString))
	_spec.Node.Schema = ard.schemaConfig.ApplicationRevision
	ctx = internal.NewSchemaConfigContext(ctx, ard.schemaConfig)
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// ApplicationRevisionDeleteOne is the builder for deleting a single ApplicationRevision entity.
type ApplicationRevisionDeleteOne struct {
	ard *ApplicationRevisionDelete
}

// Where appends a list predicates to the ApplicationRevisionDelete builder.
func (ardo *ApplicationRevisionDeleteOne) Where(ps ...predicate.ApplicationRevision) *ApplicationRevisionDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *ApplicationRevisionDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{applicationrevision.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *ApplicationRevisionDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}

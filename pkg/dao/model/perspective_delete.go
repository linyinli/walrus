// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/perspective"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// PerspectiveDelete is the builder for deleting a Perspective entity.
type PerspectiveDelete struct {
	config
	hooks    []Hook
	mutation *PerspectiveMutation
}

// Where appends a list predicates to the PerspectiveDelete builder.
func (pd *PerspectiveDelete) Where(ps ...predicate.Perspective) *PerspectiveDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PerspectiveDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PerspectiveMutation](ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PerspectiveDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PerspectiveDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(perspective.Table, sqlgraph.NewFieldSpec(perspective.FieldID, field.TypeString))
	_spec.Node.Schema = pd.schemaConfig.Perspective
	ctx = internal.NewSchemaConfigContext(ctx, pd.schemaConfig)
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PerspectiveDeleteOne is the builder for deleting a single Perspective entity.
type PerspectiveDeleteOne struct {
	pd *PerspectiveDelete
}

// Where appends a list predicates to the PerspectiveDelete builder.
func (pdo *PerspectiveDeleteOne) Where(ps ...predicate.Perspective) *PerspectiveDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PerspectiveDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{perspective.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PerspectiveDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}

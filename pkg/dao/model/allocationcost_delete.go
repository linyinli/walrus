// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// AllocationCostDelete is the builder for deleting a AllocationCost entity.
type AllocationCostDelete struct {
	config
	hooks    []Hook
	mutation *AllocationCostMutation
}

// Where appends a list predicates to the AllocationCostDelete builder.
func (acd *AllocationCostDelete) Where(ps ...predicate.AllocationCost) *AllocationCostDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AllocationCostDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AllocationCostMutation](ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AllocationCostDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AllocationCostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(allocationcost.Table, sqlgraph.NewFieldSpec(allocationcost.FieldID, field.TypeInt))
	_spec.Node.Schema = acd.schemaConfig.AllocationCost
	ctx = internal.NewSchemaConfigContext(ctx, acd.schemaConfig)
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AllocationCostDeleteOne is the builder for deleting a single AllocationCost entity.
type AllocationCostDeleteOne struct {
	acd *AllocationCostDelete
}

// Where appends a list predicates to the AllocationCostDelete builder.
func (acdo *AllocationCostDeleteOne) Where(ps ...predicate.AllocationCost) *AllocationCostDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AllocationCostDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{allocationcost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AllocationCostDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}

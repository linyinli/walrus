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
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
)

// SubjectRoleRelationshipDelete is the builder for deleting a SubjectRoleRelationship entity.
type SubjectRoleRelationshipDelete struct {
	config
	hooks    []Hook
	mutation *SubjectRoleRelationshipMutation
}

// Where appends a list predicates to the SubjectRoleRelationshipDelete builder.
func (srrd *SubjectRoleRelationshipDelete) Where(ps ...predicate.SubjectRoleRelationship) *SubjectRoleRelationshipDelete {
	srrd.mutation.Where(ps...)
	return srrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srrd *SubjectRoleRelationshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srrd.sqlExec, srrd.mutation, srrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srrd *SubjectRoleRelationshipDelete) ExecX(ctx context.Context) int {
	n, err := srrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srrd *SubjectRoleRelationshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subjectrolerelationship.Table, sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString))
	_spec.Node.Schema = srrd.schemaConfig.SubjectRoleRelationship
	ctx = internal.NewSchemaConfigContext(ctx, srrd.schemaConfig)
	if ps := srrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srrd.mutation.done = true
	return affected, err
}

// SubjectRoleRelationshipDeleteOne is the builder for deleting a single SubjectRoleRelationship entity.
type SubjectRoleRelationshipDeleteOne struct {
	srrd *SubjectRoleRelationshipDelete
}

// Where appends a list predicates to the SubjectRoleRelationshipDelete builder.
func (srrdo *SubjectRoleRelationshipDeleteOne) Where(ps ...predicate.SubjectRoleRelationship) *SubjectRoleRelationshipDeleteOne {
	srrdo.srrd.mutation.Where(ps...)
	return srrdo
}

// Exec executes the deletion query.
func (srrdo *SubjectRoleRelationshipDeleteOne) Exec(ctx context.Context) error {
	n, err := srrdo.srrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subjectrolerelationship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srrdo *SubjectRoleRelationshipDeleteOne) ExecX(ctx context.Context) {
	if err := srrdo.Exec(ctx); err != nil {
		panic(err)
	}
}

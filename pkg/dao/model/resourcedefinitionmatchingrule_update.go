// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"bytes"
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
)

// ResourceDefinitionMatchingRuleUpdate is the builder for updating ResourceDefinitionMatchingRule entities.
type ResourceDefinitionMatchingRuleUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceDefinitionMatchingRuleMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceDefinitionMatchingRule
}

// Where appends a list predicates to the ResourceDefinitionMatchingRuleUpdate builder.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Where(ps ...predicate.ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.Where(ps...)
	return rdmru
}

// SetTemplateID sets the "template_id" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetTemplateID(o object.ID) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.SetTemplateID(o)
	return rdmru
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetNillableTemplateID(o *object.ID) *ResourceDefinitionMatchingRuleUpdate {
	if o != nil {
		rdmru.SetTemplateID(*o)
	}
	return rdmru
}

// SetSelector sets the "selector" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetSelector(t types.Selector) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.SetSelector(t)
	return rdmru
}

// SetNillableSelector sets the "selector" field if the given value is not nil.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetNillableSelector(t *types.Selector) *ResourceDefinitionMatchingRuleUpdate {
	if t != nil {
		rdmru.SetSelector(*t)
	}
	return rdmru
}

// SetAttributes sets the "attributes" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetAttributes(pr property.Values) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.SetAttributes(pr)
	return rdmru
}

// ClearAttributes clears the value of the "attributes" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) ClearAttributes() *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.ClearAttributes()
	return rdmru
}

// SetOrder sets the "order" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetOrder(i int) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.ResetOrder()
	rdmru.mutation.SetOrder(i)
	return rdmru
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetNillableOrder(i *int) *ResourceDefinitionMatchingRuleUpdate {
	if i != nil {
		rdmru.SetOrder(*i)
	}
	return rdmru
}

// AddOrder adds i to the "order" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) AddOrder(i int) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.AddOrder(i)
	return rdmru
}

// SetSchemaDefaultValue sets the "schema_default_value" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetSchemaDefaultValue(b []byte) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.SetSchemaDefaultValue(b)
	return rdmru
}

// ClearSchemaDefaultValue clears the value of the "schema_default_value" field.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) ClearSchemaDefaultValue() *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.ClearSchemaDefaultValue()
	return rdmru
}

// SetTemplate sets the "template" edge to the TemplateVersion entity.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SetTemplate(t *TemplateVersion) *ResourceDefinitionMatchingRuleUpdate {
	return rdmru.SetTemplateID(t.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) AddResourceIDs(ids ...object.ID) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.AddResourceIDs(ids...)
	return rdmru
}

// AddResources adds the "resources" edges to the Resource entity.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) AddResources(r ...*Resource) *ResourceDefinitionMatchingRuleUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdmru.AddResourceIDs(ids...)
}

// Mutation returns the ResourceDefinitionMatchingRuleMutation object of the builder.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Mutation() *ResourceDefinitionMatchingRuleMutation {
	return rdmru.mutation
}

// ClearTemplate clears the "template" edge to the TemplateVersion entity.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) ClearTemplate() *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.ClearTemplate()
	return rdmru
}

// ClearResources clears all "resources" edges to the Resource entity.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) ClearResources() *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.ClearResources()
	return rdmru
}

// RemoveResourceIDs removes the "resources" edge to Resource entities by IDs.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) RemoveResourceIDs(ids ...object.ID) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.mutation.RemoveResourceIDs(ids...)
	return rdmru
}

// RemoveResources removes "resources" edges to Resource entities.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) RemoveResources(r ...*Resource) *ResourceDefinitionMatchingRuleUpdate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdmru.RemoveResourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rdmru.sqlSave, rdmru.mutation, rdmru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := rdmru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Exec(ctx context.Context) error {
	_, err := rdmru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) ExecX(ctx context.Context) {
	if err := rdmru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) check() error {
	if v, ok := rdmru.mutation.TemplateID(); ok {
		if err := resourcedefinitionmatchingrule.TemplateIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "template_id", err: fmt.Errorf(`model: validator failed for field "ResourceDefinitionMatchingRule.template_id": %w`, err)}
		}
	}
	if _, ok := rdmru.mutation.ResourceDefinitionID(); rdmru.mutation.ResourceDefinitionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceDefinitionMatchingRule.resource_definition"`)
	}
	if _, ok := rdmru.mutation.TemplateID(); rdmru.mutation.TemplateCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceDefinitionMatchingRule.template"`)
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
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Set(obj *ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleUpdate {
	// Without Default.
	rdmru.SetTemplateID(obj.TemplateID)
	rdmru.SetSelector(obj.Selector)
	if !reflect.ValueOf(obj.Attributes).IsZero() {
		rdmru.SetAttributes(obj.Attributes)
	}
	rdmru.SetOrder(obj.Order)
	if !reflect.ValueOf(obj.SchemaDefaultValue).IsZero() {
		rdmru.SetSchemaDefaultValue(obj.SchemaDefaultValue)
	}

	// With Default.

	// Record the given object.
	rdmru.object = obj

	return rdmru
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rdmru *ResourceDefinitionMatchingRuleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceDefinitionMatchingRuleUpdate {
	rdmru.modifiers = append(rdmru.modifiers, modifiers...)
	return rdmru
}

func (rdmru *ResourceDefinitionMatchingRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rdmru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcedefinitionmatchingrule.Table, resourcedefinitionmatchingrule.Columns, sqlgraph.NewFieldSpec(resourcedefinitionmatchingrule.FieldID, field.TypeString))
	if ps := rdmru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdmru.mutation.Selector(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldSelector, field.TypeJSON, value)
	}
	if value, ok := rdmru.mutation.Attributes(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldAttributes, field.TypeOther, value)
	}
	if rdmru.mutation.AttributesCleared() {
		_spec.ClearField(resourcedefinitionmatchingrule.FieldAttributes, field.TypeOther)
	}
	if value, ok := rdmru.mutation.Order(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rdmru.mutation.AddedOrder(); ok {
		_spec.AddField(resourcedefinitionmatchingrule.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rdmru.mutation.SchemaDefaultValue(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldSchemaDefaultValue, field.TypeBytes, value)
	}
	if rdmru.mutation.SchemaDefaultValueCleared() {
		_spec.ClearField(resourcedefinitionmatchingrule.FieldSchemaDefaultValue, field.TypeBytes)
	}
	if rdmru.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.TemplateTable,
			Columns: []string{resourcedefinitionmatchingrule.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmru.schemaConfig.ResourceDefinitionMatchingRule
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmru.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.TemplateTable,
			Columns: []string{resourcedefinitionmatchingrule.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmru.schemaConfig.ResourceDefinitionMatchingRule
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdmru.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmru.schemaConfig.Resource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmru.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !rdmru.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmru.schemaConfig.Resource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmru.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmru.schemaConfig.Resource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = rdmru.schemaConfig.ResourceDefinitionMatchingRule
	ctx = internal.NewSchemaConfigContext(ctx, rdmru.schemaConfig)
	_spec.AddModifiers(rdmru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rdmru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcedefinitionmatchingrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rdmru.mutation.done = true
	return n, nil
}

// ResourceDefinitionMatchingRuleUpdateOne is the builder for updating a single ResourceDefinitionMatchingRule entity.
type ResourceDefinitionMatchingRuleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceDefinitionMatchingRuleMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceDefinitionMatchingRule
}

// SetTemplateID sets the "template_id" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetTemplateID(o object.ID) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.SetTemplateID(o)
	return rdmruo
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetNillableTemplateID(o *object.ID) *ResourceDefinitionMatchingRuleUpdateOne {
	if o != nil {
		rdmruo.SetTemplateID(*o)
	}
	return rdmruo
}

// SetSelector sets the "selector" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetSelector(t types.Selector) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.SetSelector(t)
	return rdmruo
}

// SetNillableSelector sets the "selector" field if the given value is not nil.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetNillableSelector(t *types.Selector) *ResourceDefinitionMatchingRuleUpdateOne {
	if t != nil {
		rdmruo.SetSelector(*t)
	}
	return rdmruo
}

// SetAttributes sets the "attributes" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetAttributes(pr property.Values) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.SetAttributes(pr)
	return rdmruo
}

// ClearAttributes clears the value of the "attributes" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ClearAttributes() *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.ClearAttributes()
	return rdmruo
}

// SetOrder sets the "order" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetOrder(i int) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.ResetOrder()
	rdmruo.mutation.SetOrder(i)
	return rdmruo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetNillableOrder(i *int) *ResourceDefinitionMatchingRuleUpdateOne {
	if i != nil {
		rdmruo.SetOrder(*i)
	}
	return rdmruo
}

// AddOrder adds i to the "order" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) AddOrder(i int) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.AddOrder(i)
	return rdmruo
}

// SetSchemaDefaultValue sets the "schema_default_value" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetSchemaDefaultValue(b []byte) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.SetSchemaDefaultValue(b)
	return rdmruo
}

// ClearSchemaDefaultValue clears the value of the "schema_default_value" field.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ClearSchemaDefaultValue() *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.ClearSchemaDefaultValue()
	return rdmruo
}

// SetTemplate sets the "template" edge to the TemplateVersion entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SetTemplate(t *TemplateVersion) *ResourceDefinitionMatchingRuleUpdateOne {
	return rdmruo.SetTemplateID(t.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) AddResourceIDs(ids ...object.ID) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.AddResourceIDs(ids...)
	return rdmruo
}

// AddResources adds the "resources" edges to the Resource entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) AddResources(r ...*Resource) *ResourceDefinitionMatchingRuleUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdmruo.AddResourceIDs(ids...)
}

// Mutation returns the ResourceDefinitionMatchingRuleMutation object of the builder.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Mutation() *ResourceDefinitionMatchingRuleMutation {
	return rdmruo.mutation
}

// ClearTemplate clears the "template" edge to the TemplateVersion entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ClearTemplate() *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.ClearTemplate()
	return rdmruo
}

// ClearResources clears all "resources" edges to the Resource entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ClearResources() *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.ClearResources()
	return rdmruo
}

// RemoveResourceIDs removes the "resources" edge to Resource entities by IDs.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) RemoveResourceIDs(ids ...object.ID) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.RemoveResourceIDs(ids...)
	return rdmruo
}

// RemoveResources removes "resources" edges to Resource entities.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) RemoveResources(r ...*Resource) *ResourceDefinitionMatchingRuleUpdateOne {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdmruo.RemoveResourceIDs(ids...)
}

// Where appends a list predicates to the ResourceDefinitionMatchingRuleUpdate builder.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Where(ps ...predicate.ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.mutation.Where(ps...)
	return rdmruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Select(field string, fields ...string) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.fields = append([]string{field}, fields...)
	return rdmruo
}

// Save executes the query and returns the updated ResourceDefinitionMatchingRule entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Save(ctx context.Context) (*ResourceDefinitionMatchingRule, error) {
	return withHooks(ctx, rdmruo.sqlSave, rdmruo.mutation, rdmruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SaveX(ctx context.Context) *ResourceDefinitionMatchingRule {
	node, err := rdmruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := rdmruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ExecX(ctx context.Context) {
	if err := rdmruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) check() error {
	if v, ok := rdmruo.mutation.TemplateID(); ok {
		if err := resourcedefinitionmatchingrule.TemplateIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "template_id", err: fmt.Errorf(`model: validator failed for field "ResourceDefinitionMatchingRule.template_id": %w`, err)}
		}
	}
	if _, ok := rdmruo.mutation.ResourceDefinitionID(); rdmruo.mutation.ResourceDefinitionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceDefinitionMatchingRule.resource_definition"`)
	}
	if _, ok := rdmruo.mutation.TemplateID(); rdmruo.mutation.TemplateCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceDefinitionMatchingRule.template"`)
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
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Set(obj *ResourceDefinitionMatchingRule) *ResourceDefinitionMatchingRuleUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ResourceDefinitionMatchingRuleMutation)
			db, err := mt.Client().ResourceDefinitionMatchingRule.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ResourceDefinitionMatchingRule with id: %v", *mt.id)
			}

			// Without Default.
			if db.TemplateID != obj.TemplateID {
				rdmruo.SetTemplateID(obj.TemplateID)
			}
			if !reflect.DeepEqual(db.Selector, obj.Selector) {
				rdmruo.SetSelector(obj.Selector)
			}
			if !reflect.ValueOf(obj.Attributes).IsZero() {
				if !reflect.DeepEqual(db.Attributes, obj.Attributes) {
					rdmruo.SetAttributes(obj.Attributes)
				}
			}
			if db.Order != obj.Order {
				rdmruo.SetOrder(obj.Order)
			}
			if !reflect.ValueOf(obj.SchemaDefaultValue).IsZero() {
				if !bytes.Equal(db.SchemaDefaultValue, obj.SchemaDefaultValue) {
					rdmruo.SetSchemaDefaultValue(obj.SchemaDefaultValue)
				}
			}

			// With Default.

			// Record the given object.
			rdmruo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	rdmruo.hooks = append(rdmruo.hooks, h)

	return rdmruo
}

// getClientSet returns the ClientSet for the given builder.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := rdmruo.config.driver.(*txDriver); ok {
		tx := &Tx{config: rdmruo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rdmruo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ResourceDefinitionMatchingRule entity,
// which is always good for cascading update operations.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceDefinitionMatchingRule) error) (*ResourceDefinitionMatchingRule, error) {
	obj, err := rdmruo.Save(ctx)
	if err != nil &&
		(rdmruo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rdmruo.getClientSet()

	if obj == nil {
		obj = rdmruo.object
	} else if x := rdmruo.object; x != nil {
		if _, set := rdmruo.mutation.Field(resourcedefinitionmatchingrule.FieldTemplateID); set {
			obj.TemplateID = x.TemplateID
		}
		if _, set := rdmruo.mutation.Field(resourcedefinitionmatchingrule.FieldSelector); set {
			obj.Selector = x.Selector
		}
		if _, set := rdmruo.mutation.Field(resourcedefinitionmatchingrule.FieldAttributes); set {
			obj.Attributes = x.Attributes
		}
		if _, set := rdmruo.mutation.Field(resourcedefinitionmatchingrule.FieldOrder); set {
			obj.Order = x.Order
		}
		if _, set := rdmruo.mutation.Field(resourcedefinitionmatchingrule.FieldSchemaDefaultValue); set {
			obj.SchemaDefaultValue = x.SchemaDefaultValue
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
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceDefinitionMatchingRule) error) *ResourceDefinitionMatchingRule {
	obj, err := rdmruo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceDefinitionMatchingRule) error) error {
	_, err := rdmruo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceDefinitionMatchingRule) error) {
	if err := rdmruo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceDefinitionMatchingRuleUpdateOne {
	rdmruo.modifiers = append(rdmruo.modifiers, modifiers...)
	return rdmruo
}

func (rdmruo *ResourceDefinitionMatchingRuleUpdateOne) sqlSave(ctx context.Context) (_node *ResourceDefinitionMatchingRule, err error) {
	if err := rdmruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcedefinitionmatchingrule.Table, resourcedefinitionmatchingrule.Columns, sqlgraph.NewFieldSpec(resourcedefinitionmatchingrule.FieldID, field.TypeString))
	id, ok := rdmruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ResourceDefinitionMatchingRule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rdmruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcedefinitionmatchingrule.FieldID)
		for _, f := range fields {
			if !resourcedefinitionmatchingrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != resourcedefinitionmatchingrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rdmruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdmruo.mutation.Selector(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldSelector, field.TypeJSON, value)
	}
	if value, ok := rdmruo.mutation.Attributes(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldAttributes, field.TypeOther, value)
	}
	if rdmruo.mutation.AttributesCleared() {
		_spec.ClearField(resourcedefinitionmatchingrule.FieldAttributes, field.TypeOther)
	}
	if value, ok := rdmruo.mutation.Order(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rdmruo.mutation.AddedOrder(); ok {
		_spec.AddField(resourcedefinitionmatchingrule.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rdmruo.mutation.SchemaDefaultValue(); ok {
		_spec.SetField(resourcedefinitionmatchingrule.FieldSchemaDefaultValue, field.TypeBytes, value)
	}
	if rdmruo.mutation.SchemaDefaultValueCleared() {
		_spec.ClearField(resourcedefinitionmatchingrule.FieldSchemaDefaultValue, field.TypeBytes)
	}
	if rdmruo.mutation.TemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.TemplateTable,
			Columns: []string{resourcedefinitionmatchingrule.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmruo.schemaConfig.ResourceDefinitionMatchingRule
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmruo.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.TemplateTable,
			Columns: []string{resourcedefinitionmatchingrule.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmruo.schemaConfig.ResourceDefinitionMatchingRule
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdmruo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmruo.schemaConfig.Resource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmruo.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !rdmruo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmruo.schemaConfig.Resource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdmruo.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcedefinitionmatchingrule.ResourcesTable,
			Columns: []string{resourcedefinitionmatchingrule.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rdmruo.schemaConfig.Resource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = rdmruo.schemaConfig.ResourceDefinitionMatchingRule
	ctx = internal.NewSchemaConfigContext(ctx, rdmruo.schemaConfig)
	_spec.AddModifiers(rdmruo.modifiers...)
	_node = &ResourceDefinitionMatchingRule{config: rdmruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rdmruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcedefinitionmatchingrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rdmruo.mutation.done = true
	return _node, nil
}

// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package resourcedefinitionmatchingrule

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"
)

const (
	// Label holds the string label denoting the resourcedefinitionmatchingrule type in the database.
	Label = "resource_definition_matching_rule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldResourceDefinitionID holds the string denoting the resource_definition_id field in the database.
	FieldResourceDefinitionID = "resource_definition_id"
	// FieldTemplateID holds the string denoting the template_id field in the database.
	FieldTemplateID = "template_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSelector holds the string denoting the selector field in the database.
	FieldSelector = "selector"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldSchemaDefaultValue holds the string denoting the schema_default_value field in the database.
	FieldSchemaDefaultValue = "schema_default_value"
	// EdgeResourceDefinition holds the string denoting the resource_definition edge name in mutations.
	EdgeResourceDefinition = "resource_definition"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// Table holds the table name of the resourcedefinitionmatchingrule in the database.
	Table = "resource_definition_matching_rules"
	// ResourceDefinitionTable is the table that holds the resource_definition relation/edge.
	ResourceDefinitionTable = "resource_definition_matching_rules"
	// ResourceDefinitionInverseTable is the table name for the ResourceDefinition entity.
	// It exists in this package in order to avoid circular dependency with the "resourcedefinition" package.
	ResourceDefinitionInverseTable = "resource_definitions"
	// ResourceDefinitionColumn is the table column denoting the resource_definition relation/edge.
	ResourceDefinitionColumn = "resource_definition_id"
	// TemplateTable is the table that holds the template relation/edge.
	TemplateTable = "resource_definition_matching_rules"
	// TemplateInverseTable is the table name for the TemplateVersion entity.
	// It exists in this package in order to avoid circular dependency with the "templateversion" package.
	TemplateInverseTable = "template_versions"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "template_id"
	// ResourcesTable is the table that holds the resources relation/edge.
	ResourcesTable = "resources"
	// ResourcesInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourcesInverseTable = "resources"
	// ResourcesColumn is the table column denoting the resources relation/edge.
	ResourcesColumn = "resource_definition_matching_rule_id"
)

// Columns holds all SQL columns for resourcedefinitionmatchingrule fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldResourceDefinitionID,
	FieldTemplateID,
	FieldName,
	FieldSelector,
	FieldAttributes,
	FieldOrder,
	FieldSchemaDefaultValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/walrus/pkg/dao/model/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// ResourceDefinitionIDValidator is a validator for the "resource_definition_id" field. It is called by the builders before save.
	ResourceDefinitionIDValidator func(string) error
	// TemplateIDValidator is a validator for the "template_id" field. It is called by the builders before save.
	TemplateIDValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the ResourceDefinitionMatchingRule queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByResourceDefinitionID orders the results by the resource_definition_id field.
func ByResourceDefinitionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResourceDefinitionID, opts...).ToFunc()
}

// ByTemplateID orders the results by the template_id field.
func ByTemplateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAttributes orders the results by the attributes field.
func ByAttributes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttributes, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByResourceDefinitionField orders the results by resource_definition field.
func ByResourceDefinitionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourceDefinitionStep(), sql.OrderByField(field, opts...))
	}
}

// ByTemplateField orders the results by template field.
func ByTemplateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTemplateStep(), sql.OrderByField(field, opts...))
	}
}

// ByResourcesCount orders the results by resources count.
func ByResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResourcesStep(), opts...)
	}
}

// ByResources orders the results by resources terms.
func ByResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newResourceDefinitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourceDefinitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ResourceDefinitionTable, ResourceDefinitionColumn),
	)
}
func newTemplateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TemplateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TemplateTable, TemplateColumn),
	)
}
func newResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
	)
}

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return slices.Clone(Columns)
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}

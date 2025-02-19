// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// EnvironmentConnectorRelationshipQuery is the builder for querying EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipQuery struct {
	config
	ctx             *QueryContext
	order           []environmentconnectorrelationship.OrderOption
	inters          []Interceptor
	predicates      []predicate.EnvironmentConnectorRelationship
	withEnvironment *EnvironmentQuery
	withConnector   *ConnectorQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnvironmentConnectorRelationshipQuery builder.
func (ecrq *EnvironmentConnectorRelationshipQuery) Where(ps ...predicate.EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipQuery {
	ecrq.predicates = append(ecrq.predicates, ps...)
	return ecrq
}

// Limit the number of records to be returned by this query.
func (ecrq *EnvironmentConnectorRelationshipQuery) Limit(limit int) *EnvironmentConnectorRelationshipQuery {
	ecrq.ctx.Limit = &limit
	return ecrq
}

// Offset to start from.
func (ecrq *EnvironmentConnectorRelationshipQuery) Offset(offset int) *EnvironmentConnectorRelationshipQuery {
	ecrq.ctx.Offset = &offset
	return ecrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecrq *EnvironmentConnectorRelationshipQuery) Unique(unique bool) *EnvironmentConnectorRelationshipQuery {
	ecrq.ctx.Unique = &unique
	return ecrq
}

// Order specifies how the records should be ordered.
func (ecrq *EnvironmentConnectorRelationshipQuery) Order(o ...environmentconnectorrelationship.OrderOption) *EnvironmentConnectorRelationshipQuery {
	ecrq.order = append(ecrq.order, o...)
	return ecrq
}

// QueryEnvironment chains the current query on the "environment" edge.
func (ecrq *EnvironmentConnectorRelationshipQuery) QueryEnvironment() *EnvironmentQuery {
	query := (&EnvironmentClient{config: ecrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environmentconnectorrelationship.Table, environmentconnectorrelationship.FieldID, selector),
			sqlgraph.To(environment.Table, environment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, environmentconnectorrelationship.EnvironmentTable, environmentconnectorrelationship.EnvironmentColumn),
		)
		schemaConfig := ecrq.schemaConfig
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.EnvironmentConnectorRelationship
		fromU = sqlgraph.SetNeighbors(ecrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConnector chains the current query on the "connector" edge.
func (ecrq *EnvironmentConnectorRelationshipQuery) QueryConnector() *ConnectorQuery {
	query := (&ConnectorClient{config: ecrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environmentconnectorrelationship.Table, environmentconnectorrelationship.FieldID, selector),
			sqlgraph.To(connector.Table, connector.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, environmentconnectorrelationship.ConnectorTable, environmentconnectorrelationship.ConnectorColumn),
		)
		schemaConfig := ecrq.schemaConfig
		step.To.Schema = schemaConfig.Connector
		step.Edge.Schema = schemaConfig.EnvironmentConnectorRelationship
		fromU = sqlgraph.SetNeighbors(ecrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EnvironmentConnectorRelationship entity from the query.
// Returns a *NotFoundError when no EnvironmentConnectorRelationship was found.
func (ecrq *EnvironmentConnectorRelationshipQuery) First(ctx context.Context) (*EnvironmentConnectorRelationship, error) {
	nodes, err := ecrq.Limit(1).All(setContextOp(ctx, ecrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{environmentconnectorrelationship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) FirstX(ctx context.Context) *EnvironmentConnectorRelationship {
	node, err := ecrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EnvironmentConnectorRelationship ID from the query.
// Returns a *NotFoundError when no EnvironmentConnectorRelationship ID was found.
func (ecrq *EnvironmentConnectorRelationshipQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = ecrq.Limit(1).IDs(setContextOp(ctx, ecrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{environmentconnectorrelationship.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := ecrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EnvironmentConnectorRelationship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EnvironmentConnectorRelationship entity is found.
// Returns a *NotFoundError when no EnvironmentConnectorRelationship entities are found.
func (ecrq *EnvironmentConnectorRelationshipQuery) Only(ctx context.Context) (*EnvironmentConnectorRelationship, error) {
	nodes, err := ecrq.Limit(2).All(setContextOp(ctx, ecrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{environmentconnectorrelationship.Label}
	default:
		return nil, &NotSingularError{environmentconnectorrelationship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) OnlyX(ctx context.Context) *EnvironmentConnectorRelationship {
	node, err := ecrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EnvironmentConnectorRelationship ID in the query.
// Returns a *NotSingularError when more than one EnvironmentConnectorRelationship ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecrq *EnvironmentConnectorRelationshipQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = ecrq.Limit(2).IDs(setContextOp(ctx, ecrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{environmentconnectorrelationship.Label}
	default:
		err = &NotSingularError{environmentconnectorrelationship.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := ecrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EnvironmentConnectorRelationships.
func (ecrq *EnvironmentConnectorRelationshipQuery) All(ctx context.Context) ([]*EnvironmentConnectorRelationship, error) {
	ctx = setContextOp(ctx, ecrq.ctx, "All")
	if err := ecrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EnvironmentConnectorRelationship, *EnvironmentConnectorRelationshipQuery]()
	return withInterceptors[[]*EnvironmentConnectorRelationship](ctx, ecrq, qr, ecrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) AllX(ctx context.Context) []*EnvironmentConnectorRelationship {
	nodes, err := ecrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EnvironmentConnectorRelationship IDs.
func (ecrq *EnvironmentConnectorRelationshipQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if ecrq.ctx.Unique == nil && ecrq.path != nil {
		ecrq.Unique(true)
	}
	ctx = setContextOp(ctx, ecrq.ctx, "IDs")
	if err = ecrq.Select(environmentconnectorrelationship.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := ecrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecrq *EnvironmentConnectorRelationshipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecrq.ctx, "Count")
	if err := ecrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecrq, querierCount[*EnvironmentConnectorRelationshipQuery](), ecrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) CountX(ctx context.Context) int {
	count, err := ecrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecrq *EnvironmentConnectorRelationshipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecrq.ctx, "Exist")
	switch _, err := ecrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecrq *EnvironmentConnectorRelationshipQuery) ExistX(ctx context.Context) bool {
	exist, err := ecrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnvironmentConnectorRelationshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecrq *EnvironmentConnectorRelationshipQuery) Clone() *EnvironmentConnectorRelationshipQuery {
	if ecrq == nil {
		return nil
	}
	return &EnvironmentConnectorRelationshipQuery{
		config:          ecrq.config,
		ctx:             ecrq.ctx.Clone(),
		order:           append([]environmentconnectorrelationship.OrderOption{}, ecrq.order...),
		inters:          append([]Interceptor{}, ecrq.inters...),
		predicates:      append([]predicate.EnvironmentConnectorRelationship{}, ecrq.predicates...),
		withEnvironment: ecrq.withEnvironment.Clone(),
		withConnector:   ecrq.withConnector.Clone(),
		// clone intermediate query.
		sql:  ecrq.sql.Clone(),
		path: ecrq.path,
	}
}

// WithEnvironment tells the query-builder to eager-load the nodes that are connected to
// the "environment" edge. The optional arguments are used to configure the query builder of the edge.
func (ecrq *EnvironmentConnectorRelationshipQuery) WithEnvironment(opts ...func(*EnvironmentQuery)) *EnvironmentConnectorRelationshipQuery {
	query := (&EnvironmentClient{config: ecrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecrq.withEnvironment = query
	return ecrq
}

// WithConnector tells the query-builder to eager-load the nodes that are connected to
// the "connector" edge. The optional arguments are used to configure the query builder of the edge.
func (ecrq *EnvironmentConnectorRelationshipQuery) WithConnector(opts ...func(*ConnectorQuery)) *EnvironmentConnectorRelationshipQuery {
	query := (&ConnectorClient{config: ecrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecrq.withConnector = query
	return ecrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EnvironmentConnectorRelationship.Query().
//		GroupBy(environmentconnectorrelationship.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (ecrq *EnvironmentConnectorRelationshipQuery) GroupBy(field string, fields ...string) *EnvironmentConnectorRelationshipGroupBy {
	ecrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnvironmentConnectorRelationshipGroupBy{build: ecrq}
	grbuild.flds = &ecrq.ctx.Fields
	grbuild.label = environmentconnectorrelationship.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.EnvironmentConnectorRelationship.Query().
//		Select(environmentconnectorrelationship.FieldCreateTime).
//		Scan(ctx, &v)
func (ecrq *EnvironmentConnectorRelationshipQuery) Select(fields ...string) *EnvironmentConnectorRelationshipSelect {
	ecrq.ctx.Fields = append(ecrq.ctx.Fields, fields...)
	sbuild := &EnvironmentConnectorRelationshipSelect{EnvironmentConnectorRelationshipQuery: ecrq}
	sbuild.label = environmentconnectorrelationship.Label
	sbuild.flds, sbuild.scan = &ecrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnvironmentConnectorRelationshipSelect configured with the given aggregations.
func (ecrq *EnvironmentConnectorRelationshipQuery) Aggregate(fns ...AggregateFunc) *EnvironmentConnectorRelationshipSelect {
	return ecrq.Select().Aggregate(fns...)
}

func (ecrq *EnvironmentConnectorRelationshipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecrq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecrq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecrq.ctx.Fields {
		if !environmentconnectorrelationship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if ecrq.path != nil {
		prev, err := ecrq.path(ctx)
		if err != nil {
			return err
		}
		ecrq.sql = prev
	}
	return nil
}

func (ecrq *EnvironmentConnectorRelationshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EnvironmentConnectorRelationship, error) {
	var (
		nodes       = []*EnvironmentConnectorRelationship{}
		_spec       = ecrq.querySpec()
		loadedTypes = [2]bool{
			ecrq.withEnvironment != nil,
			ecrq.withConnector != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EnvironmentConnectorRelationship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EnvironmentConnectorRelationship{config: ecrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ecrq.schemaConfig.EnvironmentConnectorRelationship
	ctx = internal.NewSchemaConfigContext(ctx, ecrq.schemaConfig)
	if len(ecrq.modifiers) > 0 {
		_spec.Modifiers = ecrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecrq.withEnvironment; query != nil {
		if err := ecrq.loadEnvironment(ctx, query, nodes, nil,
			func(n *EnvironmentConnectorRelationship, e *Environment) { n.Edges.Environment = e }); err != nil {
			return nil, err
		}
	}
	if query := ecrq.withConnector; query != nil {
		if err := ecrq.loadConnector(ctx, query, nodes, nil,
			func(n *EnvironmentConnectorRelationship, e *Connector) { n.Edges.Connector = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecrq *EnvironmentConnectorRelationshipQuery) loadEnvironment(ctx context.Context, query *EnvironmentQuery, nodes []*EnvironmentConnectorRelationship, init func(*EnvironmentConnectorRelationship), assign func(*EnvironmentConnectorRelationship, *Environment)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*EnvironmentConnectorRelationship)
	for i := range nodes {
		fk := nodes[i].EnvironmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(environment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "environment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecrq *EnvironmentConnectorRelationshipQuery) loadConnector(ctx context.Context, query *ConnectorQuery, nodes []*EnvironmentConnectorRelationship, init func(*EnvironmentConnectorRelationship), assign func(*EnvironmentConnectorRelationship, *Connector)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*EnvironmentConnectorRelationship)
	for i := range nodes {
		fk := nodes[i].ConnectorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(connector.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "connector_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ecrq *EnvironmentConnectorRelationshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecrq.querySpec()
	_spec.Node.Schema = ecrq.schemaConfig.EnvironmentConnectorRelationship
	ctx = internal.NewSchemaConfigContext(ctx, ecrq.schemaConfig)
	if len(ecrq.modifiers) > 0 {
		_spec.Modifiers = ecrq.modifiers
	}
	_spec.Node.Columns = ecrq.ctx.Fields
	if len(ecrq.ctx.Fields) > 0 {
		_spec.Unique = ecrq.ctx.Unique != nil && *ecrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecrq.driver, _spec)
}

func (ecrq *EnvironmentConnectorRelationshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(environmentconnectorrelationship.Table, environmentconnectorrelationship.Columns, sqlgraph.NewFieldSpec(environmentconnectorrelationship.FieldID, field.TypeString))
	_spec.From = ecrq.sql
	if unique := ecrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecrq.path != nil {
		_spec.Unique = true
	}
	if fields := ecrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, environmentconnectorrelationship.FieldID)
		for i := range fields {
			if fields[i] != environmentconnectorrelationship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ecrq.withEnvironment != nil {
			_spec.Node.AddColumnOnce(environmentconnectorrelationship.FieldEnvironmentID)
		}
		if ecrq.withConnector != nil {
			_spec.Node.AddColumnOnce(environmentconnectorrelationship.FieldConnectorID)
		}
	}
	if ps := ecrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecrq *EnvironmentConnectorRelationshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecrq.driver.Dialect())
	t1 := builder.Table(environmentconnectorrelationship.Table)
	columns := ecrq.ctx.Fields
	if len(columns) == 0 {
		columns = environmentconnectorrelationship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecrq.sql != nil {
		selector = ecrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecrq.ctx.Unique != nil && *ecrq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(ecrq.schemaConfig.EnvironmentConnectorRelationship)
	ctx = internal.NewSchemaConfigContext(ctx, ecrq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range ecrq.modifiers {
		m(selector)
	}
	for _, p := range ecrq.predicates {
		p(selector)
	}
	for _, p := range ecrq.order {
		p(selector)
	}
	if offset := ecrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ecrq *EnvironmentConnectorRelationshipQuery) ForUpdate(opts ...sql.LockOption) *EnvironmentConnectorRelationshipQuery {
	if ecrq.driver.Dialect() == dialect.Postgres {
		ecrq.Unique(false)
	}
	ecrq.modifiers = append(ecrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ecrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ecrq *EnvironmentConnectorRelationshipQuery) ForShare(opts ...sql.LockOption) *EnvironmentConnectorRelationshipQuery {
	if ecrq.driver.Dialect() == dialect.Postgres {
		ecrq.Unique(false)
	}
	ecrq.modifiers = append(ecrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ecrq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecrq *EnvironmentConnectorRelationshipQuery) Modify(modifiers ...func(s *sql.Selector)) *EnvironmentConnectorRelationshipSelect {
	ecrq.modifiers = append(ecrq.modifiers, modifiers...)
	return ecrq.Select()
}

// WhereP appends storage-level predicates to the EnvironmentConnectorRelationshipQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (ecrq *EnvironmentConnectorRelationshipQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.EnvironmentConnectorRelationship, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.EnvironmentConnectorRelationship(ps[i]))
	}
	ecrq.predicates = append(ecrq.predicates, wps...)
}

// EnvironmentConnectorRelationshipGroupBy is the group-by builder for EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipGroupBy struct {
	selector
	build *EnvironmentConnectorRelationshipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecrgb *EnvironmentConnectorRelationshipGroupBy) Aggregate(fns ...AggregateFunc) *EnvironmentConnectorRelationshipGroupBy {
	ecrgb.fns = append(ecrgb.fns, fns...)
	return ecrgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecrgb *EnvironmentConnectorRelationshipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecrgb.build.ctx, "GroupBy")
	if err := ecrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvironmentConnectorRelationshipQuery, *EnvironmentConnectorRelationshipGroupBy](ctx, ecrgb.build, ecrgb, ecrgb.build.inters, v)
}

func (ecrgb *EnvironmentConnectorRelationshipGroupBy) sqlScan(ctx context.Context, root *EnvironmentConnectorRelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecrgb.fns))
	for _, fn := range ecrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecrgb.flds)+len(ecrgb.fns))
		for _, f := range *ecrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnvironmentConnectorRelationshipSelect is the builder for selecting fields of EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipSelect struct {
	*EnvironmentConnectorRelationshipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecrs *EnvironmentConnectorRelationshipSelect) Aggregate(fns ...AggregateFunc) *EnvironmentConnectorRelationshipSelect {
	ecrs.fns = append(ecrs.fns, fns...)
	return ecrs
}

// Scan applies the selector query and scans the result into the given value.
func (ecrs *EnvironmentConnectorRelationshipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecrs.ctx, "Select")
	if err := ecrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvironmentConnectorRelationshipQuery, *EnvironmentConnectorRelationshipSelect](ctx, ecrs.EnvironmentConnectorRelationshipQuery, ecrs, ecrs.inters, v)
}

func (ecrs *EnvironmentConnectorRelationshipSelect) sqlScan(ctx context.Context, root *EnvironmentConnectorRelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecrs.fns))
	for _, fn := range ecrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecrs *EnvironmentConnectorRelationshipSelect) Modify(modifiers ...func(s *sql.Selector)) *EnvironmentConnectorRelationshipSelect {
	ecrs.modifiers = append(ecrs.modifiers, modifiers...)
	return ecrs
}

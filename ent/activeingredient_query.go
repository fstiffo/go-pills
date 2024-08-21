// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"fstiffo/pills/ent/activeingredient"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/medicine"
	"fstiffo/pills/ent/predicate"
	"fstiffo/pills/ent/prescription"
	"fstiffo/pills/ent/stockinglog"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActiveIngredientQuery is the builder for querying ActiveIngredient entities.
type ActiveIngredientQuery struct {
	config
	ctx                 *QueryContext
	order               []activeingredient.OrderOption
	inters              []Interceptor
	predicates          []predicate.ActiveIngredient
	withMedicines       *MedicineQuery
	withPrescriptions   *PrescriptionQuery
	withStockingLogs    *StockingLogQuery
	withConsumptionLogs *ConsumptionLogQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActiveIngredientQuery builder.
func (aiq *ActiveIngredientQuery) Where(ps ...predicate.ActiveIngredient) *ActiveIngredientQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *ActiveIngredientQuery) Limit(limit int) *ActiveIngredientQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *ActiveIngredientQuery) Offset(offset int) *ActiveIngredientQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *ActiveIngredientQuery) Unique(unique bool) *ActiveIngredientQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *ActiveIngredientQuery) Order(o ...activeingredient.OrderOption) *ActiveIngredientQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// QueryMedicines chains the current query on the "medicines" edge.
func (aiq *ActiveIngredientQuery) QueryMedicines() *MedicineQuery {
	query := (&MedicineClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activeingredient.Table, activeingredient.FieldID, selector),
			sqlgraph.To(medicine.Table, medicine.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activeingredient.MedicinesTable, activeingredient.MedicinesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrescriptions chains the current query on the "prescriptions" edge.
func (aiq *ActiveIngredientQuery) QueryPrescriptions() *PrescriptionQuery {
	query := (&PrescriptionClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activeingredient.Table, activeingredient.FieldID, selector),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activeingredient.PrescriptionsTable, activeingredient.PrescriptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStockingLogs chains the current query on the "stocking_logs" edge.
func (aiq *ActiveIngredientQuery) QueryStockingLogs() *StockingLogQuery {
	query := (&StockingLogClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activeingredient.Table, activeingredient.FieldID, selector),
			sqlgraph.To(stockinglog.Table, stockinglog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activeingredient.StockingLogsTable, activeingredient.StockingLogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConsumptionLogs chains the current query on the "consumption_logs" edge.
func (aiq *ActiveIngredientQuery) QueryConsumptionLogs() *ConsumptionLogQuery {
	query := (&ConsumptionLogClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activeingredient.Table, activeingredient.FieldID, selector),
			sqlgraph.To(consumptionlog.Table, consumptionlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activeingredient.ConsumptionLogsTable, activeingredient.ConsumptionLogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ActiveIngredient entity from the query.
// Returns a *NotFoundError when no ActiveIngredient was found.
func (aiq *ActiveIngredientQuery) First(ctx context.Context) (*ActiveIngredient, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{activeingredient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) FirstX(ctx context.Context) *ActiveIngredient {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ActiveIngredient ID from the query.
// Returns a *NotFoundError when no ActiveIngredient ID was found.
func (aiq *ActiveIngredientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{activeingredient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) FirstIDX(ctx context.Context) int {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ActiveIngredient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ActiveIngredient entity is found.
// Returns a *NotFoundError when no ActiveIngredient entities are found.
func (aiq *ActiveIngredientQuery) Only(ctx context.Context) (*ActiveIngredient, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{activeingredient.Label}
	default:
		return nil, &NotSingularError{activeingredient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) OnlyX(ctx context.Context) *ActiveIngredient {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ActiveIngredient ID in the query.
// Returns a *NotSingularError when more than one ActiveIngredient ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *ActiveIngredientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{activeingredient.Label}
	default:
		err = &NotSingularError{activeingredient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) OnlyIDX(ctx context.Context) int {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ActiveIngredients.
func (aiq *ActiveIngredientQuery) All(ctx context.Context) ([]*ActiveIngredient, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryAll)
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ActiveIngredient, *ActiveIngredientQuery]()
	return withInterceptors[[]*ActiveIngredient](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) AllX(ctx context.Context) []*ActiveIngredient {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ActiveIngredient IDs.
func (aiq *ActiveIngredientQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryIDs)
	if err = aiq.Select(activeingredient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) IDsX(ctx context.Context) []int {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *ActiveIngredientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryCount)
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*ActiveIngredientQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *ActiveIngredientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryExist)
	switch _, err := aiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *ActiveIngredientQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActiveIngredientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *ActiveIngredientQuery) Clone() *ActiveIngredientQuery {
	if aiq == nil {
		return nil
	}
	return &ActiveIngredientQuery{
		config:              aiq.config,
		ctx:                 aiq.ctx.Clone(),
		order:               append([]activeingredient.OrderOption{}, aiq.order...),
		inters:              append([]Interceptor{}, aiq.inters...),
		predicates:          append([]predicate.ActiveIngredient{}, aiq.predicates...),
		withMedicines:       aiq.withMedicines.Clone(),
		withPrescriptions:   aiq.withPrescriptions.Clone(),
		withStockingLogs:    aiq.withStockingLogs.Clone(),
		withConsumptionLogs: aiq.withConsumptionLogs.Clone(),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// WithMedicines tells the query-builder to eager-load the nodes that are connected to
// the "medicines" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ActiveIngredientQuery) WithMedicines(opts ...func(*MedicineQuery)) *ActiveIngredientQuery {
	query := (&MedicineClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withMedicines = query
	return aiq
}

// WithPrescriptions tells the query-builder to eager-load the nodes that are connected to
// the "prescriptions" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ActiveIngredientQuery) WithPrescriptions(opts ...func(*PrescriptionQuery)) *ActiveIngredientQuery {
	query := (&PrescriptionClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withPrescriptions = query
	return aiq
}

// WithStockingLogs tells the query-builder to eager-load the nodes that are connected to
// the "stocking_logs" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ActiveIngredientQuery) WithStockingLogs(opts ...func(*StockingLogQuery)) *ActiveIngredientQuery {
	query := (&StockingLogClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withStockingLogs = query
	return aiq
}

// WithConsumptionLogs tells the query-builder to eager-load the nodes that are connected to
// the "consumption_logs" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ActiveIngredientQuery) WithConsumptionLogs(opts ...func(*ConsumptionLogQuery)) *ActiveIngredientQuery {
	query := (&ConsumptionLogClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withConsumptionLogs = query
	return aiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ActiveIngredient.Query().
//		GroupBy(activeingredient.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aiq *ActiveIngredientQuery) GroupBy(field string, fields ...string) *ActiveIngredientGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ActiveIngredientGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = activeingredient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ActiveIngredient.Query().
//		Select(activeingredient.FieldName).
//		Scan(ctx, &v)
func (aiq *ActiveIngredientQuery) Select(fields ...string) *ActiveIngredientSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &ActiveIngredientSelect{ActiveIngredientQuery: aiq}
	sbuild.label = activeingredient.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ActiveIngredientSelect configured with the given aggregations.
func (aiq *ActiveIngredientQuery) Aggregate(fns ...AggregateFunc) *ActiveIngredientSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *ActiveIngredientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aiq); err != nil {
				return err
			}
		}
	}
	for _, f := range aiq.ctx.Fields {
		if !activeingredient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *ActiveIngredientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ActiveIngredient, error) {
	var (
		nodes       = []*ActiveIngredient{}
		_spec       = aiq.querySpec()
		loadedTypes = [4]bool{
			aiq.withMedicines != nil,
			aiq.withPrescriptions != nil,
			aiq.withStockingLogs != nil,
			aiq.withConsumptionLogs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ActiveIngredient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ActiveIngredient{config: aiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aiq.withMedicines; query != nil {
		if err := aiq.loadMedicines(ctx, query, nodes,
			func(n *ActiveIngredient) { n.Edges.Medicines = []*Medicine{} },
			func(n *ActiveIngredient, e *Medicine) { n.Edges.Medicines = append(n.Edges.Medicines, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withPrescriptions; query != nil {
		if err := aiq.loadPrescriptions(ctx, query, nodes,
			func(n *ActiveIngredient) { n.Edges.Prescriptions = []*Prescription{} },
			func(n *ActiveIngredient, e *Prescription) { n.Edges.Prescriptions = append(n.Edges.Prescriptions, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withStockingLogs; query != nil {
		if err := aiq.loadStockingLogs(ctx, query, nodes,
			func(n *ActiveIngredient) { n.Edges.StockingLogs = []*StockingLog{} },
			func(n *ActiveIngredient, e *StockingLog) { n.Edges.StockingLogs = append(n.Edges.StockingLogs, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withConsumptionLogs; query != nil {
		if err := aiq.loadConsumptionLogs(ctx, query, nodes,
			func(n *ActiveIngredient) { n.Edges.ConsumptionLogs = []*ConsumptionLog{} },
			func(n *ActiveIngredient, e *ConsumptionLog) {
				n.Edges.ConsumptionLogs = append(n.Edges.ConsumptionLogs, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aiq *ActiveIngredientQuery) loadMedicines(ctx context.Context, query *MedicineQuery, nodes []*ActiveIngredient, init func(*ActiveIngredient), assign func(*ActiveIngredient, *Medicine)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ActiveIngredient)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Medicine(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(activeingredient.MedicinesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.active_ingredient_medicines
		if fk == nil {
			return fmt.Errorf(`foreign-key "active_ingredient_medicines" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "active_ingredient_medicines" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aiq *ActiveIngredientQuery) loadPrescriptions(ctx context.Context, query *PrescriptionQuery, nodes []*ActiveIngredient, init func(*ActiveIngredient), assign func(*ActiveIngredient, *Prescription)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ActiveIngredient)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(activeingredient.PrescriptionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.active_ingredient_prescriptions
		if fk == nil {
			return fmt.Errorf(`foreign-key "active_ingredient_prescriptions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "active_ingredient_prescriptions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aiq *ActiveIngredientQuery) loadStockingLogs(ctx context.Context, query *StockingLogQuery, nodes []*ActiveIngredient, init func(*ActiveIngredient), assign func(*ActiveIngredient, *StockingLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ActiveIngredient)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.StockingLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(activeingredient.StockingLogsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.active_ingredient_stocking_logs
		if fk == nil {
			return fmt.Errorf(`foreign-key "active_ingredient_stocking_logs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "active_ingredient_stocking_logs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aiq *ActiveIngredientQuery) loadConsumptionLogs(ctx context.Context, query *ConsumptionLogQuery, nodes []*ActiveIngredient, init func(*ActiveIngredient), assign func(*ActiveIngredient, *ConsumptionLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ActiveIngredient)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ConsumptionLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(activeingredient.ConsumptionLogsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.active_ingredient_consumption_logs
		if fk == nil {
			return fmt.Errorf(`foreign-key "active_ingredient_consumption_logs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "active_ingredient_consumption_logs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aiq *ActiveIngredientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *ActiveIngredientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(activeingredient.Table, activeingredient.Columns, sqlgraph.NewFieldSpec(activeingredient.FieldID, field.TypeInt))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activeingredient.FieldID)
		for i := range fields {
			if fields[i] != activeingredient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aiq *ActiveIngredientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(activeingredient.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = activeingredient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aiq.ctx.Unique != nil && *aiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector)
	}
	if offset := aiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ActiveIngredientGroupBy is the group-by builder for ActiveIngredient entities.
type ActiveIngredientGroupBy struct {
	selector
	build *ActiveIngredientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *ActiveIngredientGroupBy) Aggregate(fns ...AggregateFunc) *ActiveIngredientGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *ActiveIngredientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, ent.OpQueryGroupBy)
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveIngredientQuery, *ActiveIngredientGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *ActiveIngredientGroupBy) sqlScan(ctx context.Context, root *ActiveIngredientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aigb.fns))
	for _, fn := range aigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aigb.flds)+len(aigb.fns))
		for _, f := range *aigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ActiveIngredientSelect is the builder for selecting fields of ActiveIngredient entities.
type ActiveIngredientSelect struct {
	*ActiveIngredientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *ActiveIngredientSelect) Aggregate(fns ...AggregateFunc) *ActiveIngredientSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *ActiveIngredientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, ent.OpQuerySelect)
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveIngredientQuery, *ActiveIngredientSelect](ctx, ais.ActiveIngredientQuery, ais, ais.inters, v)
}

func (ais *ActiveIngredientSelect) sqlScan(ctx context.Context, root *ActiveIngredientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ais.fns))
	for _, fn := range ais.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ais.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/predicate"
	"fstiffo/pills/ent/prescription"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConsumptionLogUpdate is the builder for updating ConsumptionLog entities.
type ConsumptionLogUpdate struct {
	config
	hooks    []Hook
	mutation *ConsumptionLogMutation
}

// Where appends a list predicates to the ConsumptionLogUpdate builder.
func (clu *ConsumptionLogUpdate) Where(ps ...predicate.ConsumptionLog) *ConsumptionLogUpdate {
	clu.mutation.Where(ps...)
	return clu
}

// SetConsumedAt sets the "consumed_at" field.
func (clu *ConsumptionLogUpdate) SetConsumedAt(t time.Time) *ConsumptionLogUpdate {
	clu.mutation.SetConsumedAt(t)
	return clu
}

// SetNillableConsumedAt sets the "consumed_at" field if the given value is not nil.
func (clu *ConsumptionLogUpdate) SetNillableConsumedAt(t *time.Time) *ConsumptionLogUpdate {
	if t != nil {
		clu.SetConsumedAt(*t)
	}
	return clu
}

// SetPrescriptionID sets the "prescription" edge to the Prescription entity by ID.
func (clu *ConsumptionLogUpdate) SetPrescriptionID(id int) *ConsumptionLogUpdate {
	clu.mutation.SetPrescriptionID(id)
	return clu
}

// SetNillablePrescriptionID sets the "prescription" edge to the Prescription entity by ID if the given value is not nil.
func (clu *ConsumptionLogUpdate) SetNillablePrescriptionID(id *int) *ConsumptionLogUpdate {
	if id != nil {
		clu = clu.SetPrescriptionID(*id)
	}
	return clu
}

// SetPrescription sets the "prescription" edge to the Prescription entity.
func (clu *ConsumptionLogUpdate) SetPrescription(p *Prescription) *ConsumptionLogUpdate {
	return clu.SetPrescriptionID(p.ID)
}

// Mutation returns the ConsumptionLogMutation object of the builder.
func (clu *ConsumptionLogUpdate) Mutation() *ConsumptionLogMutation {
	return clu.mutation
}

// ClearPrescription clears the "prescription" edge to the Prescription entity.
func (clu *ConsumptionLogUpdate) ClearPrescription() *ConsumptionLogUpdate {
	clu.mutation.ClearPrescription()
	return clu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (clu *ConsumptionLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, clu.sqlSave, clu.mutation, clu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (clu *ConsumptionLogUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *ConsumptionLogUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *ConsumptionLogUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (clu *ConsumptionLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(consumptionlog.Table, consumptionlog.Columns, sqlgraph.NewFieldSpec(consumptionlog.FieldID, field.TypeInt))
	if ps := clu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.ConsumedAt(); ok {
		_spec.SetField(consumptionlog.FieldConsumedAt, field.TypeTime, value)
	}
	if clu.mutation.PrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   consumptionlog.PrescriptionTable,
			Columns: []string{consumptionlog.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.PrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   consumptionlog.PrescriptionTable,
			Columns: []string{consumptionlog.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{consumptionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	clu.mutation.done = true
	return n, nil
}

// ConsumptionLogUpdateOne is the builder for updating a single ConsumptionLog entity.
type ConsumptionLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConsumptionLogMutation
}

// SetConsumedAt sets the "consumed_at" field.
func (cluo *ConsumptionLogUpdateOne) SetConsumedAt(t time.Time) *ConsumptionLogUpdateOne {
	cluo.mutation.SetConsumedAt(t)
	return cluo
}

// SetNillableConsumedAt sets the "consumed_at" field if the given value is not nil.
func (cluo *ConsumptionLogUpdateOne) SetNillableConsumedAt(t *time.Time) *ConsumptionLogUpdateOne {
	if t != nil {
		cluo.SetConsumedAt(*t)
	}
	return cluo
}

// SetPrescriptionID sets the "prescription" edge to the Prescription entity by ID.
func (cluo *ConsumptionLogUpdateOne) SetPrescriptionID(id int) *ConsumptionLogUpdateOne {
	cluo.mutation.SetPrescriptionID(id)
	return cluo
}

// SetNillablePrescriptionID sets the "prescription" edge to the Prescription entity by ID if the given value is not nil.
func (cluo *ConsumptionLogUpdateOne) SetNillablePrescriptionID(id *int) *ConsumptionLogUpdateOne {
	if id != nil {
		cluo = cluo.SetPrescriptionID(*id)
	}
	return cluo
}

// SetPrescription sets the "prescription" edge to the Prescription entity.
func (cluo *ConsumptionLogUpdateOne) SetPrescription(p *Prescription) *ConsumptionLogUpdateOne {
	return cluo.SetPrescriptionID(p.ID)
}

// Mutation returns the ConsumptionLogMutation object of the builder.
func (cluo *ConsumptionLogUpdateOne) Mutation() *ConsumptionLogMutation {
	return cluo.mutation
}

// ClearPrescription clears the "prescription" edge to the Prescription entity.
func (cluo *ConsumptionLogUpdateOne) ClearPrescription() *ConsumptionLogUpdateOne {
	cluo.mutation.ClearPrescription()
	return cluo
}

// Where appends a list predicates to the ConsumptionLogUpdate builder.
func (cluo *ConsumptionLogUpdateOne) Where(ps ...predicate.ConsumptionLog) *ConsumptionLogUpdateOne {
	cluo.mutation.Where(ps...)
	return cluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cluo *ConsumptionLogUpdateOne) Select(field string, fields ...string) *ConsumptionLogUpdateOne {
	cluo.fields = append([]string{field}, fields...)
	return cluo
}

// Save executes the query and returns the updated ConsumptionLog entity.
func (cluo *ConsumptionLogUpdateOne) Save(ctx context.Context) (*ConsumptionLog, error) {
	return withHooks(ctx, cluo.sqlSave, cluo.mutation, cluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *ConsumptionLogUpdateOne) SaveX(ctx context.Context) *ConsumptionLog {
	node, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cluo *ConsumptionLogUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *ConsumptionLogUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cluo *ConsumptionLogUpdateOne) sqlSave(ctx context.Context) (_node *ConsumptionLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(consumptionlog.Table, consumptionlog.Columns, sqlgraph.NewFieldSpec(consumptionlog.FieldID, field.TypeInt))
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ConsumptionLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, consumptionlog.FieldID)
		for _, f := range fields {
			if !consumptionlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != consumptionlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cluo.mutation.ConsumedAt(); ok {
		_spec.SetField(consumptionlog.FieldConsumedAt, field.TypeTime, value)
	}
	if cluo.mutation.PrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   consumptionlog.PrescriptionTable,
			Columns: []string{consumptionlog.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.PrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   consumptionlog.PrescriptionTable,
			Columns: []string{consumptionlog.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ConsumptionLog{config: cluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{consumptionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cluo.mutation.done = true
	return _node, nil
}

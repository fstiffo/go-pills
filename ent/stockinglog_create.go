// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fstiffo/pills/ent/medicine"
	"fstiffo/pills/ent/stockinglog"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockingLogCreate is the builder for creating a StockingLog entity.
type StockingLogCreate struct {
	config
	mutation *StockingLogMutation
	hooks    []Hook
}

// SetStockedAt sets the "stocked_at" field.
func (slc *StockingLogCreate) SetStockedAt(t time.Time) *StockingLogCreate {
	slc.mutation.SetStockedAt(t)
	return slc
}

// SetQuantity sets the "quantity" field.
func (slc *StockingLogCreate) SetQuantity(i int) *StockingLogCreate {
	slc.mutation.SetQuantity(i)
	return slc
}

// SetMedicineID sets the "medicine" edge to the Medicine entity by ID.
func (slc *StockingLogCreate) SetMedicineID(id int) *StockingLogCreate {
	slc.mutation.SetMedicineID(id)
	return slc
}

// SetNillableMedicineID sets the "medicine" edge to the Medicine entity by ID if the given value is not nil.
func (slc *StockingLogCreate) SetNillableMedicineID(id *int) *StockingLogCreate {
	if id != nil {
		slc = slc.SetMedicineID(*id)
	}
	return slc
}

// SetMedicine sets the "medicine" edge to the Medicine entity.
func (slc *StockingLogCreate) SetMedicine(m *Medicine) *StockingLogCreate {
	return slc.SetMedicineID(m.ID)
}

// Mutation returns the StockingLogMutation object of the builder.
func (slc *StockingLogCreate) Mutation() *StockingLogMutation {
	return slc.mutation
}

// Save creates the StockingLog in the database.
func (slc *StockingLogCreate) Save(ctx context.Context) (*StockingLog, error) {
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *StockingLogCreate) SaveX(ctx context.Context) *StockingLog {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *StockingLogCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *StockingLogCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *StockingLogCreate) check() error {
	if _, ok := slc.mutation.StockedAt(); !ok {
		return &ValidationError{Name: "stocked_at", err: errors.New(`ent: missing required field "StockingLog.stocked_at"`)}
	}
	if _, ok := slc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "StockingLog.quantity"`)}
	}
	if v, ok := slc.mutation.Quantity(); ok {
		if err := stockinglog.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "StockingLog.quantity": %w`, err)}
		}
	}
	return nil
}

func (slc *StockingLogCreate) sqlSave(ctx context.Context) (*StockingLog, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *StockingLogCreate) createSpec() (*StockingLog, *sqlgraph.CreateSpec) {
	var (
		_node = &StockingLog{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(stockinglog.Table, sqlgraph.NewFieldSpec(stockinglog.FieldID, field.TypeInt))
	)
	if value, ok := slc.mutation.StockedAt(); ok {
		_spec.SetField(stockinglog.FieldStockedAt, field.TypeTime, value)
		_node.StockedAt = value
	}
	if value, ok := slc.mutation.Quantity(); ok {
		_spec.SetField(stockinglog.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if nodes := slc.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stockinglog.MedicineTable,
			Columns: []string{stockinglog.MedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicine.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.medicine_stocking_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StockingLogCreateBulk is the builder for creating many StockingLog entities in bulk.
type StockingLogCreateBulk struct {
	config
	err      error
	builders []*StockingLogCreate
}

// Save creates the StockingLog entities in the database.
func (slcb *StockingLogCreateBulk) Save(ctx context.Context) ([]*StockingLog, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*StockingLog, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockingLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *StockingLogCreateBulk) SaveX(ctx context.Context) []*StockingLog {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *StockingLogCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *StockingLogCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}

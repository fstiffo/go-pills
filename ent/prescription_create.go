// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fstiffo/pills/ent/activeingredient"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/prescription"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrescriptionCreate is the builder for creating a Prescription entity.
type PrescriptionCreate struct {
	config
	mutation *PrescriptionMutation
	hooks    []Hook
}

// SetDosage sets the "dosage" field.
func (pc *PrescriptionCreate) SetDosage(i int) *PrescriptionCreate {
	pc.mutation.SetDosage(i)
	return pc
}

// SetDosageFrequency sets the "dosage_frequency" field.
func (pc *PrescriptionCreate) SetDosageFrequency(i int) *PrescriptionCreate {
	pc.mutation.SetDosageFrequency(i)
	return pc
}

// SetNillableDosageFrequency sets the "dosage_frequency" field if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableDosageFrequency(i *int) *PrescriptionCreate {
	if i != nil {
		pc.SetDosageFrequency(*i)
	}
	return pc
}

// SetStartDate sets the "start_date" field.
func (pc *PrescriptionCreate) SetStartDate(t time.Time) *PrescriptionCreate {
	pc.mutation.SetStartDate(t)
	return pc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableStartDate(t *time.Time) *PrescriptionCreate {
	if t != nil {
		pc.SetStartDate(*t)
	}
	return pc
}

// SetEndDate sets the "end_date" field.
func (pc *PrescriptionCreate) SetEndDate(t time.Time) *PrescriptionCreate {
	pc.mutation.SetEndDate(t)
	return pc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableEndDate(t *time.Time) *PrescriptionCreate {
	if t != nil {
		pc.SetEndDate(*t)
	}
	return pc
}

// AddComsumptionLogIDs adds the "comsumption_logs" edge to the ConsumptionLog entity by IDs.
func (pc *PrescriptionCreate) AddComsumptionLogIDs(ids ...int) *PrescriptionCreate {
	pc.mutation.AddComsumptionLogIDs(ids...)
	return pc
}

// AddComsumptionLogs adds the "comsumption_logs" edges to the ConsumptionLog entity.
func (pc *PrescriptionCreate) AddComsumptionLogs(c ...*ConsumptionLog) *PrescriptionCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddComsumptionLogIDs(ids...)
}

// SetActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID.
func (pc *PrescriptionCreate) SetActiveIngredientID(id int) *PrescriptionCreate {
	pc.mutation.SetActiveIngredientID(id)
	return pc
}

// SetNillableActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableActiveIngredientID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetActiveIngredientID(*id)
	}
	return pc
}

// SetActiveIngredient sets the "active_ingredient" edge to the ActiveIngredient entity.
func (pc *PrescriptionCreate) SetActiveIngredient(a *ActiveIngredient) *PrescriptionCreate {
	return pc.SetActiveIngredientID(a.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pc *PrescriptionCreate) Mutation() *PrescriptionMutation {
	return pc.mutation
}

// Save creates the Prescription in the database.
func (pc *PrescriptionCreate) Save(ctx context.Context) (*Prescription, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrescriptionCreate) SaveX(ctx context.Context) *Prescription {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PrescriptionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PrescriptionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PrescriptionCreate) defaults() {
	if _, ok := pc.mutation.DosageFrequency(); !ok {
		v := prescription.DefaultDosageFrequency
		pc.mutation.SetDosageFrequency(v)
	}
	if _, ok := pc.mutation.StartDate(); !ok {
		v := prescription.DefaultStartDate()
		pc.mutation.SetStartDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PrescriptionCreate) check() error {
	if _, ok := pc.mutation.Dosage(); !ok {
		return &ValidationError{Name: "dosage", err: errors.New(`ent: missing required field "Prescription.dosage"`)}
	}
	if v, ok := pc.mutation.Dosage(); ok {
		if err := prescription.DosageValidator(v); err != nil {
			return &ValidationError{Name: "dosage", err: fmt.Errorf(`ent: validator failed for field "Prescription.dosage": %w`, err)}
		}
	}
	if v, ok := pc.mutation.DosageFrequency(); ok {
		if err := prescription.DosageFrequencyValidator(v); err != nil {
			return &ValidationError{Name: "dosage_frequency", err: fmt.Errorf(`ent: validator failed for field "Prescription.dosage_frequency": %w`, err)}
		}
	}
	return nil
}

func (pc *PrescriptionCreate) sqlSave(ctx context.Context) (*Prescription, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PrescriptionCreate) createSpec() (*Prescription, *sqlgraph.CreateSpec) {
	var (
		_node = &Prescription{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(prescription.Table, sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Dosage(); ok {
		_spec.SetField(prescription.FieldDosage, field.TypeInt, value)
		_node.Dosage = value
	}
	if value, ok := pc.mutation.DosageFrequency(); ok {
		_spec.SetField(prescription.FieldDosageFrequency, field.TypeInt, value)
		_node.DosageFrequency = value
	}
	if value, ok := pc.mutation.StartDate(); ok {
		_spec.SetField(prescription.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := pc.mutation.EndDate(); ok {
		_spec.SetField(prescription.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if nodes := pc.mutation.ComsumptionLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prescription.ComsumptionLogsTable,
			Columns: []string{prescription.ComsumptionLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(consumptionlog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ActiveIngredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.ActiveIngredientTable,
			Columns: []string{prescription.ActiveIngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.active_ingredient_prescriptions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PrescriptionCreateBulk is the builder for creating many Prescription entities in bulk.
type PrescriptionCreateBulk struct {
	config
	err      error
	builders []*PrescriptionCreate
}

// Save creates the Prescription entities in the database.
func (pcb *PrescriptionCreateBulk) Save(ctx context.Context) ([]*Prescription, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Prescription, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PrescriptionCreateBulk) SaveX(ctx context.Context) []*Prescription {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PrescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PrescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
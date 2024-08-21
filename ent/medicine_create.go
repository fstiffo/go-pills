// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fstiffo/pills/ent/activeingredient"
	"fstiffo/pills/ent/medicine"
	"fstiffo/pills/ent/stockinglog"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MedicineCreate is the builder for creating a Medicine entity.
type MedicineCreate struct {
	config
	mutation *MedicineMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MedicineCreate) SetName(s string) *MedicineCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetMah sets the "mah" field.
func (mc *MedicineCreate) SetMah(s string) *MedicineCreate {
	mc.mutation.SetMah(s)
	return mc
}

// SetDosage sets the "dosage" field.
func (mc *MedicineCreate) SetDosage(f float64) *MedicineCreate {
	mc.mutation.SetDosage(f)
	return mc
}

// SetAtc sets the "atc" field.
func (mc *MedicineCreate) SetAtc(s string) *MedicineCreate {
	mc.mutation.SetAtc(s)
	return mc
}

// SetPackage sets the "package" field.
func (mc *MedicineCreate) SetPackage(s string) *MedicineCreate {
	mc.mutation.SetPackage(s)
	return mc
}

// SetForm sets the "form" field.
func (mc *MedicineCreate) SetForm(s string) *MedicineCreate {
	mc.mutation.SetForm(s)
	return mc
}

// SetBoxSize sets the "box_size" field.
func (mc *MedicineCreate) SetBoxSize(i int) *MedicineCreate {
	mc.mutation.SetBoxSize(i)
	return mc
}

// AddStockingLogIDs adds the "stocking_logs" edge to the StockingLog entity by IDs.
func (mc *MedicineCreate) AddStockingLogIDs(ids ...int) *MedicineCreate {
	mc.mutation.AddStockingLogIDs(ids...)
	return mc
}

// AddStockingLogs adds the "stocking_logs" edges to the StockingLog entity.
func (mc *MedicineCreate) AddStockingLogs(s ...*StockingLog) *MedicineCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddStockingLogIDs(ids...)
}

// SetActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID.
func (mc *MedicineCreate) SetActiveIngredientID(id int) *MedicineCreate {
	mc.mutation.SetActiveIngredientID(id)
	return mc
}

// SetNillableActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID if the given value is not nil.
func (mc *MedicineCreate) SetNillableActiveIngredientID(id *int) *MedicineCreate {
	if id != nil {
		mc = mc.SetActiveIngredientID(*id)
	}
	return mc
}

// SetActiveIngredient sets the "active_ingredient" edge to the ActiveIngredient entity.
func (mc *MedicineCreate) SetActiveIngredient(a *ActiveIngredient) *MedicineCreate {
	return mc.SetActiveIngredientID(a.ID)
}

// Mutation returns the MedicineMutation object of the builder.
func (mc *MedicineCreate) Mutation() *MedicineMutation {
	return mc.mutation
}

// Save creates the Medicine in the database.
func (mc *MedicineCreate) Save(ctx context.Context) (*Medicine, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicineCreate) SaveX(ctx context.Context) *Medicine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MedicineCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MedicineCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MedicineCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Medicine.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Medicine.name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Mah(); !ok {
		return &ValidationError{Name: "mah", err: errors.New(`ent: missing required field "Medicine.mah"`)}
	}
	if v, ok := mc.mutation.Mah(); ok {
		if err := medicine.MahValidator(v); err != nil {
			return &ValidationError{Name: "mah", err: fmt.Errorf(`ent: validator failed for field "Medicine.mah": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Dosage(); !ok {
		return &ValidationError{Name: "dosage", err: errors.New(`ent: missing required field "Medicine.dosage"`)}
	}
	if v, ok := mc.mutation.Dosage(); ok {
		if err := medicine.DosageValidator(v); err != nil {
			return &ValidationError{Name: "dosage", err: fmt.Errorf(`ent: validator failed for field "Medicine.dosage": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Atc(); !ok {
		return &ValidationError{Name: "atc", err: errors.New(`ent: missing required field "Medicine.atc"`)}
	}
	if v, ok := mc.mutation.Atc(); ok {
		if err := medicine.AtcValidator(v); err != nil {
			return &ValidationError{Name: "atc", err: fmt.Errorf(`ent: validator failed for field "Medicine.atc": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Package(); !ok {
		return &ValidationError{Name: "package", err: errors.New(`ent: missing required field "Medicine.package"`)}
	}
	if _, ok := mc.mutation.Form(); !ok {
		return &ValidationError{Name: "form", err: errors.New(`ent: missing required field "Medicine.form"`)}
	}
	if _, ok := mc.mutation.BoxSize(); !ok {
		return &ValidationError{Name: "box_size", err: errors.New(`ent: missing required field "Medicine.box_size"`)}
	}
	if v, ok := mc.mutation.BoxSize(); ok {
		if err := medicine.BoxSizeValidator(v); err != nil {
			return &ValidationError{Name: "box_size", err: fmt.Errorf(`ent: validator failed for field "Medicine.box_size": %w`, err)}
		}
	}
	return nil
}

func (mc *MedicineCreate) sqlSave(ctx context.Context) (*Medicine, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MedicineCreate) createSpec() (*Medicine, *sqlgraph.CreateSpec) {
	var (
		_node = &Medicine{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(medicine.Table, sqlgraph.NewFieldSpec(medicine.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(medicine.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Mah(); ok {
		_spec.SetField(medicine.FieldMah, field.TypeString, value)
		_node.Mah = value
	}
	if value, ok := mc.mutation.Dosage(); ok {
		_spec.SetField(medicine.FieldDosage, field.TypeFloat64, value)
		_node.Dosage = value
	}
	if value, ok := mc.mutation.Atc(); ok {
		_spec.SetField(medicine.FieldAtc, field.TypeString, value)
		_node.Atc = value
	}
	if value, ok := mc.mutation.Package(); ok {
		_spec.SetField(medicine.FieldPackage, field.TypeString, value)
		_node.Package = value
	}
	if value, ok := mc.mutation.Form(); ok {
		_spec.SetField(medicine.FieldForm, field.TypeString, value)
		_node.Form = value
	}
	if value, ok := mc.mutation.BoxSize(); ok {
		_spec.SetField(medicine.FieldBoxSize, field.TypeInt, value)
		_node.BoxSize = value
	}
	if nodes := mc.mutation.StockingLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.StockingLogsTable,
			Columns: []string{medicine.StockingLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockinglog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ActiveIngredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicine.ActiveIngredientTable,
			Columns: []string{medicine.ActiveIngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.active_ingredient_medicines = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MedicineCreateBulk is the builder for creating many Medicine entities in bulk.
type MedicineCreateBulk struct {
	config
	err      error
	builders []*MedicineCreate
}

// Save creates the Medicine entities in the database.
func (mcb *MedicineCreateBulk) Save(ctx context.Context) ([]*Medicine, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Medicine, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MedicineMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MedicineCreateBulk) SaveX(ctx context.Context) []*Medicine {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MedicineCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MedicineCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

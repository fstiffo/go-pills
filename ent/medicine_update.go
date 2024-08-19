// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fstiffo/pills/ent/activeingredient"
	"fstiffo/pills/ent/medicine"
	"fstiffo/pills/ent/predicate"
	"fstiffo/pills/ent/purchase"
	"fstiffo/pills/ent/stockinglog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MedicineUpdate is the builder for updating Medicine entities.
type MedicineUpdate struct {
	config
	hooks    []Hook
	mutation *MedicineMutation
}

// Where appends a list predicates to the MedicineUpdate builder.
func (mu *MedicineUpdate) Where(ps ...predicate.Medicine) *MedicineUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MedicineUpdate) SetName(s string) *MedicineUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableName(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetMah sets the "mah" field.
func (mu *MedicineUpdate) SetMah(s string) *MedicineUpdate {
	mu.mutation.SetMah(s)
	return mu
}

// SetNillableMah sets the "mah" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableMah(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetMah(*s)
	}
	return mu
}

// SetDosage sets the "dosage" field.
func (mu *MedicineUpdate) SetDosage(f float64) *MedicineUpdate {
	mu.mutation.ResetDosage()
	mu.mutation.SetDosage(f)
	return mu
}

// SetNillableDosage sets the "dosage" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableDosage(f *float64) *MedicineUpdate {
	if f != nil {
		mu.SetDosage(*f)
	}
	return mu
}

// AddDosage adds f to the "dosage" field.
func (mu *MedicineUpdate) AddDosage(f float64) *MedicineUpdate {
	mu.mutation.AddDosage(f)
	return mu
}

// SetUnit sets the "unit" field.
func (mu *MedicineUpdate) SetUnit(s string) *MedicineUpdate {
	mu.mutation.SetUnit(s)
	return mu
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableUnit(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetUnit(*s)
	}
	return mu
}

// SetAtc sets the "atc" field.
func (mu *MedicineUpdate) SetAtc(s string) *MedicineUpdate {
	mu.mutation.SetAtc(s)
	return mu
}

// SetNillableAtc sets the "atc" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableAtc(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetAtc(*s)
	}
	return mu
}

// SetPackage sets the "package" field.
func (mu *MedicineUpdate) SetPackage(s string) *MedicineUpdate {
	mu.mutation.SetPackage(s)
	return mu
}

// SetNillablePackage sets the "package" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillablePackage(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetPackage(*s)
	}
	return mu
}

// SetForm sets the "form" field.
func (mu *MedicineUpdate) SetForm(s string) *MedicineUpdate {
	mu.mutation.SetForm(s)
	return mu
}

// SetNillableForm sets the "form" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableForm(s *string) *MedicineUpdate {
	if s != nil {
		mu.SetForm(*s)
	}
	return mu
}

// SetBoxSize sets the "box_size" field.
func (mu *MedicineUpdate) SetBoxSize(i int) *MedicineUpdate {
	mu.mutation.ResetBoxSize()
	mu.mutation.SetBoxSize(i)
	return mu
}

// SetNillableBoxSize sets the "box_size" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableBoxSize(i *int) *MedicineUpdate {
	if i != nil {
		mu.SetBoxSize(*i)
	}
	return mu
}

// AddBoxSize adds i to the "box_size" field.
func (mu *MedicineUpdate) AddBoxSize(i int) *MedicineUpdate {
	mu.mutation.AddBoxSize(i)
	return mu
}

// SetStock sets the "stock" field.
func (mu *MedicineUpdate) SetStock(f float32) *MedicineUpdate {
	mu.mutation.ResetStock()
	mu.mutation.SetStock(f)
	return mu
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableStock(f *float32) *MedicineUpdate {
	if f != nil {
		mu.SetStock(*f)
	}
	return mu
}

// AddStock adds f to the "stock" field.
func (mu *MedicineUpdate) AddStock(f float32) *MedicineUpdate {
	mu.mutation.AddStock(f)
	return mu
}

// ClearStock clears the value of the "stock" field.
func (mu *MedicineUpdate) ClearStock() *MedicineUpdate {
	mu.mutation.ClearStock()
	return mu
}

// SetLastStockUpdate sets the "last_stock_update" field.
func (mu *MedicineUpdate) SetLastStockUpdate(t time.Time) *MedicineUpdate {
	mu.mutation.SetLastStockUpdate(t)
	return mu
}

// SetNillableLastStockUpdate sets the "last_stock_update" field if the given value is not nil.
func (mu *MedicineUpdate) SetNillableLastStockUpdate(t *time.Time) *MedicineUpdate {
	if t != nil {
		mu.SetLastStockUpdate(*t)
	}
	return mu
}

// AddPurchaseIDs adds the "purchases" edge to the Purchase entity by IDs.
func (mu *MedicineUpdate) AddPurchaseIDs(ids ...int) *MedicineUpdate {
	mu.mutation.AddPurchaseIDs(ids...)
	return mu
}

// AddPurchases adds the "purchases" edges to the Purchase entity.
func (mu *MedicineUpdate) AddPurchases(p ...*Purchase) *MedicineUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPurchaseIDs(ids...)
}

// AddStockingLogIDs adds the "stocking_logs" edge to the StockingLog entity by IDs.
func (mu *MedicineUpdate) AddStockingLogIDs(ids ...int) *MedicineUpdate {
	mu.mutation.AddStockingLogIDs(ids...)
	return mu
}

// AddStockingLogs adds the "stocking_logs" edges to the StockingLog entity.
func (mu *MedicineUpdate) AddStockingLogs(s ...*StockingLog) *MedicineUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddStockingLogIDs(ids...)
}

// SetActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID.
func (mu *MedicineUpdate) SetActiveIngredientID(id int) *MedicineUpdate {
	mu.mutation.SetActiveIngredientID(id)
	return mu
}

// SetNillableActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID if the given value is not nil.
func (mu *MedicineUpdate) SetNillableActiveIngredientID(id *int) *MedicineUpdate {
	if id != nil {
		mu = mu.SetActiveIngredientID(*id)
	}
	return mu
}

// SetActiveIngredient sets the "active_ingredient" edge to the ActiveIngredient entity.
func (mu *MedicineUpdate) SetActiveIngredient(a *ActiveIngredient) *MedicineUpdate {
	return mu.SetActiveIngredientID(a.ID)
}

// Mutation returns the MedicineMutation object of the builder.
func (mu *MedicineUpdate) Mutation() *MedicineMutation {
	return mu.mutation
}

// ClearPurchases clears all "purchases" edges to the Purchase entity.
func (mu *MedicineUpdate) ClearPurchases() *MedicineUpdate {
	mu.mutation.ClearPurchases()
	return mu
}

// RemovePurchaseIDs removes the "purchases" edge to Purchase entities by IDs.
func (mu *MedicineUpdate) RemovePurchaseIDs(ids ...int) *MedicineUpdate {
	mu.mutation.RemovePurchaseIDs(ids...)
	return mu
}

// RemovePurchases removes "purchases" edges to Purchase entities.
func (mu *MedicineUpdate) RemovePurchases(p ...*Purchase) *MedicineUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePurchaseIDs(ids...)
}

// ClearStockingLogs clears all "stocking_logs" edges to the StockingLog entity.
func (mu *MedicineUpdate) ClearStockingLogs() *MedicineUpdate {
	mu.mutation.ClearStockingLogs()
	return mu
}

// RemoveStockingLogIDs removes the "stocking_logs" edge to StockingLog entities by IDs.
func (mu *MedicineUpdate) RemoveStockingLogIDs(ids ...int) *MedicineUpdate {
	mu.mutation.RemoveStockingLogIDs(ids...)
	return mu
}

// RemoveStockingLogs removes "stocking_logs" edges to StockingLog entities.
func (mu *MedicineUpdate) RemoveStockingLogs(s ...*StockingLog) *MedicineUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveStockingLogIDs(ids...)
}

// ClearActiveIngredient clears the "active_ingredient" edge to the ActiveIngredient entity.
func (mu *MedicineUpdate) ClearActiveIngredient() *MedicineUpdate {
	mu.mutation.ClearActiveIngredient()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MedicineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MedicineUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MedicineUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MedicineUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MedicineUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Medicine.name": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Mah(); ok {
		if err := medicine.MahValidator(v); err != nil {
			return &ValidationError{Name: "mah", err: fmt.Errorf(`ent: validator failed for field "Medicine.mah": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Dosage(); ok {
		if err := medicine.DosageValidator(v); err != nil {
			return &ValidationError{Name: "dosage", err: fmt.Errorf(`ent: validator failed for field "Medicine.dosage": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Unit(); ok {
		if err := medicine.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "Medicine.unit": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Atc(); ok {
		if err := medicine.AtcValidator(v); err != nil {
			return &ValidationError{Name: "atc", err: fmt.Errorf(`ent: validator failed for field "Medicine.atc": %w`, err)}
		}
	}
	if v, ok := mu.mutation.BoxSize(); ok {
		if err := medicine.BoxSizeValidator(v); err != nil {
			return &ValidationError{Name: "box_size", err: fmt.Errorf(`ent: validator failed for field "Medicine.box_size": %w`, err)}
		}
	}
	return nil
}

func (mu *MedicineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(medicine.Table, medicine.Columns, sqlgraph.NewFieldSpec(medicine.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(medicine.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Mah(); ok {
		_spec.SetField(medicine.FieldMah, field.TypeString, value)
	}
	if value, ok := mu.mutation.Dosage(); ok {
		_spec.SetField(medicine.FieldDosage, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedDosage(); ok {
		_spec.AddField(medicine.FieldDosage, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.Unit(); ok {
		_spec.SetField(medicine.FieldUnit, field.TypeString, value)
	}
	if value, ok := mu.mutation.Atc(); ok {
		_spec.SetField(medicine.FieldAtc, field.TypeString, value)
	}
	if value, ok := mu.mutation.Package(); ok {
		_spec.SetField(medicine.FieldPackage, field.TypeString, value)
	}
	if value, ok := mu.mutation.Form(); ok {
		_spec.SetField(medicine.FieldForm, field.TypeString, value)
	}
	if value, ok := mu.mutation.BoxSize(); ok {
		_spec.SetField(medicine.FieldBoxSize, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedBoxSize(); ok {
		_spec.AddField(medicine.FieldBoxSize, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Stock(); ok {
		_spec.SetField(medicine.FieldStock, field.TypeFloat32, value)
	}
	if value, ok := mu.mutation.AddedStock(); ok {
		_spec.AddField(medicine.FieldStock, field.TypeFloat32, value)
	}
	if mu.mutation.StockCleared() {
		_spec.ClearField(medicine.FieldStock, field.TypeFloat32)
	}
	if value, ok := mu.mutation.LastStockUpdate(); ok {
		_spec.SetField(medicine.FieldLastStockUpdate, field.TypeTime, value)
	}
	if mu.mutation.PurchasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPurchasesIDs(); len(nodes) > 0 && !mu.mutation.PurchasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PurchasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.StockingLogsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedStockingLogsIDs(); len(nodes) > 0 && !mu.mutation.StockingLogsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.StockingLogsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ActiveIngredientCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ActiveIngredientIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MedicineUpdateOne is the builder for updating a single Medicine entity.
type MedicineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MedicineMutation
}

// SetName sets the "name" field.
func (muo *MedicineUpdateOne) SetName(s string) *MedicineUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableName(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetMah sets the "mah" field.
func (muo *MedicineUpdateOne) SetMah(s string) *MedicineUpdateOne {
	muo.mutation.SetMah(s)
	return muo
}

// SetNillableMah sets the "mah" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableMah(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetMah(*s)
	}
	return muo
}

// SetDosage sets the "dosage" field.
func (muo *MedicineUpdateOne) SetDosage(f float64) *MedicineUpdateOne {
	muo.mutation.ResetDosage()
	muo.mutation.SetDosage(f)
	return muo
}

// SetNillableDosage sets the "dosage" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableDosage(f *float64) *MedicineUpdateOne {
	if f != nil {
		muo.SetDosage(*f)
	}
	return muo
}

// AddDosage adds f to the "dosage" field.
func (muo *MedicineUpdateOne) AddDosage(f float64) *MedicineUpdateOne {
	muo.mutation.AddDosage(f)
	return muo
}

// SetUnit sets the "unit" field.
func (muo *MedicineUpdateOne) SetUnit(s string) *MedicineUpdateOne {
	muo.mutation.SetUnit(s)
	return muo
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableUnit(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetUnit(*s)
	}
	return muo
}

// SetAtc sets the "atc" field.
func (muo *MedicineUpdateOne) SetAtc(s string) *MedicineUpdateOne {
	muo.mutation.SetAtc(s)
	return muo
}

// SetNillableAtc sets the "atc" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableAtc(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetAtc(*s)
	}
	return muo
}

// SetPackage sets the "package" field.
func (muo *MedicineUpdateOne) SetPackage(s string) *MedicineUpdateOne {
	muo.mutation.SetPackage(s)
	return muo
}

// SetNillablePackage sets the "package" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillablePackage(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetPackage(*s)
	}
	return muo
}

// SetForm sets the "form" field.
func (muo *MedicineUpdateOne) SetForm(s string) *MedicineUpdateOne {
	muo.mutation.SetForm(s)
	return muo
}

// SetNillableForm sets the "form" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableForm(s *string) *MedicineUpdateOne {
	if s != nil {
		muo.SetForm(*s)
	}
	return muo
}

// SetBoxSize sets the "box_size" field.
func (muo *MedicineUpdateOne) SetBoxSize(i int) *MedicineUpdateOne {
	muo.mutation.ResetBoxSize()
	muo.mutation.SetBoxSize(i)
	return muo
}

// SetNillableBoxSize sets the "box_size" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableBoxSize(i *int) *MedicineUpdateOne {
	if i != nil {
		muo.SetBoxSize(*i)
	}
	return muo
}

// AddBoxSize adds i to the "box_size" field.
func (muo *MedicineUpdateOne) AddBoxSize(i int) *MedicineUpdateOne {
	muo.mutation.AddBoxSize(i)
	return muo
}

// SetStock sets the "stock" field.
func (muo *MedicineUpdateOne) SetStock(f float32) *MedicineUpdateOne {
	muo.mutation.ResetStock()
	muo.mutation.SetStock(f)
	return muo
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableStock(f *float32) *MedicineUpdateOne {
	if f != nil {
		muo.SetStock(*f)
	}
	return muo
}

// AddStock adds f to the "stock" field.
func (muo *MedicineUpdateOne) AddStock(f float32) *MedicineUpdateOne {
	muo.mutation.AddStock(f)
	return muo
}

// ClearStock clears the value of the "stock" field.
func (muo *MedicineUpdateOne) ClearStock() *MedicineUpdateOne {
	muo.mutation.ClearStock()
	return muo
}

// SetLastStockUpdate sets the "last_stock_update" field.
func (muo *MedicineUpdateOne) SetLastStockUpdate(t time.Time) *MedicineUpdateOne {
	muo.mutation.SetLastStockUpdate(t)
	return muo
}

// SetNillableLastStockUpdate sets the "last_stock_update" field if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableLastStockUpdate(t *time.Time) *MedicineUpdateOne {
	if t != nil {
		muo.SetLastStockUpdate(*t)
	}
	return muo
}

// AddPurchaseIDs adds the "purchases" edge to the Purchase entity by IDs.
func (muo *MedicineUpdateOne) AddPurchaseIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.AddPurchaseIDs(ids...)
	return muo
}

// AddPurchases adds the "purchases" edges to the Purchase entity.
func (muo *MedicineUpdateOne) AddPurchases(p ...*Purchase) *MedicineUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPurchaseIDs(ids...)
}

// AddStockingLogIDs adds the "stocking_logs" edge to the StockingLog entity by IDs.
func (muo *MedicineUpdateOne) AddStockingLogIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.AddStockingLogIDs(ids...)
	return muo
}

// AddStockingLogs adds the "stocking_logs" edges to the StockingLog entity.
func (muo *MedicineUpdateOne) AddStockingLogs(s ...*StockingLog) *MedicineUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddStockingLogIDs(ids...)
}

// SetActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID.
func (muo *MedicineUpdateOne) SetActiveIngredientID(id int) *MedicineUpdateOne {
	muo.mutation.SetActiveIngredientID(id)
	return muo
}

// SetNillableActiveIngredientID sets the "active_ingredient" edge to the ActiveIngredient entity by ID if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableActiveIngredientID(id *int) *MedicineUpdateOne {
	if id != nil {
		muo = muo.SetActiveIngredientID(*id)
	}
	return muo
}

// SetActiveIngredient sets the "active_ingredient" edge to the ActiveIngredient entity.
func (muo *MedicineUpdateOne) SetActiveIngredient(a *ActiveIngredient) *MedicineUpdateOne {
	return muo.SetActiveIngredientID(a.ID)
}

// Mutation returns the MedicineMutation object of the builder.
func (muo *MedicineUpdateOne) Mutation() *MedicineMutation {
	return muo.mutation
}

// ClearPurchases clears all "purchases" edges to the Purchase entity.
func (muo *MedicineUpdateOne) ClearPurchases() *MedicineUpdateOne {
	muo.mutation.ClearPurchases()
	return muo
}

// RemovePurchaseIDs removes the "purchases" edge to Purchase entities by IDs.
func (muo *MedicineUpdateOne) RemovePurchaseIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.RemovePurchaseIDs(ids...)
	return muo
}

// RemovePurchases removes "purchases" edges to Purchase entities.
func (muo *MedicineUpdateOne) RemovePurchases(p ...*Purchase) *MedicineUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePurchaseIDs(ids...)
}

// ClearStockingLogs clears all "stocking_logs" edges to the StockingLog entity.
func (muo *MedicineUpdateOne) ClearStockingLogs() *MedicineUpdateOne {
	muo.mutation.ClearStockingLogs()
	return muo
}

// RemoveStockingLogIDs removes the "stocking_logs" edge to StockingLog entities by IDs.
func (muo *MedicineUpdateOne) RemoveStockingLogIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.RemoveStockingLogIDs(ids...)
	return muo
}

// RemoveStockingLogs removes "stocking_logs" edges to StockingLog entities.
func (muo *MedicineUpdateOne) RemoveStockingLogs(s ...*StockingLog) *MedicineUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveStockingLogIDs(ids...)
}

// ClearActiveIngredient clears the "active_ingredient" edge to the ActiveIngredient entity.
func (muo *MedicineUpdateOne) ClearActiveIngredient() *MedicineUpdateOne {
	muo.mutation.ClearActiveIngredient()
	return muo
}

// Where appends a list predicates to the MedicineUpdate builder.
func (muo *MedicineUpdateOne) Where(ps ...predicate.Medicine) *MedicineUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MedicineUpdateOne) Select(field string, fields ...string) *MedicineUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Medicine entity.
func (muo *MedicineUpdateOne) Save(ctx context.Context) (*Medicine, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MedicineUpdateOne) SaveX(ctx context.Context) *Medicine {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MedicineUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MedicineUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MedicineUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Medicine.name": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Mah(); ok {
		if err := medicine.MahValidator(v); err != nil {
			return &ValidationError{Name: "mah", err: fmt.Errorf(`ent: validator failed for field "Medicine.mah": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Dosage(); ok {
		if err := medicine.DosageValidator(v); err != nil {
			return &ValidationError{Name: "dosage", err: fmt.Errorf(`ent: validator failed for field "Medicine.dosage": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Unit(); ok {
		if err := medicine.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "Medicine.unit": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Atc(); ok {
		if err := medicine.AtcValidator(v); err != nil {
			return &ValidationError{Name: "atc", err: fmt.Errorf(`ent: validator failed for field "Medicine.atc": %w`, err)}
		}
	}
	if v, ok := muo.mutation.BoxSize(); ok {
		if err := medicine.BoxSizeValidator(v); err != nil {
			return &ValidationError{Name: "box_size", err: fmt.Errorf(`ent: validator failed for field "Medicine.box_size": %w`, err)}
		}
	}
	return nil
}

func (muo *MedicineUpdateOne) sqlSave(ctx context.Context) (_node *Medicine, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(medicine.Table, medicine.Columns, sqlgraph.NewFieldSpec(medicine.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Medicine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, medicine.FieldID)
		for _, f := range fields {
			if !medicine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != medicine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(medicine.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Mah(); ok {
		_spec.SetField(medicine.FieldMah, field.TypeString, value)
	}
	if value, ok := muo.mutation.Dosage(); ok {
		_spec.SetField(medicine.FieldDosage, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedDosage(); ok {
		_spec.AddField(medicine.FieldDosage, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.Unit(); ok {
		_spec.SetField(medicine.FieldUnit, field.TypeString, value)
	}
	if value, ok := muo.mutation.Atc(); ok {
		_spec.SetField(medicine.FieldAtc, field.TypeString, value)
	}
	if value, ok := muo.mutation.Package(); ok {
		_spec.SetField(medicine.FieldPackage, field.TypeString, value)
	}
	if value, ok := muo.mutation.Form(); ok {
		_spec.SetField(medicine.FieldForm, field.TypeString, value)
	}
	if value, ok := muo.mutation.BoxSize(); ok {
		_spec.SetField(medicine.FieldBoxSize, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedBoxSize(); ok {
		_spec.AddField(medicine.FieldBoxSize, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Stock(); ok {
		_spec.SetField(medicine.FieldStock, field.TypeFloat32, value)
	}
	if value, ok := muo.mutation.AddedStock(); ok {
		_spec.AddField(medicine.FieldStock, field.TypeFloat32, value)
	}
	if muo.mutation.StockCleared() {
		_spec.ClearField(medicine.FieldStock, field.TypeFloat32)
	}
	if value, ok := muo.mutation.LastStockUpdate(); ok {
		_spec.SetField(medicine.FieldLastStockUpdate, field.TypeTime, value)
	}
	if muo.mutation.PurchasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPurchasesIDs(); len(nodes) > 0 && !muo.mutation.PurchasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PurchasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.PurchasesTable,
			Columns: []string{medicine.PurchasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.StockingLogsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedStockingLogsIDs(); len(nodes) > 0 && !muo.mutation.StockingLogsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.StockingLogsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ActiveIngredientCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ActiveIngredientIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Medicine{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}

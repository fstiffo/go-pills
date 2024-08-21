// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConsumptionLogDelete is the builder for deleting a ConsumptionLog entity.
type ConsumptionLogDelete struct {
	config
	hooks    []Hook
	mutation *ConsumptionLogMutation
}

// Where appends a list predicates to the ConsumptionLogDelete builder.
func (cld *ConsumptionLogDelete) Where(ps ...predicate.ConsumptionLog) *ConsumptionLogDelete {
	cld.mutation.Where(ps...)
	return cld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cld *ConsumptionLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cld.sqlExec, cld.mutation, cld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cld *ConsumptionLogDelete) ExecX(ctx context.Context) int {
	n, err := cld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cld *ConsumptionLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(consumptionlog.Table, sqlgraph.NewFieldSpec(consumptionlog.FieldID, field.TypeInt))
	if ps := cld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cld.mutation.done = true
	return affected, err
}

// ConsumptionLogDeleteOne is the builder for deleting a single ConsumptionLog entity.
type ConsumptionLogDeleteOne struct {
	cld *ConsumptionLogDelete
}

// Where appends a list predicates to the ConsumptionLogDelete builder.
func (cldo *ConsumptionLogDeleteOne) Where(ps ...predicate.ConsumptionLog) *ConsumptionLogDeleteOne {
	cldo.cld.mutation.Where(ps...)
	return cldo
}

// Exec executes the deletion query.
func (cldo *ConsumptionLogDeleteOne) Exec(ctx context.Context) error {
	n, err := cldo.cld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{consumptionlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cldo *ConsumptionLogDeleteOne) ExecX(ctx context.Context) {
	if err := cldo.Exec(ctx); err != nil {
		panic(err)
	}
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/prescription"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ConsumptionLog is the model entity for the ConsumptionLog schema.
type ConsumptionLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ConsumedAt holds the value of the "consumed_at" field.
	ConsumedAt time.Time `json:"consumed_at,omitempty"`
	// Units holds the value of the "units" field.
	Units int `json:"units,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConsumptionLogQuery when eager-loading is set.
	Edges                              ConsumptionLogEdges `json:"edges"`
	active_ingredient_consumption_logs *int
	prescription_comsumption_logs      *int
	selectValues                       sql.SelectValues
}

// ConsumptionLogEdges holds the relations/edges for other nodes in the graph.
type ConsumptionLogEdges struct {
	// Prescription holds the value of the prescription edge.
	Prescription *Prescription `json:"prescription,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PrescriptionOrErr returns the Prescription value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConsumptionLogEdges) PrescriptionOrErr() (*Prescription, error) {
	if e.Prescription != nil {
		return e.Prescription, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: prescription.Label}
	}
	return nil, &NotLoadedError{edge: "prescription"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ConsumptionLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case consumptionlog.FieldID, consumptionlog.FieldUnits:
			values[i] = new(sql.NullInt64)
		case consumptionlog.FieldConsumedAt:
			values[i] = new(sql.NullTime)
		case consumptionlog.ForeignKeys[0]: // active_ingredient_consumption_logs
			values[i] = new(sql.NullInt64)
		case consumptionlog.ForeignKeys[1]: // prescription_comsumption_logs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ConsumptionLog fields.
func (cl *ConsumptionLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case consumptionlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cl.ID = int(value.Int64)
		case consumptionlog.FieldConsumedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field consumed_at", values[i])
			} else if value.Valid {
				cl.ConsumedAt = value.Time
			}
		case consumptionlog.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				cl.Units = int(value.Int64)
			}
		case consumptionlog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field active_ingredient_consumption_logs", value)
			} else if value.Valid {
				cl.active_ingredient_consumption_logs = new(int)
				*cl.active_ingredient_consumption_logs = int(value.Int64)
			}
		case consumptionlog.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field prescription_comsumption_logs", value)
			} else if value.Valid {
				cl.prescription_comsumption_logs = new(int)
				*cl.prescription_comsumption_logs = int(value.Int64)
			}
		default:
			cl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ConsumptionLog.
// This includes values selected through modifiers, order, etc.
func (cl *ConsumptionLog) Value(name string) (ent.Value, error) {
	return cl.selectValues.Get(name)
}

// QueryPrescription queries the "prescription" edge of the ConsumptionLog entity.
func (cl *ConsumptionLog) QueryPrescription() *PrescriptionQuery {
	return NewConsumptionLogClient(cl.config).QueryPrescription(cl)
}

// Update returns a builder for updating this ConsumptionLog.
// Note that you need to call ConsumptionLog.Unwrap() before calling this method if this ConsumptionLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (cl *ConsumptionLog) Update() *ConsumptionLogUpdateOne {
	return NewConsumptionLogClient(cl.config).UpdateOne(cl)
}

// Unwrap unwraps the ConsumptionLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cl *ConsumptionLog) Unwrap() *ConsumptionLog {
	_tx, ok := cl.config.driver.(*txDriver)
	if !ok {
		panic("ent: ConsumptionLog is not a transactional entity")
	}
	cl.config.driver = _tx.drv
	return cl
}

// String implements the fmt.Stringer.
func (cl *ConsumptionLog) String() string {
	var builder strings.Builder
	builder.WriteString("ConsumptionLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cl.ID))
	builder.WriteString("consumed_at=")
	builder.WriteString(cl.ConsumedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", cl.Units))
	builder.WriteByte(')')
	return builder.String()
}

// ConsumptionLogs is a parsable slice of ConsumptionLog.
type ConsumptionLogs []*ConsumptionLog

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActiveIngredientsColumns holds the columns for the "active_ingredients" table.
	ActiveIngredientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ActiveIngredientsTable holds the schema information for the "active_ingredients" table.
	ActiveIngredientsTable = &schema.Table{
		Name:       "active_ingredients",
		Columns:    ActiveIngredientsColumns,
		PrimaryKey: []*schema.Column{ActiveIngredientsColumns[0]},
	}
	// ConsumptionLogsColumns holds the columns for the "consumption_logs" table.
	ConsumptionLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "consumed_at", Type: field.TypeTime},
		{Name: "prescription_comsumption_logs", Type: field.TypeInt, Nullable: true},
	}
	// ConsumptionLogsTable holds the schema information for the "consumption_logs" table.
	ConsumptionLogsTable = &schema.Table{
		Name:       "consumption_logs",
		Columns:    ConsumptionLogsColumns,
		PrimaryKey: []*schema.Column{ConsumptionLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "consumption_logs_prescriptions_comsumption_logs",
				Columns:    []*schema.Column{ConsumptionLogsColumns[2]},
				RefColumns: []*schema.Column{PrescriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MedicinesColumns holds the columns for the "medicines" table.
	MedicinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "mah", Type: field.TypeString},
		{Name: "dosage", Type: field.TypeFloat64},
		{Name: "unit", Type: field.TypeString},
		{Name: "atc", Type: field.TypeString, Unique: true},
		{Name: "package", Type: field.TypeString},
		{Name: "form", Type: field.TypeString},
		{Name: "box_size", Type: field.TypeInt},
		{Name: "stock", Type: field.TypeFloat32, Nullable: true, Default: 0},
		{Name: "last_stock_update", Type: field.TypeTime},
		{Name: "active_ingredient_medicines", Type: field.TypeInt, Nullable: true},
	}
	// MedicinesTable holds the schema information for the "medicines" table.
	MedicinesTable = &schema.Table{
		Name:       "medicines",
		Columns:    MedicinesColumns,
		PrimaryKey: []*schema.Column{MedicinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "medicines_active_ingredients_medicines",
				Columns:    []*schema.Column{MedicinesColumns[11]},
				RefColumns: []*schema.Column{ActiveIngredientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrescriptionsColumns holds the columns for the "prescriptions" table.
	PrescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dosage", Type: field.TypeInt},
		{Name: "unit", Type: field.TypeString},
		{Name: "dosage_frequency", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "active_ingredient_prescriptions", Type: field.TypeInt, Nullable: true},
	}
	// PrescriptionsTable holds the schema information for the "prescriptions" table.
	PrescriptionsTable = &schema.Table{
		Name:       "prescriptions",
		Columns:    PrescriptionsColumns,
		PrimaryKey: []*schema.Column{PrescriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prescriptions_active_ingredients_prescriptions",
				Columns:    []*schema.Column{PrescriptionsColumns[6]},
				RefColumns: []*schema.Column{ActiveIngredientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PurchasesColumns holds the columns for the "purchases" table.
	PurchasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "puchased_at", Type: field.TypeTime},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "medicine_purchases", Type: field.TypeInt, Nullable: true},
	}
	// PurchasesTable holds the schema information for the "purchases" table.
	PurchasesTable = &schema.Table{
		Name:       "purchases",
		Columns:    PurchasesColumns,
		PrimaryKey: []*schema.Column{PurchasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "purchases_medicines_purchases",
				Columns:    []*schema.Column{PurchasesColumns[3]},
				RefColumns: []*schema.Column{MedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StockingLogsColumns holds the columns for the "stocking_logs" table.
	StockingLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "stocked_at", Type: field.TypeTime},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "medicine_stocking_logs", Type: field.TypeInt, Nullable: true},
	}
	// StockingLogsTable holds the schema information for the "stocking_logs" table.
	StockingLogsTable = &schema.Table{
		Name:       "stocking_logs",
		Columns:    StockingLogsColumns,
		PrimaryKey: []*schema.Column{StockingLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stocking_logs_medicines_stocking_logs",
				Columns:    []*schema.Column{StockingLogsColumns[3]},
				RefColumns: []*schema.Column{MedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActiveIngredientsTable,
		ConsumptionLogsTable,
		MedicinesTable,
		PrescriptionsTable,
		PurchasesTable,
		StockingLogsTable,
	}
)

func init() {
	ConsumptionLogsTable.ForeignKeys[0].RefTable = PrescriptionsTable
	MedicinesTable.ForeignKeys[0].RefTable = ActiveIngredientsTable
	PrescriptionsTable.ForeignKeys[0].RefTable = ActiveIngredientsTable
	PurchasesTable.ForeignKeys[0].RefTable = MedicinesTable
	StockingLogsTable.ForeignKeys[0].RefTable = MedicinesTable
}

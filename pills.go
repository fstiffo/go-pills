package main

import (
	"context"
	"fmt"
	"fstiffo/pills/ent"
	"fstiffo/pills/ent/consumptionlog"
	"fstiffo/pills/ent/medicine"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("ent: one-to-many (O2M) association")
	client, err := ent.Open("postgres", "host=aws-0-eu-central-1.pooler.supabase.com port=6543 user=postgres.vnnadpjjyjzlvhtxyubb dbname=postgres password=Q5G66QriNLTVN0gO")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create a new medicine.
	// _, err = CreateMedicine(context.Background(), client)
	// if err != nil {
	// 	log.Fatalf("failed creating medicine: %v", err)
	// }

	// // Query a medicine.
	// ai, error := client.ActiveIngredient.
	// 	Query().
	// 	Where(activeingredient.Name("METOPROLOLO")).
	// 	Only(context.Background())
	// if error != nil {
	// 	log.Fatalf("failed querying active ingredient: %v", error)
	// }
	// QueryMedicinesIngredient(context.Background(), ai)
}

// CreateActiveIngredient creates a new active ingredient.
// This function creates a new active ingredient in the database.
func CreateActiveIngredient(ctx context.Context, client *ent.Client) (*ent.ActiveIngredient, error) {
	ai, err := client.ActiveIngredient.
		Create().
		SetName("METOPROLOLO").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating active ingredient: %v", err)
	}
	log.Println("active ingredient was created: ", ai)
	return ai, nil
}

// CreateMedicine creates a new medicine.
// This function creates a new medicine in the database and add it to an active ingredient.
func CreateMedicine(ctx context.Context, client *ent.Client) (*ent.ActiveIngredient, error) {
	m, err := client.Medicine.
		Create().
		SetName("METOPROLOLO AUROBINDO").
		SetMah("AUROBINDO PHARMA (ITALIA) S.R.L.").
		SetDosage(100).
		SetAtc("C07AB02 - METOPROLOLO").
		SetPackage("BLISTER PVC/PVDC/AL").
		SetForm("Compresse rivestite con film").
		SetBoxSize(30).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating medicine: %v", err)
	}
	log.Println("medicine was created: ", m)

	ai, err := client.ActiveIngredient.
		Create().
		SetName("METOPROLOLO").
		AddMedicines(m).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating active ingredient: %v", err)
	}
	log.Println("active ingredient was created: ", ai)

	return ai, nil
}

// QueryMedicine queries the medicine with the given name.
// This function returns the medicine with the given name.
func QueryMedicine(ctx context.Context, client *ent.Client, name string) (*ent.Medicine, error) {
	m, err := client.Medicine.
		Query().
		Where(medicine.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying medicine: %v", err)
	}
	log.Println("medicine returned: ", m)
	return m, nil
}

// QueryMedicinesIngredient queries the active ingredients.
// This function returns the active ingredient.
func QueryMedicinesIngredient(ctx context.Context, ai *ent.ActiveIngredient) error {
	medicines, err := ai.QueryMedicines().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying medicine active ingredient: %v", err)
	}
	for _, m := range medicines {
		ai, err := m.QueryActiveIngredient().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying medicine %q active ingredient: %w", m.Name, err)
		}
		log.Printf("medicine %q active ingredient %q:", m.Name, ai.Name)
	}
	return nil
}

// LastTakenAt returns the last taken consumption log.
func LastTakenAt(ctx context.Context, pr *ent.Prescription) (time.Time, error) {
	log, err := pr.QueryComsumptionLogs().Order(ent.Desc(consumptionlog.FieldConsumedAt)).First(ctx)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed querying last consumption log: %v", err)
	}
	return log.ConsumedAt, nil
}

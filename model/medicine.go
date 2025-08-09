package model

import "gorm.io/gorm"

// InsertActiveIngredient inserts a new ActiveIngredient record.
func InsertActiveIngredient(db *gorm.DB, ai *ActiveIngredient) error {
	return db.Create(ai).Error
}

// InsertMedicine inserts a new Medicine record.
func InsertMedicine(db *gorm.DB, med *Medicine) error {
	return db.Create(med).Error
}

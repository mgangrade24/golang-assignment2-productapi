package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

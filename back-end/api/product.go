package main

import "gorm.io/gorm"

var ProductDB *gorm.DB

type Product struct {
	gorm.Model
}

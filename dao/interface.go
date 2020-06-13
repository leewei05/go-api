package dao

import "time"

// Dao is an interface
type Dao interface {
	Get() (*Product, error)
	Create(*Product) error
	Update()
	Delete()
}

// Product is a struct
type Product struct {
	// ID is Product's primary key.
	ID uint64 `gorm:"primary_key"`
	// Name is Product's name.
	Name string `gorm:"size:150" sql:"not null"`
	// Price is Product's pricing.
	Price float64 `sql:"not null"`
	// CreatedAt for gorm to insert create time.
	CreatedAt time.Time
	// UpdatedAt for gorm to insert update time.
	UpdatedAt time.Time
}

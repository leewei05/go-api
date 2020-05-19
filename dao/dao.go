package dao

import "github.com/jinzhu/gorm"

type dao struct {
	db *gorm.DB
}

// NewDao is a function that defines a new dao instance
func NewDao(db *gorm.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) Get() {}

func (d *dao) Create() {}

func (d *dao) Update() {}

func (d *dao) Delete() {}

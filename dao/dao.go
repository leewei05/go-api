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

func (d *dao) Get() (*Product, error) {
	var product Product

	if err := d.db.Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (d *dao) Create(p Product) (*Product, error) {

	return &Product{}, nil
}

func (d *dao) Update() {}

func (d *dao) Delete() {}

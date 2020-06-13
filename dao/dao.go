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
	var p Product

	if err := d.db.Find(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (d *dao) Create(p *Product) error {
	if err := d.db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (d *dao) Update() {}

func (d *dao) Delete() {}

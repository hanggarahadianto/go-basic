package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct{
	ID					string
	SKU					string
	Name				string
	Slug				string
	Price				decimal.Decimal
	Stock				int64
	Weight				decimal.Decimal
	ShortDescription 	string
	Description			string
	Status				int
	CreatedAt			time.Time
	UpdatedAt			time.Time
	ProductImages		[]ProductImage
	Categories			[]Category `gorm:"many2many:product_categories"`
	ParentId			string
	User				User
	UserID				string



}
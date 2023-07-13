package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct{

	ID				string
	Qty				int
	BasePrice		decimal.Decimal
	BaseTotal		decimal.Decimal
	TaxAmount		decimal.Decimal
	TaxPercent		decimal.Decimal
	DiscountAmount	decimal.Decimal
	DiscountPercent	decimal.Decimal
	SubTotal		decimal.Decimal
	SKU				string
	AppName			string
	Weight			decimal.Decimal
	Order			Order
	OrderID			string
	Product			Product
	ProductID		string
	CreatedAt		time.Time
	UpdatedAt		time.Time


}
package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct{

	ID			string
	Number		string
	Amount 		decimal.Decimal
	Method		string
	Status		string
	Token 		string
	Payload		string
	PaymentType	string
	VaNumber	string
	BillCode	string
	BillKey		string
	Order		Order
	OrderID		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletdAt	gorm.DeletedAt

}
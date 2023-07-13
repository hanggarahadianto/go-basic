package models

import (
	"time"

	"github.com/shopspring/decimal"
)
type Order struct{
	ID					string
	Code 				string
	Status				int
	OrderDate			time.Time
	PaymentDue			time.Time
	PaymentStatus		string
	PaymentToken		string
	BaseTotalPrice		decimal.Decimal
	TaxAmount			decimal.Decimal
	TaxPercent			decimal.Decimal
	DiscountAmount		decimal.Decimal
	ShoppingCost		decimal.Decimal
	GrandTotal			decimal.Decimal
	Note 				string
	ShippingCourier		string
	ShippingServiceName	string
	ApprovedBy			string
	ApproveAt			time.Time
	CancelledBy			string
	CencelledAt			time.Time
	CencellationNote	string
	OrderCustomer		*OrderCustomer
	OrderItems			[]OrderItem
	UserID				string
	User				User
	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			time.Time

	
}
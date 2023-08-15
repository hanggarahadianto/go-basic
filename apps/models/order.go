package models

import (
	"time"

	"github.com/google/uuid"
)


type Order struct{
	ID			uuid.UUID			`json:"order_id"`
	Price		int
	Discount	int
	Payment		Payment
	OrderAt		time.Time

}
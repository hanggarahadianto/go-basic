package models

import "github.com/google/uuid"



type ProductUser struct{

	ID			uuid.UUID	`json:"product_id"`
	ProductName	string		`json:"product_name"`
	Price		uint64		`json:"price"`
	Rating		uint8		`json:"rating"`
	Image		string		`json:"image"`
}
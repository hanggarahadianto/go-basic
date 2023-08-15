package models

import "github.com/google/uuid"

type Address struct{

	ID			uuid.UUID			`json:"adrdess_id"`
	House		string
	Street		string
	City		string
	Pincode		string		

}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct{
	ID				string
	TrackNumber		string
	Status			string
	TotalQty		int
	TotalWeight		string
	FirstName		string
	LastName		string
	CityID			string
	ProvienceID		string
	Address1		string
	Address2		string
	Phone			string
	Email			string
	PostCode		string
	ShippedBy		string
	ShippedAt		time.Time
	User			User
	UserID			string
	Order			Order
	OrderID			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt

}
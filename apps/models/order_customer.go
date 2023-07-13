package models

import "time"


type OrderCustomer struct{
	
	ID			string
	FirstName	string
	LastName	string
	CityID		string
	ProvienceID	string
	Address1	string
	Address2	string
	Phone		string
	Email		string
	PostCode	string
	User		User
	UserID		string
	Order		Order
	OrderID		string
	CreatedAt	time.Time
	UpdatedAt	time.Time


}
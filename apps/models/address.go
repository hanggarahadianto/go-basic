package models

import "time"

type Address struct {	
	ID			string
	Name 		string
	IsPrimary	bool
	ProvienceID	string
	CityID		string
	Address1	string
	Address2	string
	Phone		string
	Email		string
	PostCode	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	User User
	UserID 		string

}
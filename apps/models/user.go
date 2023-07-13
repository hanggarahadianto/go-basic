package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	ID				string
	FirstName 		string
	LastName		string
	Email			string
	Password		string
	RememberToken	string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt
	Address			[]Address


}
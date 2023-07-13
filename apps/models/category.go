package models

import "time"

type Category struct{

	ID			string
	Name 		string
	Slug		string
	ParentID		string
	Section		Section
	SectionID	string
	Products	[]Product		`gorm:"many2many:product_categories;"`
	CreatedAt	time.Time
	UpdatedAt	time.Time


}
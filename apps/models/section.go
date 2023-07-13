package models

import "time"

type Section struct{
	ID			string 
	Name		string
	Slug		string
	Categories	[]Category
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
package models

type ProductImage struct {
	ID			string
	Path 		string
	ExtraLarge	string
	Large 		string
	Medium		string
	Small		string
	Product		Product
	ProductID	string
	
}
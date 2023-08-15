package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct{

	ID        			uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName			string		`json:"first_name" gorm:"not null"`
	LastName			string		`json:"last_name" gorm:"not null"`
	Password			string		`json:"password" gorm:"not null"`
	Email				string		`json:"email" gorm:"not null"`
	Phone				string		`json:"phone" gorm:"not null"`
	PasswordResetToken	string
	PasswordResetAt		time.Time
	CreatedAt			time.Time
	UpdatedAt			time.Time
	
}


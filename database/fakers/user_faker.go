package fakers

import (
	"origin/apps/models"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
) 


func UserFaker(db *gorm.DB) *models.User{
	return &models.User{
		ID: uuid.New().String(),
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Password: faker.Password(),


		
	}
}
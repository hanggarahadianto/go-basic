package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB(){
	db, err := gorm.Open(mysql.Open("root:12345678/go"), &gorm.Config{})

	if err != nil{
		fmt.Println(err.Error())
		panic("filed to connect db")
	}
	
	DB = db

}


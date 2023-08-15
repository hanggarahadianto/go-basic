package main

import (
	"fmt"
	"go-basic/db"
	"log"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine


func init(){
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()

}

func main() {
	db.RedisInit()

	fmt.Println("running on port 8080")
	log.Fatal(server.Run(":8080"))


}
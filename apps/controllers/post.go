package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"go-basic/apps/models"
// 	"go-basic/db"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var key = "iagijglsk93302002jkn2knnvs"

// func AddPost(c *gin.Context){

// 	var postData models.Post

// 	rdb:= db.RedisConnect()
// 	newPost := models.Post{
// 		Title: title,
// 	}

// 	req, _:= json.Marshal(newPost)

// 	err := rdb.HSet(key, id, req).Err()
// 	if err != nil {
// 		fmt.Printf("error set redis %s", err)
// 		return
// 	}
// 	// return

// 	c.JSON(http.StatusOK, gin.H{
// 		"status" : "success",
// 		"data" : newPost,
// 	})
// }
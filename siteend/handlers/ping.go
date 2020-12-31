package handlers

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func Ping(context *gin.Context){
	fmt.Println("hello world!")
	context.JSON(200, gin.H{"message": "pong"})
}
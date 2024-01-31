package main

import (
	"practice/customer-labs-test/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/sendData", controllers.SendData)
	r.Run(":9092")
}

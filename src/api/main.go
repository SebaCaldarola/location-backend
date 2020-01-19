package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sebastian/location-backend/src/api/controllers"
	"net/http"
)

var (
	Router *gin.Engine
)

func main() {
	StartApplication()
}


func StartApplication(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/locations", controllers.UpdateLocationData)

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
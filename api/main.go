package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantomlinson/shipit/api/controllers"
)

func main() {
	router := gin.Default()

	features := new(controllers.FeatureController)

	v1 := router.Group("/api/v1/features")
	{
		v1.GET("/", features.All)
	}
}

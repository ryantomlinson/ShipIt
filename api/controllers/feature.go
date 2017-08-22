package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantomlinson/shipit/api/model"
)

type FeatureController struct {}

var featureModel = new(model.FeatureFlag)

func (f *FeatureController) All(c *gin.Context) {
	data, err := featureModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the feature toggles", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}
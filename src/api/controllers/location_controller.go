package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sebastian/location-backend/src/api/core"
	"io/ioutil"
	"net/http"
)

func UpdateLocationData (c *gin.Context) {
	location := new(core.Location)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &location)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

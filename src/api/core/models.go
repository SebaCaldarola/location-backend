package core

import "github.com/sebastian/location-backend/src/api/utils"

type Location struct {
	Model utils.Model
	Latitude float64
	Longitude float64
	DriverId int
}
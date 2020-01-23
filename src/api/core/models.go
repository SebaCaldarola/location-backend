package core

import "github.com/sebastian/location-backend/src/api/utils"

type Location struct {
	Model     utils.Model
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DriverId  int     `json:"driver_id"`
}

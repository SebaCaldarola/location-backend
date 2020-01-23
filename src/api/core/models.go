package core

import "github.com/sebastian/location-backend/src/api/utils"

type Location struct {
	Model     utils.Model
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	DriverId  int     `json:"driver_id"`
}
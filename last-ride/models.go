package function

import "time"

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Ride struct {
	PassengerID uint      `json:"passengerID"`
	Time        time.Time `json:"time"`
	Origin      Point     `json:"origin"`
	Destination Point     `json:"destination"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}

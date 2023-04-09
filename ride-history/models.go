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

type RideSummary struct {
	Time        time.Time `json:"time"`
	Destination Point     `json:"destination"`
}

type RideHistory struct {
	Rides []RideSummary `json:"rides"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}

package function

import "time"

const (
	RecommendationRepeat  = "repeat_ride"
	RecommendationReverse = "reverse_ride"
	RecommendationNothing = "no_ride"
)

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

type UserInfo struct {
	ID                     uint     `json:"id"`
	FirstName              string   `json:"first_name"`
	LastName               string   `json:"last_name"`
	PhoneNumber            string   `json:"phone_number"`
	CurrentAddressLocation Point    `json:"current_address_location"`
	Addresses              []string `json:"addresses"`
}

type Recommendation struct {
	Type           string `json:"type"`
	Recommendation *Point `json:"recommendation"`
	BannerText     string `json:"banner_text"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint  `json:"user_id"`
	Origin *Point `json:"origin"`
}

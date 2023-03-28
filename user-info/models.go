package function

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type UserInfo struct {
	ID                     uint     `json:"id"`
	FirstName              string   `json:"first_name"`
	LastName               string   `json:"last_name"`
	PhoneNumber            string   `json:"phone_number"`
	CurrentAddressLocation Point    `json:"current_address_location"`
	Addresses              []string `json:"addresses"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}

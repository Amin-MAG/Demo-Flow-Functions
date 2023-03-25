package function

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type UserInfo struct {
	ID              uint   `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	CurrentLocation Point  `json:"current_location"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}

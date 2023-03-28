package function

import (
	"errors"
	"math/rand"
	"time"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	// Add a sleep time for simulating database connection
	time.Sleep(time.Duration(rand.Intn(10)+50) * time.Millisecond)

	userInfo := UserInfo{
		ID: *request.Args.UserID,
	}
	switch *request.Args.UserID {
	case 10:
		userInfo = UserInfo{
			FirstName:   "Amin",
			LastName:    "Ghasvari",
			PhoneNumber: "09336205449",
			CurrentAddressLocation: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
			Addresses: []string{"Azadegan ST. 13"},
		}
		break
	case 20:
		userInfo = UserInfo{
			FirstName:   "Arman",
			LastName:    "Heydari",
			PhoneNumber: "09362817764",
			CurrentAddressLocation: Point{
				Lat: 22.11,
				Lon: 22.22,
			},
			Addresses: []string{"Saadat abad ST. 52"},
		}
		break
	case 30:
		userInfo = UserInfo{
			FirstName:   "Masoud",
			LastName:    "Golestane",
			PhoneNumber: "09197846219",
			CurrentAddressLocation: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
			Addresses: []string{"Fadak ST. 92", "Tajrish ST. 32"},
		}
		break
	case 40:
		userInfo = UserInfo{
			FirstName:   "Ali",
			LastName:    "Sedaghi",
			PhoneNumber: "09376199092",
			CurrentAddressLocation: Point{
				Lat: 44.11,
				Lon: 44.22,
			},
			Addresses: []string{"Fallah ST. 11", "Gheytarieh ST. 82"},
		}
		break
	default:
		return nil, errors.New("can not find the user")
	}

	return createFlowOutput("user_info", userInfo)
}

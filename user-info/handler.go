package function

import (
	"errors"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	userInfo := UserInfo{
		ID: *request.Args.UserID,
	}
	switch *request.Args.UserID {
	case 10:
		userInfo = UserInfo{
			FirstName:   "Amin",
			LastName:    "Ghasvari",
			PhoneNumber: "09336205449",
			CurrentLocation: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
		}
		break
	case 20:
		userInfo = UserInfo{
			FirstName:   "Arman",
			LastName:    "Heydari",
			PhoneNumber: "09362817764",
			CurrentLocation: Point{
				Lat: 22.11,
				Lon: 22.22,
			},
		}
		break
	case 30:
		userInfo = UserInfo{
			FirstName:   "Masoud",
			LastName:    "Golestane",
			PhoneNumber: "09197846219",
			CurrentLocation: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
		}
		break
	case 40:
		userInfo = UserInfo{
			FirstName:   "Ali",
			LastName:    "Sedaghi",
			PhoneNumber: "09376199092",
			CurrentLocation: Point{
				Lat: 44.11,
				Lon: 44.22,
			},
		}
		break
	default:
		return nil, errors.New("can not find the user")
	}

	return createFlowOutput("user_info", userInfo)
}

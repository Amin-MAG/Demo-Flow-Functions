package function

import (
	"errors"
	"time"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	lastRide := Ride{
		PassengerID: *request.Args.UserID,
	}
	switch *request.Args.UserID {
	case 10:
		lastRide = Ride{
			Time: time.Now().Add(-10 * time.Minute),
			Origin: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
			Destination: Point{
				Lat: 20.20,
				Lon: 30.30,
			},
		}
		break
	case 20:
		lastRide = Ride{
			Time: time.Now().Add(-20 * time.Minute),
			Origin: Point{
				Lat: 20.20,
				Lon: 30.30,
			},
			Destination: Point{
				Lat: 30.30,
				Lon: 20.20,
			},
		}
		break
	case 30:
		lastRide = Ride{
			Time: time.Now().Add(-30 * time.Minute),
			Origin: Point{
				Lat: 30.30,
				Lon: 20.20,
			},
			Destination: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
		}
		break
	case 40:
		lastRide = Ride{
			Time: time.Now().Add(-40 * time.Minute),
			Origin: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
			Destination: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
		}
		break
	default:
		return nil, errors.New("can not find the user")
	}

	return createFlowOutput("last_ride", lastRide)
}

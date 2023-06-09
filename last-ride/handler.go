package function

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

func ExecFlow(request FlowInput) ([]byte, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	// Add a sleep time for simulating database connection
	time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)

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

	return json.Marshal(lastRide)
}

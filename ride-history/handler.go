package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ExecFlow(request FlowInput) ([]byte, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	var lastRide Ride
	lastRideResponse, ok := request.Children["last_ride_of_passenger"]
	if !ok || lastRideResponse == nil {
		return nil, errors.New("response of last_ride_of_passenger is required to process")
	}
	err := json.Unmarshal(lastRideResponse.Data, &lastRide)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", lastRide)

	return json.Marshal(RideHistory{Rides: []RideSummary{
		{
			Time:        lastRide.Time,
			Destination: lastRide.Destination,
		},
		{
			Time:        lastRide.Time,
			Destination: lastRide.Origin,
		},
		{
			Time:        lastRide.Time,
			Destination: lastRide.Destination,
		},
	}})
}

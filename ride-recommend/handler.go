package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}
	if request.Args.Origin == nil {
		return nil, errors.New("origin is required")
	}

	var lastRide Ride
	lastRideResponse, ok := request.Children["last_ride"]
	if !ok {
		return nil, errors.New("response of last_ride is required to process")
	}
	if lastRideResponse == nil {
		return nil, errors.New("the response of last_ride is nil")
	}
	jsonStr, err := json.Marshal(lastRideResponse.Data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in marshalling: %s", err))
	}
	err = json.Unmarshal(jsonStr, &lastRide)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", lastRide)

	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info"]
	if !ok {
		return nil, errors.New("response of user_info is required to process")
	}
	if userInfoResponse == nil {
		return nil, errors.New("the response of user_info is nil")
	}
	jsonStr, err = json.Marshal(userInfoResponse.Data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in marshalling: %s", err))
	}
	err = json.Unmarshal(jsonStr, &userInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", userInfo)

	var recommendation Recommendation
	// Repeat recommendation
	if request.Args.Origin.Lat == lastRide.Origin.Lat && request.Args.Origin.Lon == lastRide.Origin.Lon {
		recommendation = Recommendation{
			Type:           RecommendationRepeat,
			Recommendation: &lastRide.Destination,
			BannerText:     fmt.Sprintf("Dear %s, Here is your repeat recommendation.", userInfo.FirstName),
		}
		return createFlowOutput("last_ride", recommendation)
	}

	// Reverse recommendation
	if request.Args.Origin.Lat == lastRide.Destination.Lat && request.Args.Origin.Lon == lastRide.Destination.Lon {
		recommendation = Recommendation{
			Type:           RecommendationReverse,
			Recommendation: &lastRide.Origin,
			BannerText:     fmt.Sprintf("Dear %s, Here is your reverse recommendation.", userInfo.FirstName),
		}
		return createFlowOutput("last_ride", recommendation)
	}

	// No recommendation
	recommendation = Recommendation{
		Type:       RecommendationNothing,
		BannerText: fmt.Sprintf("Dear %s, There is no recommendation.", userInfo.FirstName),
	}

	return createFlowOutput("ride_recommend", recommendation)
}

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
	if request.Args.Origin == nil {
		return nil, errors.New("origin is required")
	}

	var recommendation Recommendation
	recommendationResponse, ok := request.Children["ride_recommendation"]
	if !ok || recommendationResponse == nil {
		return nil, errors.New("response of ride_recommendation is required to process")
	}
	err := json.Unmarshal(recommendationResponse.Data, &recommendation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", recommendation)

	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info_of_passenger"]
	if !ok || userInfoResponse == nil {
		return nil, errors.New("response of user_info_of_passenger is required to process")
	}
	err = json.Unmarshal(userInfoResponse.Data, &userInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", userInfo)

	var homepageDetails HomepageDetails
	if recommendation.Type != RecommendationNothing {
		homepageDetails = HomepageDetails{
			IsAnythingRecommended:    true,
			RecommendationBannerText: &recommendation.BannerText,
			RecommendationType:       &recommendation.Type,
		}
	}
	homepageDetails.UserCurrentLocation = userInfo.CurrentAddressLocation
	homepageDetails.UserAddresses = userInfo.Addresses
	fmt.Printf("%+v\n", homepageDetails)

	return json.Marshal(homepageDetails)
}

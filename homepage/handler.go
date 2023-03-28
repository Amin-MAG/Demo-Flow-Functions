package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	// Read environment variables
	isFlowCachingOn := os.Getenv("FLOW_CACHING")
	if isFlowCachingOn == "" {
		isFlowCachingOn = "false"
	}

	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}
	if request.Args.Origin == nil {
		return nil, errors.New("origin is required")
	}

	var recommendation Recommendation
	recommendationResponse, ok := request.Children["ride_recommendation"]
	if !ok {
		return nil, errors.New("response of ride_recommendation is required to process")
	}
	if recommendationResponse == nil {
		return nil, errors.New("the response of ride_recommendation is nil")
	}
	jsonStr, err := json.Marshal(recommendationResponse.Data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in marshalling: %s", err))
	}
	err = json.Unmarshal(jsonStr, &recommendation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", recommendation)

	var userInfo UserInfo
	if isFlowCachingOn == "false" {
		userInfoResponse, ok := request.Children["user_info_of_passenger"]
		if !ok {
			return nil, errors.New("response of user_info_of_passenger is required to process")
		}
		if userInfoResponse == nil {
			return nil, errors.New("the response of user_info_of_passenger is nil")
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
	} else {
		fmt.Println("Use caching...")
	}

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

	return createFlowOutput("homepage", homepageDetails)
}

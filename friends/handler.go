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

	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info_of_passenger"]
	if !ok || userInfoResponse == nil {
		return nil, errors.New("response of user_info_of_passenger is required to process")
	}
	err := json.Unmarshal(userInfoResponse.Data, &userInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", userInfo)

	return json.Marshal(FriendsInfo{
		NumberOfFriends: userInfo.PhoneNumber,
	})
}

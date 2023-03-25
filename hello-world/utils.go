package function

import (
	"encoding/json"
	"errors"
)

func createFlowOutput(fName string, data any) (*FlowOutput, error) {
	mapData, err := structToMap(data)
	if err != nil {
		return nil, errors.New("can not convert struct to map")
	}

	return &FlowOutput{
		Data:     mapData,
		Function: "last_ride",
	}, nil
}

func structToMap(s any) (map[string]interface{}, error) {
	var myMap map[string]interface{}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &myMap)
	if err != nil {
		return nil, err
	}

	return myMap, nil
}

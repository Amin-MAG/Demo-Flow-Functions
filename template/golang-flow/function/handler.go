package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ExecFlow(request FlowInput) ([]byte, error) {
	if request.Args.Name == nil {
		return nil, errors.New("name is required")
	}

	message := Message{Text: fmt.Sprintf("Hello %s!", *request.Args.Name)}

	return json.Marshal(message)
}

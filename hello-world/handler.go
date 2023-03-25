package function

import (
	"errors"
	"fmt"
)

func ExecFlow(request FlowInput) (*FlowOutput, error) {
	if request.Args.Name == nil {
		return nil, errors.New("name is required")
	}

	message := Message{Text: fmt.Sprintf("Hello %s!", *request.Args.Name)}

	return createFlowOutput("hello_world", message)
}

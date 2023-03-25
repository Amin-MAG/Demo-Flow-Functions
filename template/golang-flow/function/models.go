package function

type Message struct {
	Text string `json:"text"`
}

// Input is the argument of your flow function
type Input struct {
	Name *string `json:"name"`
}

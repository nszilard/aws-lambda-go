package internal

// Event maps the received payload
type Event struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

package models

// Event maps the received payload
type Event struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Response maps the response
type Response struct {
	Value int `json:"value"`
}

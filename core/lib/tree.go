package lib

type Tree struct {
	Name   string  `json:"name,omitempty"`
	Alias  string  `json:"alias,omitempty"`
	Routes []Route `json:"routes,omitempty"`
}

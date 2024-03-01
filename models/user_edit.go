package models

type UserEdit struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Title    string `json:"title,omitempty"`
}

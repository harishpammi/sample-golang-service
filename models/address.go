package models

//Address : Model for address info
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

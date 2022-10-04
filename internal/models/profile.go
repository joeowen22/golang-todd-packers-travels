package models

type Profile struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Travels   []Travel `json:"travelled"`
}

package models

type Pokemon struct {
	Id        int            `json:"Id"`
	Name      string         `json:"Name"`
	Power     string         `json:"Type"`
	Abilities map[string]int `json:"Abilities"`
}

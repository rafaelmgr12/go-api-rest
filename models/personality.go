package models

type Personality struct {
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Personalities []Personality

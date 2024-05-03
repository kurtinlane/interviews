package starwars

import "time"

type Result struct {
	Count    int         `json:"count"`
	Next     any         `json:"next"`
	Previous any         `json:"previous"`
	Results  []Character `json:"results"`
}

// Character represents a Star Wars character.
type Character struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

type Vehicle struct {
	Name                 string    `json:"name"`
	Model                string    `json:"model"`
	Manufacturer         string    `json:"manufacturer"`
	CostInCredits        string    `json:"cost_in_credits"`
	Length               string    `json:"length"`
	MaxAtmospheringSpeed string    `json:"max_atmosphering_speed"`
	Crew                 string    `json:"crew"`
	Passengers           string    `json:"passengers"`
	CargoCapacity        string    `json:"cargo_capacity"`
	Consumables          string    `json:"consumables"`
	VehicleClass         string    `json:"vehicle_class"`
	Pilots               []string  `json:"pilots"`
	Films                []string  `json:"films"`
	Created              time.Time `json:"created"`
	Edited               time.Time `json:"edited"`
	URL                  string    `json:"url"`
}

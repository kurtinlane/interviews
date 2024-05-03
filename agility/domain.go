package main

// Character represents a Star Wars character.
type Character struct {
	Name      string     `json:"name,omitempty"`
	Height    string     `json:"height,omitempty"`
	Mass      string     `json:"mass,omitempty"`
	HairColor string     `json:"hair_color,omitempty"`
	SkinColor string     `json:"skin_color,omitempty"`
	EyeColor  string     `json:"eye_color,omitempty"`
	BirthYear string     `json:"birth_year,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	Homeworld Homeworld  `json:"homeworld,omitempty"`
	Films     []string   `json:"films,omitempty"`
	Species   []Species  `json:"species,omitempty"`
	Vehicles  []Vehicle  `json:"vehicles,omitempty"`
	Starships []Starship `json:"starships,omitempty"`
	URL       string     `json:"url,omitempty"`
}

type Vehicle struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	VehicleClass         string   `json:"vehicle_class,omitempty"`
	Pilots               []string `json:"pilots,omitempty"`
	Films                []string `json:"films,omitempty"`
	URL                  string   `json:"url,omitempty"`
}

type Starship struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	HyperdriveRating     string   `json:"hyperdrive_rating,omitempty"`
	Mglt                 string   `json:"MGLT,omitempty"`
	StarshipClass        string   `json:"starship_class,omitempty"`
	Pilots               []string `json:"pilots,omitempty"`
	Films                []string `json:"films,omitempty"`
	URL                  string   `json:"url,omitempty"`
}

type Homeworld struct {
	Name           string   `json:"name,omitempty"`
	RotationPeriod string   `json:"rotation_period,omitempty"`
	OrbitalPeriod  string   `json:"orbital_period,omitempty"`
	Diameter       string   `json:"diameter,omitempty"`
	Climate        string   `json:"climate,omitempty"`
	Gravity        string   `json:"gravity,omitempty"`
	Terrain        string   `json:"terrain,omitempty"`
	SurfaceWater   string   `json:"surface_water,omitempty"`
	Population     string   `json:"population,omitempty"`
	Residents      []string `json:"residents,omitempty"`
	Films          []any    `json:"films,omitempty"`
	URL            string   `json:"url,omitempty"`
}

type Species struct {
	Name            string   `json:"name,omitempty"`
	Classification  string   `json:"classification,omitempty"`
	Designation     string   `json:"designation,omitempty"`
	AverageHeight   string   `json:"average_height,omitempty"`
	SkinColors      string   `json:"skin_colors,omitempty"`
	HairColors      string   `json:"hair_colors,omitempty"`
	EyeColors       string   `json:"eye_colors,omitempty"`
	AverageLifespan string   `json:"average_lifespan,omitempty"`
	Homeworld       string   `json:"homeworld,omitempty"`
	Language        string   `json:"language,omitempty"`
	People          []string `json:"people,omitempty"`
	Films           []string `json:"films,omitempty"`
	URL             string   `json:"url,omitempty"`
}

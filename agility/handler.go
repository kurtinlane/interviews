package main

import (
	"agility/starwars"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func characterHandler(w http.ResponseWriter, r *http.Request) {
	// Extract character name from URL path
	name := strings.TrimPrefix(r.URL.Path, "/character/")

	//QueryUnescape name from URL
	name, err := url.QueryUnescape(name)
	if err != nil {
		http.Error(w, "Error unescaping URL", http.StatusInternalServerError)
		return
	}

	// Fetch character
	character, err := getCharacter(name)
	if err != nil {
		http.Error(w, "Error fetching character data", http.StatusInternalServerError)
		return
	}

	// Marshal character data to JSON
	jsonData, err := json.Marshal(character)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getCharacter(name string) ([]Character, error) {
	// Fetch character data from Star Wars API
	resp, err := http.Get(fmt.Sprintf("https://swapi.dev/api/people/?search=%s", url.QueryEscape(name)))
	if err != nil {
		fmt.Println(err, "error in httpGet")
		return nil, err
	}
	defer resp.Body.Close()

	result := starwars.Result{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check if character was found
	if len(result.Results) == 0 {
		return nil, fmt.Errorf("character not found")
	}

	var characters []Character

	//range over the results and populate the vehicles and starships
	for _, result := range result.Results {
		responseCharacter := Character{
			Name: result.Name,
		}
		//add responseCharacter to characters slice
		characters = append(characters, responseCharacter)

		//Check if the character has species and hydrate it
		if len(result.Species) > 0 {
			var species []Species
			for _, speciesURL := range result.Species {
				s, err := getSpecies(speciesURL)
				if err != nil {
					return nil, fmt.Errorf("error getting species data")
				}
				a := Species{
					Name:            s.Name,
					Language:        s.Language,
					AverageLifespan: s.AverageLifespan,
				}
				species = append(species, a)
			}
			characters[len(characters)-1].Species = species
		}
		if len(result.Starships) > 0 {
			var starships []Starship
			for _, starshipURL := range result.Starships {
				s, err := getStarship(starshipURL)
				if err != nil {
					return nil, fmt.Errorf("error getting s data")
				}
				a := Starship{
					Name:          s.Name,
					CargoCapacity: s.CargoCapacity,
					StarshipClass: s.StarshipClass,
				}
				starships = append(starships, a)
			}
			characters[len(characters)-1].Starships = starships
		}

		if len(result.Homeworld) > 0 {
			s, err := getHomeworld(result.Homeworld)
			if err != nil {
				return nil, fmt.Errorf("error getting homeworld data")
			}
			h := Homeworld{
				Name:       s.Name,
				Population: s.Population,
				Climate:    s.Climate,
			}
			characters[len(characters)-1].Homeworld = h
		}

	}
	return characters, nil
}

func getSpecies(url string) (Species, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "error in vehicle httpGet")
		return Species{}, err
	}
	defer resp.Body.Close()

	v := Species{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		fmt.Println(err)
		return Species{}, err
	}

	return v, nil
}

func getStarship(url string) (Starship, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "error in Starship httpGet")
		return Starship{}, err
	}
	defer resp.Body.Close()

	v := Starship{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		fmt.Println(err)
		return Starship{}, err
	}

	return v, nil
}

func getHomeworld(url string) (Homeworld, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "error in Homeworld httpGet")
		return Homeworld{}, err
	}
	defer resp.Body.Close()

	v := Homeworld{}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		fmt.Println(err)
		return Homeworld{}, err
	}

	return v, nil
}

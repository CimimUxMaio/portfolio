package model

import (
	"encoding/json"
	"os"
)

type Profile struct {
	Picture string `json:"picture"`
	// Summary string `json:"summary"`
}

type Contact struct {
	Email    string `json:"email"`
	GitHub   string `json:"gitHub"`
	LinkedIn string `json:"linkedIn"`
}

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	Repository   string   `json:"repository"`
	Image        string   `json:"image"`
}

type Portfolio struct {
	Profile  Profile   `json:"profile"`
	Contact  Contact   `json:"contact"`
	Projects []Project `json:"projects"`
}

func LoadPortfolio() (Portfolio, error) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		return Portfolio{}, err
	}

	data := Portfolio{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return Portfolio{}, err
	}

	return data, nil
}

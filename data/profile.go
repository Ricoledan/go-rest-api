package data

import "time"

type Technique struct {
	Name        string
	Description string
}

type Profile struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Level      string      `json:"level"`
	Type       string      `json:"type"`
	Attribute  string      `json:"attribute"`
	Fields     []string    `json:"fields"`
	Group      []string    `json:"group"`
	Techniques []Technique `json:"techniques"`
	Artwork    string      `json:artwork`
	Profile    string      `json:profile`
	CreatedAt  string      `json:createdAt`
	UpdatedAt  string      `json:updatedAt`
	DeletedAt  string      `json:deletedAt`
}

var profileList = []*Profile{
	&Profile{
		ID:        1,
		Name:      "Agumon",
		Level:     "Child",
		Type:      "Reptile",
		Attribute: "Vaccine",
		Fields:    ["Metal Empire", "Nature Spirits", "Virus Busters", "Unknown", "Dragon's Roar", "Nightmare Soldiers"],
		Group: nil,
		Techniques: []Technique{
			Name:        "Baby Flame",
			Description: "Releases a stream of fire from its mouth.",
		},
		Artwork:   "https://wikimon.net/images/7/72/Agumon.jpg",
		Profile:   "A Reptile Digimon with an appearance resembling a small dinosaur, it has grown and become able to walk on two legs. Its strength is weak as it is still in the process of growing, but it has a fearless and rather ferocious personality. Hard, sharp claws grow from both its hands and feet, and their power is displayed in battle. It also foreshadows an evolution into a great and powerful Digimon. Its Special Move is spitting a fiery breath from its mouth to attack the opponent",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Profile{
		ID:        2,
		Name:      "Agumon",
		Level:     "Child",
		Type:      "Reptile",
		Attribute: "Vaccine",
		Fields:    ["Metal Empire", "Nature Spirits", "Virus Busters", "Unknown", "Dragon's Roar", "Nightmare Soldiers"],
		Group: nil,
		Techniques: []Technique{
			Name:        "Baby Flame",
			Description: "Releases a stream of fire from its mouth.",
		},
		Artwork:   "https://wikimon.net/images/7/72/Agumon.jpg",
		Profile:   "A Reptile Digimon with an appearance resembling a small dinosaur, it has grown and become able to walk on two legs. Its strength is weak as it is still in the process of growing, but it has a fearless and rather ferocious personality. Hard, sharp claws grow from both its hands and feet, and their power is displayed in battle. It also foreshadows an evolution into a great and powerful Digimon. Its Special Move is spitting a fiery breath from its mouth to attack the opponent",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}

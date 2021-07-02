package data

import "time"

type Technique struct {
	Name        string
	Description string
}

type Profile struct {
	ID        int
	Name      string
	Level     string
	Type      string
	Attribute string
	Field     [string]
	Group     []string
	Technique []Technique
	Artwork   string
	Profile   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

var profileList = []*Profile{
	&Profile{
		ID:        1,
		Name:      "Agumon",
		Level:     "Child",
		Type:      "Reptile",
		Attribute: "Vaccine",
		Field:     "Metal Empire", "Nature Spirits", "Virus Busters", "Unknown", "Dragon's Roar", "Nightmare Soldiers",
		Group: nil,
		Technique: []*Technique{
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
		Field:     "Metal Empire", "Nature Spirits", "Virus Busters", "Unknown", "Dragon's Roar", "Nightmare Soldiers",
		Group: nil,
		Technique: []*Technique{
			Name:        "Baby Flame",
			Description: "Releases a stream of fire from its mouth.",
		},
		Artwork:   "https://wikimon.net/images/7/72/Agumon.jpg",
		Profile:   "A Reptile Digimon with an appearance resembling a small dinosaur, it has grown and become able to walk on two legs. Its strength is weak as it is still in the process of growing, but it has a fearless and rather ferocious personality. Hard, sharp claws grow from both its hands and feet, and their power is displayed in battle. It also foreshadows an evolution into a great and powerful Digimon. Its Special Move is spitting a fiery breath from its mouth to attack the opponent",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}

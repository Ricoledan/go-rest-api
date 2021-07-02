package data

type Technique struct {
	Name string
	Description string
}

type Profile struct {
	ID        int
	Name      string
	Level     string
	Type      string
	Attribute string
	Field     []string
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
		ID: 1,
		Name: "Agumon",
		Level: "Child",
		Type: "Reptile",
		Attribute: "Vaccine",
		Field: ["Metal Empire","Nature Spirits","Virus Busters","Unknown","Dragon's Roar","Nightmare Soldiers"]
	}
}

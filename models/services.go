package models

type Service struct {
	ID          string   `json:"id"`
	Member      string   `json:"member"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Duration    string   `json:"duration"`
	Concessions bool     `json:"concessions"`
	Price       float64  `json:"price"`
	Regions     []string `json:"regions"`
	Ailments    []string `json:"ailments"`
}

package models

type Service struct {
	ID          string   `json:"id" firestore:"id"`
	Member      string   `json:"member" firestore:"member"`
	Name        string   `json:"name" firestore:"name"`
	Image       string   `json:"image" firestore:"image"`
	Description string   `json:"description" firestore:"description"`
	Duration    string   `json:"duration" firestore:"duration"`
	Concessions bool     `json:"concessions" firestore:"concessions"`
	Price       float64  `json:"price" firestore:"price"`
	Regions     []string `json:"regions" firestore:"regions"`
	Ailments    []string `json:"ailments" firestore:"ailments"`
}

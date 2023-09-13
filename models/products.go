package models

import "github.com/google/uuid"

type Product struct {
	Meta        Internals
	ID          string  `json:"id" firestore:"id"`
	Name        string  `json:"name" firestore:"name"`
	Description string  `json:"description" firestore:"description"`
	Price       float64 `json:"price" firestore:"price"`
	Length      string  `json:"length" firestore:"length"`
	Width       string  `json:"width" firestore:"width"`
	Height      string  `json:"height" firestore:"height"`
	Weight      string  `json:"weight" firestore:"weight"`
}

func (user *User) NewProduct() *Product {
	return &Product{
		Meta: NewInternals(),
		ID:   uuid.NewString(),
	}
}

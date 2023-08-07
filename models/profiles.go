package models

type Profiles struct {
	Social       SocialProfile       `json:"social" firestore:"social"`
	Personal     PersonalProfile     `json:"personal" firestore:"personal"`
	Practitioner PractitionerProfile `json:"practitioner" firestore:"practitioner"`
	Business     BusinessProfile     `json:"business" firestore:"business"`
}

type PersonalProfile struct {
	Internals
	Firstname string `json:"firstname" firestore:"firstname"`
	Lastname  string `json:"lastname" firestore:"lastname"`
	Phone     string `json:"phone" firestore:"phone"`
}

type PractitionerProfile struct {
	Internals
	Description string   `json:"description" firestore:"description"`
	Phone       string   `json:"phone" firestore:"phone"`
	Website     string   `json:"website" firestore:"website"`
	Keywords    []string `json:"keywords" firestore:"keywords"`
}

type BusinessProfile struct {
	Internals
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	Address     string `json:"address" firestore:"address"`
	Phone       string `json:"phone" firestore:"phone"`
	Website     string `json:"website" firestore:"website"`
	VAT         string `json:"vat" firestore:"vat"`
}

type SocialProfile struct {
	Internals
	Facebook    string `json:"facebook" firestore:"facebook"`
	Telegram    string `json:"telegram" firestore:"telegram"`
	WhatsApp    string `json:"whatsapp" firestore:"whatsapp"`
	Instagram   string `json:"instagram" firestore:"instagram"`
	Snapchat    string `json:"snapchat" firestore:"snapchat"`
	Meta        string `json:"meta" firestore:"meta"`
	MetaThreads string `json:"metathreads" firestore:"metathreads"`
	Twitter     string `json:"twitter" firestore:"twitter"`
	What3Words  string `json:"w3w" firestore:"w3w"`
}

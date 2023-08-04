package models

type Profiles struct {
	Social       SocialProfile       `json:"social" firestore:"social"`
	Practitioner PractitionerProfile `json:"practitioner" firestore:"practitioner"`
	Business     BusinessProfile     `json:"business" firestore:"business"`
}

type PractitionerProfile struct {
	Internals
	Image       string   `json:"image" firestore:"image"`
	Firstname   string   `json:"firstname" firestore:"firstname"`
	Lastname    string   `json:"lastname" firestore:"lastname"`
	Description string   `json:"description" firestore:"description"`
	Keywords    []string `json:"keywords" firestore:"keywords"`
}

type BusinessProfile struct {
	Internals
	Name    string `json:"name" firestore:"name"`
	Address string `json:"address" firestore:"address"`
	VAT     string `json:"vat" firestore:"vat"`
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

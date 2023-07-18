package models

type Socials struct {
	Facebook    string `json:"facebook" firestore:"facebook"`
	Telegram    string `json:"telegram" firestore:"telegram"`
	WhatsApp    string `json:"whatsapp" firestore:"whatsapp"`
	Instagram   string `json:"instagram" firestore:"instagram"`
	Snapchat    string `json:"snapchat" firestore:"snapchat"`
	Meta        string `json:"meta" firestore:"meta"`
	MetaThreads string `json:"metathreads" firestore:"metathreads"`
	Twitter     string `json:"twitter" firestore:"twitter"`
}

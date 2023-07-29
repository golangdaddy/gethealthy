package models

type Mail struct {
	Recipient string `json:"recipient" firestore:"recipient"`
	Sender    string `json:"sender" firestore:"sender"`
	Subject   string `json:"subject" firestore:"subject"`
	Body      string `json:"body" firestore:"body"`
}

package models

type MediaFile struct {
	Meta        Internals
	Folder      string `json:"folder" firestore:"folder"`
	Kind        string `json:"kind" firestore:"kind"`
	URI         string `json:"uri" firestore:"uri"`
	Owner       string `json:"owner" firestore:"owner"`
	Description string `json:"description" firestore:"description"`
}

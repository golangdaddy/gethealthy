package models

type MediaFolder struct {
	Meta  Internals
	Owner string `json:"owner" firestore:"owner"`
	Name  string `json:"name" firestore:"name"`
}

func (user *User) NewMediaFolder(name string) *MediaFolder {
	return &MediaFolder{
		Meta:  NewInternals(),
		Owner: user.Username,
		Name:  name,
	}
}

type MediaFile struct {
	Meta        Internals
	Folder      string `json:"folder" firestore:"folder"`
	Kind        string `json:"kind" firestore:"kind"`
	URI         string `json:"uri" firestore:"uri"`
	Description string `json:"description" firestore:"description"`
}

func (folder *MediaFolder) NewMediaFile(kind, uri, description string) *MediaFile {
	return &MediaFile{
		Meta:        NewInternals(),
		Folder:      folder.Name,
		Kind:        kind,
		URI:         uri,
		Description: description,
	}
}

package model

type User struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Bucket struct {
	BucketID string `json:"bucketId"`
	Name     string `json:"name"`
}

type Data11 struct {
	Code       string `json:"code"`
	Uploadedby string `json:"uploaded_by"`
	Filename   string `json:"file"`
}

type Data12 struct {
	Code         string `json:"code"`
	Downloadedby string `json:"downloaded_by"`
}

type Share struct {
	DocID       string   `json:"documentId"`
	Data1       Data11   `json:"data"`
	Permissions []string `json:"permissions"`
}

type Download struct {
	DocID       string   `json:"documentId"`
	Data1       Data12   `json:"data"`
	Permissions []string `json:"permissions"`
}

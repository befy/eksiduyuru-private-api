package models

type PostPreview struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ID       string `json:"id"`
	Type     int    `json:"type"`
}

type PostType struct {
	Type int `json:"type"`
}

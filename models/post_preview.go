package models

type PostPreview struct {
	Title    string
	Subtitle string
	ID       string
	Type     PostType
}

type PostType struct {
	Type int
}

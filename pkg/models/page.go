package models

type Page struct {
	PageId    uint   `json:"pageid"`
	Namespace uint8  `json:"ns"`
	Title     string `json:"title"`
	Summary   string `json:"extract"`
	Links     []Page `json:"links"`
}

func NewPage(title string) Page {
	return Page{Title: title}
}

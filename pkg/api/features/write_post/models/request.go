package models

type Request struct {
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
	RootDomain  string `json:"rootDomain"`
}

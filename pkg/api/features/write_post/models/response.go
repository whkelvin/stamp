package models

type Response struct {
	PostId      string `json:"postId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	CreatedDate string `json:"createdDate"`
	RootDomain  string `json:"rootDomain"`
}

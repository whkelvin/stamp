package models

type Request struct {
	Link        string `bson:"link"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	RootDomain  string `bson:"rootDomain"`
	CreatedDate string `bson:"createdDate"`
}

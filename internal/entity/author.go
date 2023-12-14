package entity

type Author struct {
	Username string `json:"username" bson:"username"`
	ID       string `json:"id" bson:"id"`
}

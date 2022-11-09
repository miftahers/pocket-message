package dto

type PocketMessageWithRandomID struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	RandomID string `json:"random_id"`
}

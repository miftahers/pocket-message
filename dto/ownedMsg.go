package dto

type OwnedMessage struct {
	RandomID string `json:"random_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Visit    int    `json:"visit"`
}

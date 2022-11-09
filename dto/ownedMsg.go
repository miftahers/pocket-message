package dto

import "github.com/google/uuid"

type OwnedMessage struct {
	UUID     uuid.UUID `json:"uuid"`
	RandomID string    `json:"random_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Visit    int       `json:"visit"`
}

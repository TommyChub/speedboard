package level

import "github.com/google/uuid"

type Level struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	GameId      string    `json:"submittedTime"`
}

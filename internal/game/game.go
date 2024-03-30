package level

import (
	"time"

	"github.com/google/uuid"
)

type Level struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ReleasedAt  string    `json:"submittedTime"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

package submission

import (
	"time"

	"github.com/google/uuid"
)

type Submission struct {
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	LevelId       uint      `json:"levelId"`
	SubmittedTime time.Time `json:"submittedTime"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

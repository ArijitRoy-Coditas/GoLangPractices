package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Tasks struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    int       `json:"priority"`
	CreatedAT   time.Time `json:"created_at"`
	UpdatedAT   time.Time `json:"updated_at"`
}

var Task []Tasks

func (t Tasks) MarshalJson() ([]byte, error) {
	type Alias Tasks

	return json.Marshal(&struct {
		Alias
		CreatedAT string `json:"created_at"`
		UpdatedAT string `json:"updated_at"`
	}{
		Alias:     Alias(t),
		CreatedAT: t.CreatedAT.Format("2006-01-02 T15:04:05"),
		UpdatedAT: t.UpdatedAT.Format("2006-01-02 T15:04:05"),
	})
}

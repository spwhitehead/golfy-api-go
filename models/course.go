package models

import (
	"github.com/google/uuid"
)

// Course represents the course model for the API
type Course struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Location string    `json:"location"`
    NumHoles int       `json:"num_holes"`
}

package model

import "github.com/google/uuid"

// StationWithDistanceUnstructured is the model for the metarWithDistance type.
type StationWithDistanceUnstructured struct {
	ID       uuid.UUID `json:"id"`
	Distance float64   `json:"distance"`
	Stations int       `json:"airport_stations"`
}

package model

// StationWithDistanceUnstructured is the model for the metarWithDistance type.
type StationWithDistanceUnstructured struct {
	ID       int     `json:"id"`
	Distance float64 `json:"distance"`
	Stations int     `json:"airport_stations"`
}

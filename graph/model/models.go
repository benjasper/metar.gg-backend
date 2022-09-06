package model

// MetarWithDistanceUnstructured is the model for the metarWithDistance type.
type MetarWithDistanceUnstructured struct {
	ID            int     `json:"id"`
	Distance      float64 `json:"distance"`
	AirportMetars int     `json:"airport_metars"`
}

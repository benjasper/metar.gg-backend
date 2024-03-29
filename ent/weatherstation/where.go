// Code generated by ent, DO NOT EDIT.

package weatherstation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"metar.gg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldID, id))
}

// StationID applies equality check predicate on the "station_id" field. It's identical to StationIDEQ.
func StationID(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldStationID, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldLongitude, v))
}

// Elevation applies equality check predicate on the "elevation" field. It's identical to ElevationEQ.
func Elevation(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldElevation, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldHash, v))
}

// StationIDEQ applies the EQ predicate on the "station_id" field.
func StationIDEQ(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldStationID, v))
}

// StationIDNEQ applies the NEQ predicate on the "station_id" field.
func StationIDNEQ(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldStationID, v))
}

// StationIDIn applies the In predicate on the "station_id" field.
func StationIDIn(vs ...string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldStationID, vs...))
}

// StationIDNotIn applies the NotIn predicate on the "station_id" field.
func StationIDNotIn(vs ...string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldStationID, vs...))
}

// StationIDGT applies the GT predicate on the "station_id" field.
func StationIDGT(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldStationID, v))
}

// StationIDGTE applies the GTE predicate on the "station_id" field.
func StationIDGTE(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldStationID, v))
}

// StationIDLT applies the LT predicate on the "station_id" field.
func StationIDLT(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldStationID, v))
}

// StationIDLTE applies the LTE predicate on the "station_id" field.
func StationIDLTE(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldStationID, v))
}

// StationIDContains applies the Contains predicate on the "station_id" field.
func StationIDContains(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldContains(FieldStationID, v))
}

// StationIDHasPrefix applies the HasPrefix predicate on the "station_id" field.
func StationIDHasPrefix(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldHasPrefix(FieldStationID, v))
}

// StationIDHasSuffix applies the HasSuffix predicate on the "station_id" field.
func StationIDHasSuffix(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldHasSuffix(FieldStationID, v))
}

// StationIDEqualFold applies the EqualFold predicate on the "station_id" field.
func StationIDEqualFold(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEqualFold(FieldStationID, v))
}

// StationIDContainsFold applies the ContainsFold predicate on the "station_id" field.
func StationIDContainsFold(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldContainsFold(FieldStationID, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotNull(FieldLatitude))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotNull(FieldLongitude))
}

// ElevationEQ applies the EQ predicate on the "elevation" field.
func ElevationEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldElevation, v))
}

// ElevationNEQ applies the NEQ predicate on the "elevation" field.
func ElevationNEQ(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldElevation, v))
}

// ElevationIn applies the In predicate on the "elevation" field.
func ElevationIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldElevation, vs...))
}

// ElevationNotIn applies the NotIn predicate on the "elevation" field.
func ElevationNotIn(vs ...float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldElevation, vs...))
}

// ElevationGT applies the GT predicate on the "elevation" field.
func ElevationGT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldElevation, v))
}

// ElevationGTE applies the GTE predicate on the "elevation" field.
func ElevationGTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldElevation, v))
}

// ElevationLT applies the LT predicate on the "elevation" field.
func ElevationLT(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldElevation, v))
}

// ElevationLTE applies the LTE predicate on the "elevation" field.
func ElevationLTE(v float64) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldElevation, v))
}

// ElevationIsNil applies the IsNil predicate on the "elevation" field.
func ElevationIsNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIsNull(FieldElevation))
}

// ElevationNotNil applies the NotNil predicate on the "elevation" field.
func ElevationNotNil() predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotNull(FieldElevation))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.WeatherStation {
	return predicate.WeatherStation(sql.FieldContainsFold(FieldHash, v))
}

// HasAirport applies the HasEdge predicate on the "airport" edge.
func HasAirport() predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, AirportTable, AirportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAirportWith applies the HasEdge predicate on the "airport" edge with a given conditions (other predicates).
func HasAirportWith(preds ...predicate.Airport) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := newAirportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetars applies the HasEdge predicate on the "metars" edge.
func HasMetars() predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetarsTable, MetarsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetarsWith applies the HasEdge predicate on the "metars" edge with a given conditions (other predicates).
func HasMetarsWith(preds ...predicate.Metar) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := newMetarsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTafs applies the HasEdge predicate on the "tafs" edge.
func HasTafs() predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TafsTable, TafsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTafsWith applies the HasEdge predicate on the "tafs" edge with a given conditions (other predicates).
func HasTafsWith(preds ...predicate.Taf) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		step := newTafsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WeatherStation) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WeatherStation) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WeatherStation) predicate.WeatherStation {
	return predicate.WeatherStation(func(s *sql.Selector) {
		p(s.Not())
	})
}

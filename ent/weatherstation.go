// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/airport"
	"metar.gg/ent/weatherstation"
)

// WeatherStation is the model entity for the WeatherStation schema.
type WeatherStation struct {
	config `json:"-"`
	// ID of the ent.
	// The unique identifier of the record.
	ID uuid.UUID `json:"id,omitempty"`
	// The ICAO identifier of the station that provided the weather data or identifier of the weather station.
	StationID string `json:"station_id,omitempty"`
	// The latitude in decimal degrees of the station.
	Latitude *float64 `json:"latitude,omitempty"`
	// The longitude in decimal degrees of the station.
	Longitude *float64 `json:"longitude,omitempty"`
	// The elevation in meters of the station.
	Elevation *float64 `json:"elevation,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WeatherStationQuery when eager-loading is set.
	Edges           WeatherStationEdges `json:"edges"`
	airport_station *uuid.UUID
	selectValues    sql.SelectValues
}

// WeatherStationEdges holds the relations/edges for other nodes in the graph.
type WeatherStationEdges struct {
	// The airport that hosts this station. This can also be empty if the metar is from a weather station outside an airport.
	Airport *Airport `json:"airport,omitempty"`
	// The metars that were reported by this station.
	Metars []*Metar `json:"metars,omitempty"`
	// The tafs that were reported by this station.
	Tafs []*Taf `json:"tafs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedMetars map[string][]*Metar
	namedTafs   map[string][]*Taf
}

// AirportOrErr returns the Airport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WeatherStationEdges) AirportOrErr() (*Airport, error) {
	if e.loadedTypes[0] {
		if e.Airport == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Airport, nil
	}
	return nil, &NotLoadedError{edge: "airport"}
}

// MetarsOrErr returns the Metars value or an error if the edge
// was not loaded in eager-loading.
func (e WeatherStationEdges) MetarsOrErr() ([]*Metar, error) {
	if e.loadedTypes[1] {
		return e.Metars, nil
	}
	return nil, &NotLoadedError{edge: "metars"}
}

// TafsOrErr returns the Tafs value or an error if the edge
// was not loaded in eager-loading.
func (e WeatherStationEdges) TafsOrErr() ([]*Taf, error) {
	if e.loadedTypes[2] {
		return e.Tafs, nil
	}
	return nil, &NotLoadedError{edge: "tafs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WeatherStation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case weatherstation.FieldLatitude, weatherstation.FieldLongitude, weatherstation.FieldElevation:
			values[i] = new(sql.NullFloat64)
		case weatherstation.FieldStationID, weatherstation.FieldHash:
			values[i] = new(sql.NullString)
		case weatherstation.FieldID:
			values[i] = new(uuid.UUID)
		case weatherstation.ForeignKeys[0]: // airport_station
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WeatherStation fields.
func (ws *WeatherStation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case weatherstation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ws.ID = *value
			}
		case weatherstation.FieldStationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field station_id", values[i])
			} else if value.Valid {
				ws.StationID = value.String
			}
		case weatherstation.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				ws.Latitude = new(float64)
				*ws.Latitude = value.Float64
			}
		case weatherstation.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				ws.Longitude = new(float64)
				*ws.Longitude = value.Float64
			}
		case weatherstation.FieldElevation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field elevation", values[i])
			} else if value.Valid {
				ws.Elevation = new(float64)
				*ws.Elevation = value.Float64
			}
		case weatherstation.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				ws.Hash = value.String
			}
		case weatherstation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airport_station", values[i])
			} else if value.Valid {
				ws.airport_station = new(uuid.UUID)
				*ws.airport_station = *value.S.(*uuid.UUID)
			}
		default:
			ws.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WeatherStation.
// This includes values selected through modifiers, order, etc.
func (ws *WeatherStation) Value(name string) (ent.Value, error) {
	return ws.selectValues.Get(name)
}

// QueryAirport queries the "airport" edge of the WeatherStation entity.
func (ws *WeatherStation) QueryAirport() *AirportQuery {
	return NewWeatherStationClient(ws.config).QueryAirport(ws)
}

// QueryMetars queries the "metars" edge of the WeatherStation entity.
func (ws *WeatherStation) QueryMetars() *MetarQuery {
	return NewWeatherStationClient(ws.config).QueryMetars(ws)
}

// QueryTafs queries the "tafs" edge of the WeatherStation entity.
func (ws *WeatherStation) QueryTafs() *TafQuery {
	return NewWeatherStationClient(ws.config).QueryTafs(ws)
}

// Update returns a builder for updating this WeatherStation.
// Note that you need to call WeatherStation.Unwrap() before calling this method if this WeatherStation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ws *WeatherStation) Update() *WeatherStationUpdateOne {
	return NewWeatherStationClient(ws.config).UpdateOne(ws)
}

// Unwrap unwraps the WeatherStation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ws *WeatherStation) Unwrap() *WeatherStation {
	_tx, ok := ws.config.driver.(*txDriver)
	if !ok {
		panic("ent: WeatherStation is not a transactional entity")
	}
	ws.config.driver = _tx.drv
	return ws
}

// String implements the fmt.Stringer.
func (ws *WeatherStation) String() string {
	var builder strings.Builder
	builder.WriteString("WeatherStation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ws.ID))
	builder.WriteString("station_id=")
	builder.WriteString(ws.StationID)
	builder.WriteString(", ")
	if v := ws.Latitude; v != nil {
		builder.WriteString("latitude=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ws.Longitude; v != nil {
		builder.WriteString("longitude=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ws.Elevation; v != nil {
		builder.WriteString("elevation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(ws.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// NamedMetars returns the Metars named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ws *WeatherStation) NamedMetars(name string) ([]*Metar, error) {
	if ws.Edges.namedMetars == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ws.Edges.namedMetars[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ws *WeatherStation) appendNamedMetars(name string, edges ...*Metar) {
	if ws.Edges.namedMetars == nil {
		ws.Edges.namedMetars = make(map[string][]*Metar)
	}
	if len(edges) == 0 {
		ws.Edges.namedMetars[name] = []*Metar{}
	} else {
		ws.Edges.namedMetars[name] = append(ws.Edges.namedMetars[name], edges...)
	}
}

// NamedTafs returns the Tafs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ws *WeatherStation) NamedTafs(name string) ([]*Taf, error) {
	if ws.Edges.namedTafs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ws.Edges.namedTafs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ws *WeatherStation) appendNamedTafs(name string, edges ...*Taf) {
	if ws.Edges.namedTafs == nil {
		ws.Edges.namedTafs = make(map[string][]*Taf)
	}
	if len(edges) == 0 {
		ws.Edges.namedTafs[name] = []*Taf{}
	} else {
		ws.Edges.namedTafs[name] = append(ws.Edges.namedTafs[name], edges...)
	}
}

// WeatherStations is a parsable slice of WeatherStation.
type WeatherStations []*WeatherStation

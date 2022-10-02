// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/airport"
	"metar.gg/ent/country"
	"metar.gg/ent/region"
	"metar.gg/ent/weatherstation"
)

// Airport is the model entity for the Airport schema.
type Airport struct {
	config `json:"-"`
	// ID of the ent.
	// The unique identifier of the record.
	ID uuid.UUID `json:"id,omitempty"`
	// The unique identifier of the import.
	ImportID int `json:"import_id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// ImportFlag holds the value of the "import_flag" field.
	ImportFlag bool `json:"import_flag,omitempty"`
	// The last time the record was updated/created.
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// The four-letter ICAO code of the airport.
	IcaoCode string `json:"icao_code,omitempty"`
	// The three-letter IATA code for the airport.
	IataCode *string `json:"iata_code,omitempty"`
	// This will be the ICAO code if available. Otherwise, it will be a local airport code (if no conflict), or if nothing else is available, an internally-generated code starting with the ISO2 country code, followed by a dash and a four-digit number.
	Identifier string `json:"identifier,omitempty"`
	// Type of airport.
	Type airport.Type `json:"type,omitempty"`
	// Importance of the airport.
	Importance int `json:"importance,omitempty"`
	// The official airport name, including "Airport", "Airstrip", etc.
	Name string `json:"name,omitempty"`
	// Latitude of the airport in decimal degrees (positive is north).
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude of the airport in decimal degrees (positive is east).
	Longitude float64 `json:"longitude,omitempty"`
	// The timezone of the airport.
	Timezone *string `json:"timezone,omitempty"`
	// Elevation of the airport, in feet.
	Elevation *int `json:"elevation,omitempty"`
	// The primary municipality that the airport serves (when available). Note that this is not necessarily the municipality where the airport is physically located.
	Municipality *string `json:"municipality,omitempty"`
	// Whether the airport has scheduled airline service.
	ScheduledService bool `json:"scheduled_service,omitempty"`
	// The code that an aviation GPS database (such as Jeppesen's or Garmin's) would normally use for the airport. This will always be the ICAO code if one exists. Note that, unlike the ident column, this is not guaranteed to be globally unique.
	GpsCode *string `json:"gps_code,omitempty"`
	// The local country code for the airport, if different from the gps_code and iata_code fields (used mainly for US airports).
	LocalCode *string `json:"local_code,omitempty"`
	// The URL of the airport's website.
	Website *string `json:"website,omitempty"`
	// The URL of the airport's Wikipedia page.
	Wikipedia *string `json:"wikipedia,omitempty"`
	// Extra keywords/phrases to assist with search. May include former names for the airport, alternate codes, names in other languages, nearby tourist destinations, etc.
	Keywords []string `json:"keywords,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AirportQuery when eager-loading is set.
	Edges            AirportEdges `json:"edges"`
	country_airports *uuid.UUID
	region_airports  *uuid.UUID
}

// AirportEdges holds the relations/edges for other nodes in the graph.
type AirportEdges struct {
	// The region that the airport is located in.
	Region *Region `json:"region,omitempty"`
	// The country that the airport is located in.
	Country *Country `json:"country,omitempty"`
	// Runways at the airport.
	Runways []*Runway `json:"runways,omitempty"`
	// Frequencies at the airport.
	Frequencies []*Frequency `json:"frequencies,omitempty"`
	// Weather station at the airport.
	Station *WeatherStation `json:"station,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedRunways     map[string][]*Runway
	namedFrequencies map[string][]*Frequency
}

// RegionOrErr returns the Region value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AirportEdges) RegionOrErr() (*Region, error) {
	if e.loadedTypes[0] {
		if e.Region == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: region.Label}
		}
		return e.Region, nil
	}
	return nil, &NotLoadedError{edge: "region"}
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AirportEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[1] {
		if e.Country == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// RunwaysOrErr returns the Runways value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) RunwaysOrErr() ([]*Runway, error) {
	if e.loadedTypes[2] {
		return e.Runways, nil
	}
	return nil, &NotLoadedError{edge: "runways"}
}

// FrequenciesOrErr returns the Frequencies value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) FrequenciesOrErr() ([]*Frequency, error) {
	if e.loadedTypes[3] {
		return e.Frequencies, nil
	}
	return nil, &NotLoadedError{edge: "frequencies"}
}

// StationOrErr returns the Station value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AirportEdges) StationOrErr() (*WeatherStation, error) {
	if e.loadedTypes[4] {
		if e.Station == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: weatherstation.Label}
		}
		return e.Station, nil
	}
	return nil, &NotLoadedError{edge: "station"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Airport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case airport.FieldKeywords:
			values[i] = new([]byte)
		case airport.FieldImportFlag, airport.FieldScheduledService:
			values[i] = new(sql.NullBool)
		case airport.FieldLatitude, airport.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case airport.FieldImportID, airport.FieldImportance, airport.FieldElevation:
			values[i] = new(sql.NullInt64)
		case airport.FieldHash, airport.FieldIcaoCode, airport.FieldIataCode, airport.FieldIdentifier, airport.FieldType, airport.FieldName, airport.FieldTimezone, airport.FieldMunicipality, airport.FieldGpsCode, airport.FieldLocalCode, airport.FieldWebsite, airport.FieldWikipedia:
			values[i] = new(sql.NullString)
		case airport.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case airport.FieldID:
			values[i] = new(uuid.UUID)
		case airport.ForeignKeys[0]: // country_airports
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case airport.ForeignKeys[1]: // region_airports
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Airport", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Airport fields.
func (a *Airport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airport.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case airport.FieldImportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field import_id", values[i])
			} else if value.Valid {
				a.ImportID = int(value.Int64)
			}
		case airport.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				a.Hash = value.String
			}
		case airport.FieldImportFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field import_flag", values[i])
			} else if value.Valid {
				a.ImportFlag = value.Bool
			}
		case airport.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated", values[i])
			} else if value.Valid {
				a.LastUpdated = value.Time
			}
		case airport.FieldIcaoCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icao_code", values[i])
			} else if value.Valid {
				a.IcaoCode = value.String
			}
		case airport.FieldIataCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iata_code", values[i])
			} else if value.Valid {
				a.IataCode = new(string)
				*a.IataCode = value.String
			}
		case airport.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				a.Identifier = value.String
			}
		case airport.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = airport.Type(value.String)
			}
		case airport.FieldImportance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field importance", values[i])
			} else if value.Valid {
				a.Importance = int(value.Int64)
			}
		case airport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case airport.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				a.Latitude = value.Float64
			}
		case airport.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				a.Longitude = value.Float64
			}
		case airport.FieldTimezone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone", values[i])
			} else if value.Valid {
				a.Timezone = new(string)
				*a.Timezone = value.String
			}
		case airport.FieldElevation:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field elevation", values[i])
			} else if value.Valid {
				a.Elevation = new(int)
				*a.Elevation = int(value.Int64)
			}
		case airport.FieldMunicipality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field municipality", values[i])
			} else if value.Valid {
				a.Municipality = new(string)
				*a.Municipality = value.String
			}
		case airport.FieldScheduledService:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field scheduled_service", values[i])
			} else if value.Valid {
				a.ScheduledService = value.Bool
			}
		case airport.FieldGpsCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gps_code", values[i])
			} else if value.Valid {
				a.GpsCode = new(string)
				*a.GpsCode = value.String
			}
		case airport.FieldLocalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_code", values[i])
			} else if value.Valid {
				a.LocalCode = new(string)
				*a.LocalCode = value.String
			}
		case airport.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				a.Website = new(string)
				*a.Website = value.String
			}
		case airport.FieldWikipedia:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wikipedia", values[i])
			} else if value.Valid {
				a.Wikipedia = new(string)
				*a.Wikipedia = value.String
			}
		case airport.FieldKeywords:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field keywords", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Keywords); err != nil {
					return fmt.Errorf("unmarshal field keywords: %w", err)
				}
			}
		case airport.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field country_airports", values[i])
			} else if value.Valid {
				a.country_airports = new(uuid.UUID)
				*a.country_airports = *value.S.(*uuid.UUID)
			}
		case airport.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field region_airports", values[i])
			} else if value.Valid {
				a.region_airports = new(uuid.UUID)
				*a.region_airports = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryRegion queries the "region" edge of the Airport entity.
func (a *Airport) QueryRegion() *RegionQuery {
	return (&AirportClient{config: a.config}).QueryRegion(a)
}

// QueryCountry queries the "country" edge of the Airport entity.
func (a *Airport) QueryCountry() *CountryQuery {
	return (&AirportClient{config: a.config}).QueryCountry(a)
}

// QueryRunways queries the "runways" edge of the Airport entity.
func (a *Airport) QueryRunways() *RunwayQuery {
	return (&AirportClient{config: a.config}).QueryRunways(a)
}

// QueryFrequencies queries the "frequencies" edge of the Airport entity.
func (a *Airport) QueryFrequencies() *FrequencyQuery {
	return (&AirportClient{config: a.config}).QueryFrequencies(a)
}

// QueryStation queries the "station" edge of the Airport entity.
func (a *Airport) QueryStation() *WeatherStationQuery {
	return (&AirportClient{config: a.config}).QueryStation(a)
}

// Update returns a builder for updating this Airport.
// Note that you need to call Airport.Unwrap() before calling this method if this Airport
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Airport) Update() *AirportUpdateOne {
	return (&AirportClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Airport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Airport) Unwrap() *Airport {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Airport is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Airport) String() string {
	var builder strings.Builder
	builder.WriteString("Airport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("import_id=")
	builder.WriteString(fmt.Sprintf("%v", a.ImportID))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(a.Hash)
	builder.WriteString(", ")
	builder.WriteString("import_flag=")
	builder.WriteString(fmt.Sprintf("%v", a.ImportFlag))
	builder.WriteString(", ")
	builder.WriteString("last_updated=")
	builder.WriteString(a.LastUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("icao_code=")
	builder.WriteString(a.IcaoCode)
	builder.WriteString(", ")
	if v := a.IataCode; v != nil {
		builder.WriteString("iata_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("identifier=")
	builder.WriteString(a.Identifier)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("importance=")
	builder.WriteString(fmt.Sprintf("%v", a.Importance))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", a.Latitude))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", a.Longitude))
	builder.WriteString(", ")
	if v := a.Timezone; v != nil {
		builder.WriteString("timezone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.Elevation; v != nil {
		builder.WriteString("elevation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.Municipality; v != nil {
		builder.WriteString("municipality=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("scheduled_service=")
	builder.WriteString(fmt.Sprintf("%v", a.ScheduledService))
	builder.WriteString(", ")
	if v := a.GpsCode; v != nil {
		builder.WriteString("gps_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.LocalCode; v != nil {
		builder.WriteString("local_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.Website; v != nil {
		builder.WriteString("website=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.Wikipedia; v != nil {
		builder.WriteString("wikipedia=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("keywords=")
	builder.WriteString(fmt.Sprintf("%v", a.Keywords))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRunways returns the Runways named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Airport) NamedRunways(name string) ([]*Runway, error) {
	if a.Edges.namedRunways == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedRunways[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Airport) appendNamedRunways(name string, edges ...*Runway) {
	if a.Edges.namedRunways == nil {
		a.Edges.namedRunways = make(map[string][]*Runway)
	}
	if len(edges) == 0 {
		a.Edges.namedRunways[name] = []*Runway{}
	} else {
		a.Edges.namedRunways[name] = append(a.Edges.namedRunways[name], edges...)
	}
}

// NamedFrequencies returns the Frequencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Airport) NamedFrequencies(name string) ([]*Frequency, error) {
	if a.Edges.namedFrequencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedFrequencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Airport) appendNamedFrequencies(name string, edges ...*Frequency) {
	if a.Edges.namedFrequencies == nil {
		a.Edges.namedFrequencies = make(map[string][]*Frequency)
	}
	if len(edges) == 0 {
		a.Edges.namedFrequencies[name] = []*Frequency{}
	} else {
		a.Edges.namedFrequencies[name] = append(a.Edges.namedFrequencies[name], edges...)
	}
}

// Airports is a parsable slice of Airport.
type Airports []*Airport

func (a Airports) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}

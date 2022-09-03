// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent/airport"
	"metar.gg/ent/frequency"
)

// Frequency is the model entity for the Frequency schema.
type Frequency struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// ImportFlag holds the value of the "import_flag" field.
	ImportFlag bool `json:"import_flag,omitempty"`
	// A code for the frequency type. Some common values are "TWR" (tower), "ATF" or "CTAF" (common traffic frequency), "GND" (ground control), "RMP" (ramp control), "ATIS" (automated weather), "RCO" (remote radio outlet), "ARR" (arrivals), "DEP" (departures), "UNICOM" (monitored ground station), and "RDO" (a flight-service station).
	Type string `json:"type,omitempty"`
	// A description of the frequency.
	Description string `json:"description,omitempty"`
	// Radio frequency in megahertz. Note that the same frequency may appear multiple times for an airport, serving different functions
	Frequency float64 `json:"frequency,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FrequencyQuery when eager-loading is set.
	Edges               FrequencyEdges `json:"edges"`
	airport_frequencies *int
}

// FrequencyEdges holds the relations/edges for other nodes in the graph.
type FrequencyEdges struct {
	// Airport holds the value of the airport edge.
	Airport *Airport `json:"airport,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// AirportOrErr returns the Airport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FrequencyEdges) AirportOrErr() (*Airport, error) {
	if e.loadedTypes[0] {
		if e.Airport == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Airport, nil
	}
	return nil, &NotLoadedError{edge: "airport"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Frequency) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case frequency.FieldImportFlag:
			values[i] = new(sql.NullBool)
		case frequency.FieldFrequency:
			values[i] = new(sql.NullFloat64)
		case frequency.FieldID:
			values[i] = new(sql.NullInt64)
		case frequency.FieldHash, frequency.FieldType, frequency.FieldDescription:
			values[i] = new(sql.NullString)
		case frequency.ForeignKeys[0]: // airport_frequencies
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Frequency", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Frequency fields.
func (f *Frequency) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case frequency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case frequency.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				f.Hash = value.String
			}
		case frequency.FieldImportFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field import_flag", values[i])
			} else if value.Valid {
				f.ImportFlag = value.Bool
			}
		case frequency.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = value.String
			}
		case frequency.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		case frequency.FieldFrequency:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field frequency", values[i])
			} else if value.Valid {
				f.Frequency = value.Float64
			}
		case frequency.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field airport_frequencies", value)
			} else if value.Valid {
				f.airport_frequencies = new(int)
				*f.airport_frequencies = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAirport queries the "airport" edge of the Frequency entity.
func (f *Frequency) QueryAirport() *AirportQuery {
	return (&FrequencyClient{config: f.config}).QueryAirport(f)
}

// Update returns a builder for updating this Frequency.
// Note that you need to call Frequency.Unwrap() before calling this method if this Frequency
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Frequency) Update() *FrequencyUpdateOne {
	return (&FrequencyClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Frequency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Frequency) Unwrap() *Frequency {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Frequency is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Frequency) String() string {
	var builder strings.Builder
	builder.WriteString("Frequency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("hash=")
	builder.WriteString(f.Hash)
	builder.WriteString(", ")
	builder.WriteString("import_flag=")
	builder.WriteString(fmt.Sprintf("%v", f.ImportFlag))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(f.Type)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(f.Description)
	builder.WriteString(", ")
	builder.WriteString("frequency=")
	builder.WriteString(fmt.Sprintf("%v", f.Frequency))
	builder.WriteByte(')')
	return builder.String()
}

// Frequencies is a parsable slice of Frequency.
type Frequencies []*Frequency

func (f Frequencies) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/country"
)

// Country is the model entity for the Country schema.
type Country struct {
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
	// The ISO 3166-1 alpha-2 code of the country. A handful of unofficial, non-ISO codes are also in use, such as "XK" for Kosovo.
	Code string `json:"code,omitempty"`
	// The name of the country.
	Name string `json:"name,omitempty"`
	// Where the airport is (primarily) located.
	Continent country.Continent `json:"continent,omitempty"`
	// The wikipedia link of the country.
	WikipediaLink string `json:"wikipedia_link,omitempty"`
	// Keywords that can be used to search for the country.
	Keywords []string `json:"keywords,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges CountryEdges `json:"edges"`
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// Airports holds the value of the airports edge.
	Airports []*Airport `json:"airports,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool

	namedAirports map[string][]*Airport
}

// AirportsOrErr returns the Airports value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) AirportsOrErr() ([]*Airport, error) {
	if e.loadedTypes[0] {
		return e.Airports, nil
	}
	return nil, &NotLoadedError{edge: "airports"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldKeywords:
			values[i] = new([]byte)
		case country.FieldImportFlag:
			values[i] = new(sql.NullBool)
		case country.FieldImportID:
			values[i] = new(sql.NullInt64)
		case country.FieldHash, country.FieldCode, country.FieldName, country.FieldContinent, country.FieldWikipediaLink:
			values[i] = new(sql.NullString)
		case country.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case country.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Country", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case country.FieldImportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field import_id", values[i])
			} else if value.Valid {
				c.ImportID = int(value.Int64)
			}
		case country.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				c.Hash = value.String
			}
		case country.FieldImportFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field import_flag", values[i])
			} else if value.Valid {
				c.ImportFlag = value.Bool
			}
		case country.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated", values[i])
			} else if value.Valid {
				c.LastUpdated = value.Time
			}
		case country.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		case country.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case country.FieldContinent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field continent", values[i])
			} else if value.Valid {
				c.Continent = country.Continent(value.String)
			}
		case country.FieldWikipediaLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wikipedia_link", values[i])
			} else if value.Valid {
				c.WikipediaLink = value.String
			}
		case country.FieldKeywords:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field keywords", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Keywords); err != nil {
					return fmt.Errorf("unmarshal field keywords: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryAirports queries the "airports" edge of the Country entity.
func (c *Country) QueryAirports() *AirportQuery {
	return NewCountryClient(c.config).QueryAirports(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("import_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ImportID))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(c.Hash)
	builder.WriteString(", ")
	builder.WriteString("import_flag=")
	builder.WriteString(fmt.Sprintf("%v", c.ImportFlag))
	builder.WriteString(", ")
	builder.WriteString("last_updated=")
	builder.WriteString(c.LastUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(c.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("continent=")
	builder.WriteString(fmt.Sprintf("%v", c.Continent))
	builder.WriteString(", ")
	builder.WriteString("wikipedia_link=")
	builder.WriteString(c.WikipediaLink)
	builder.WriteString(", ")
	builder.WriteString("keywords=")
	builder.WriteString(fmt.Sprintf("%v", c.Keywords))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAirports returns the Airports named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedAirports(name string) ([]*Airport, error) {
	if c.Edges.namedAirports == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAirports[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedAirports(name string, edges ...*Airport) {
	if c.Edges.namedAirports == nil {
		c.Edges.namedAirports = make(map[string][]*Airport)
	}
	if len(edges) == 0 {
		c.Edges.namedAirports[name] = []*Airport{}
	} else {
		c.Edges.namedAirports[name] = append(c.Edges.namedAirports[name], edges...)
	}
}

// Countries is a parsable slice of Country.
type Countries []*Country

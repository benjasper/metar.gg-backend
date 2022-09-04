// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent/metar"
	"metar.gg/ent/skycondition"
)

// SkyCondition is the model entity for the SkyCondition schema.
type SkyCondition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SkyCover holds the value of the "sky_cover" field.
	SkyCover skycondition.SkyCover `json:"sky_cover,omitempty"`
	// Cloud base in feet.
	CloudBase *int `json:"cloud_base,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkyConditionQuery when eager-loading is set.
	Edges                SkyConditionEdges `json:"edges"`
	metar_sky_conditions *int
}

// SkyConditionEdges holds the relations/edges for other nodes in the graph.
type SkyConditionEdges struct {
	// Metar holds the value of the metar edge.
	Metar *Metar `json:"metar,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// MetarOrErr returns the Metar value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkyConditionEdges) MetarOrErr() (*Metar, error) {
	if e.loadedTypes[0] {
		if e.Metar == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: metar.Label}
		}
		return e.Metar, nil
	}
	return nil, &NotLoadedError{edge: "metar"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SkyCondition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case skycondition.FieldID, skycondition.FieldCloudBase:
			values[i] = new(sql.NullInt64)
		case skycondition.FieldSkyCover:
			values[i] = new(sql.NullString)
		case skycondition.ForeignKeys[0]: // metar_sky_conditions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SkyCondition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SkyCondition fields.
func (sc *SkyCondition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skycondition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case skycondition.FieldSkyCover:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sky_cover", values[i])
			} else if value.Valid {
				sc.SkyCover = skycondition.SkyCover(value.String)
			}
		case skycondition.FieldCloudBase:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cloud_base", values[i])
			} else if value.Valid {
				sc.CloudBase = new(int)
				*sc.CloudBase = int(value.Int64)
			}
		case skycondition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field metar_sky_conditions", value)
			} else if value.Valid {
				sc.metar_sky_conditions = new(int)
				*sc.metar_sky_conditions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetar queries the "metar" edge of the SkyCondition entity.
func (sc *SkyCondition) QueryMetar() *MetarQuery {
	return (&SkyConditionClient{config: sc.config}).QueryMetar(sc)
}

// Update returns a builder for updating this SkyCondition.
// Note that you need to call SkyCondition.Unwrap() before calling this method if this SkyCondition
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SkyCondition) Update() *SkyConditionUpdateOne {
	return (&SkyConditionClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the SkyCondition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SkyCondition) Unwrap() *SkyCondition {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SkyCondition is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SkyCondition) String() string {
	var builder strings.Builder
	builder.WriteString("SkyCondition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("sky_cover=")
	builder.WriteString(fmt.Sprintf("%v", sc.SkyCover))
	builder.WriteString(", ")
	if v := sc.CloudBase; v != nil {
		builder.WriteString("cloud_base=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SkyConditions is a parsable slice of SkyCondition.
type SkyConditions []*SkyCondition

func (sc SkyConditions) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}

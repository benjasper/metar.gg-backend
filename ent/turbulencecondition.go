// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent/turbulencecondition"
)

// TurbulenceCondition is the model entity for the TurbulenceCondition schema.
type TurbulenceCondition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// The intensity of the turbulence.
	Intensity string `json:"intensity,omitempty"`
	// The minimum altitude in feet that the turbulence is present.
	MinAltitude int `json:"min_altitude,omitempty"`
	// The maximum altitude in feet that the turbulence is present.
	MaxAltitude                    int `json:"max_altitude,omitempty"`
	forecast_turbulence_conditions *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TurbulenceCondition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case turbulencecondition.FieldID, turbulencecondition.FieldMinAltitude, turbulencecondition.FieldMaxAltitude:
			values[i] = new(sql.NullInt64)
		case turbulencecondition.FieldIntensity:
			values[i] = new(sql.NullString)
		case turbulencecondition.ForeignKeys[0]: // forecast_turbulence_conditions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TurbulenceCondition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TurbulenceCondition fields.
func (tc *TurbulenceCondition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case turbulencecondition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int(value.Int64)
		case turbulencecondition.FieldIntensity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intensity", values[i])
			} else if value.Valid {
				tc.Intensity = value.String
			}
		case turbulencecondition.FieldMinAltitude:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_altitude", values[i])
			} else if value.Valid {
				tc.MinAltitude = int(value.Int64)
			}
		case turbulencecondition.FieldMaxAltitude:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_altitude", values[i])
			} else if value.Valid {
				tc.MaxAltitude = int(value.Int64)
			}
		case turbulencecondition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field forecast_turbulence_conditions", value)
			} else if value.Valid {
				tc.forecast_turbulence_conditions = new(int)
				*tc.forecast_turbulence_conditions = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TurbulenceCondition.
// Note that you need to call TurbulenceCondition.Unwrap() before calling this method if this TurbulenceCondition
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TurbulenceCondition) Update() *TurbulenceConditionUpdateOne {
	return (&TurbulenceConditionClient{config: tc.config}).UpdateOne(tc)
}

// Unwrap unwraps the TurbulenceCondition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TurbulenceCondition) Unwrap() *TurbulenceCondition {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TurbulenceCondition is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TurbulenceCondition) String() string {
	var builder strings.Builder
	builder.WriteString("TurbulenceCondition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("intensity=")
	builder.WriteString(tc.Intensity)
	builder.WriteString(", ")
	builder.WriteString("min_altitude=")
	builder.WriteString(fmt.Sprintf("%v", tc.MinAltitude))
	builder.WriteString(", ")
	builder.WriteString("max_altitude=")
	builder.WriteString(fmt.Sprintf("%v", tc.MaxAltitude))
	builder.WriteByte(')')
	return builder.String()
}

// TurbulenceConditions is a parsable slice of TurbulenceCondition.
type TurbulenceConditions []*TurbulenceCondition

func (tc TurbulenceConditions) config(cfg config) {
	for _i := range tc {
		tc[_i].config = cfg
	}
}

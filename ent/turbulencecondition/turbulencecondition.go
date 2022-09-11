// Code generated by ent, DO NOT EDIT.

package turbulencecondition

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the turbulencecondition type in the database.
	Label = "turbulence_condition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIntensity holds the string denoting the intensity field in the database.
	FieldIntensity = "intensity"
	// FieldMinAltitude holds the string denoting the min_altitude field in the database.
	FieldMinAltitude = "min_altitude"
	// FieldMaxAltitude holds the string denoting the max_altitude field in the database.
	FieldMaxAltitude = "max_altitude"
	// Table holds the table name of the turbulencecondition in the database.
	Table = "turbulence_conditions"
)

// Columns holds all SQL columns for turbulencecondition fields.
var Columns = []string{
	FieldID,
	FieldIntensity,
	FieldMinAltitude,
	FieldMaxAltitude,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "turbulence_conditions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"forecast_turbulence_conditions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

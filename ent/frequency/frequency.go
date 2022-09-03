// Code generated by ent, DO NOT EDIT.

package frequency

const (
	// Label holds the string label denoting the frequency type in the database.
	Label = "frequency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldImportFlag holds the string denoting the import_flag field in the database.
	FieldImportFlag = "import_flag"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFrequency holds the string denoting the frequency field in the database.
	FieldFrequency = "frequency"
	// EdgeAirport holds the string denoting the airport edge name in mutations.
	EdgeAirport = "airport"
	// Table holds the table name of the frequency in the database.
	Table = "frequencies"
	// AirportTable is the table that holds the airport relation/edge.
	AirportTable = "frequencies"
	// AirportInverseTable is the table name for the Airport entity.
	// It exists in this package in order to avoid circular dependency with the "airport" package.
	AirportInverseTable = "airports"
	// AirportColumn is the table column denoting the airport relation/edge.
	AirportColumn = "airport_frequencies"
)

// Columns holds all SQL columns for frequency fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldImportFlag,
	FieldType,
	FieldDescription,
	FieldFrequency,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "frequencies"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"airport_frequencies",
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
	// DefaultImportFlag holds the default value on creation for the "import_flag" field.
	DefaultImportFlag bool
)

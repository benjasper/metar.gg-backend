// Code generated by ent, DO NOT EDIT.

package runway

const (
	// Label holds the string label denoting the runway type in the database.
	Label = "runway"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldImportFlag holds the string denoting the import_flag field in the database.
	FieldImportFlag = "import_flag"
	// FieldAirportIdentifier holds the string denoting the airport_identifier field in the database.
	FieldAirportIdentifier = "airport_identifier"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldSurface holds the string denoting the surface field in the database.
	FieldSurface = "surface"
	// FieldLighted holds the string denoting the lighted field in the database.
	FieldLighted = "lighted"
	// FieldClosed holds the string denoting the closed field in the database.
	FieldClosed = "closed"
	// FieldLowRunwayIdentifier holds the string denoting the low_runway_identifier field in the database.
	FieldLowRunwayIdentifier = "low_runway_identifier"
	// FieldLowRunwayLatitude holds the string denoting the low_runway_latitude field in the database.
	FieldLowRunwayLatitude = "low_runway_latitude"
	// FieldLowRunwayLongitude holds the string denoting the low_runway_longitude field in the database.
	FieldLowRunwayLongitude = "low_runway_longitude"
	// FieldLowRunwayElevation holds the string denoting the low_runway_elevation field in the database.
	FieldLowRunwayElevation = "low_runway_elevation"
	// FieldLowRunwayHeading holds the string denoting the low_runway_heading field in the database.
	FieldLowRunwayHeading = "low_runway_heading"
	// FieldLowRunwayDisplaced holds the string denoting the low_runway_displaced field in the database.
	FieldLowRunwayDisplaced = "low_runway_displaced"
	// FieldHighRunwayIdentifier holds the string denoting the high_runway_identifier field in the database.
	FieldHighRunwayIdentifier = "high_runway_identifier"
	// FieldHighRunwayLatitude holds the string denoting the high_runway_latitude field in the database.
	FieldHighRunwayLatitude = "high_runway_latitude"
	// FieldHighRunwayLongitude holds the string denoting the high_runway_longitude field in the database.
	FieldHighRunwayLongitude = "high_runway_longitude"
	// FieldHighRunwayElevation holds the string denoting the high_runway_elevation field in the database.
	FieldHighRunwayElevation = "high_runway_elevation"
	// FieldHighRunwayHeading holds the string denoting the high_runway_heading field in the database.
	FieldHighRunwayHeading = "high_runway_heading"
	// FieldHighRunwayDisplaced holds the string denoting the high_runway_displaced field in the database.
	FieldHighRunwayDisplaced = "high_runway_displaced"
	// EdgeAirport holds the string denoting the airport edge name in mutations.
	EdgeAirport = "airport"
	// Table holds the table name of the runway in the database.
	Table = "runways"
	// AirportTable is the table that holds the airport relation/edge.
	AirportTable = "runways"
	// AirportInverseTable is the table name for the Airport entity.
	// It exists in this package in order to avoid circular dependency with the "airport" package.
	AirportInverseTable = "airports"
	// AirportColumn is the table column denoting the airport relation/edge.
	AirportColumn = "airport_runways"
)

// Columns holds all SQL columns for runway fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldImportFlag,
	FieldAirportIdentifier,
	FieldLength,
	FieldWidth,
	FieldSurface,
	FieldLighted,
	FieldClosed,
	FieldLowRunwayIdentifier,
	FieldLowRunwayLatitude,
	FieldLowRunwayLongitude,
	FieldLowRunwayElevation,
	FieldLowRunwayHeading,
	FieldLowRunwayDisplaced,
	FieldHighRunwayIdentifier,
	FieldHighRunwayLatitude,
	FieldHighRunwayLongitude,
	FieldHighRunwayElevation,
	FieldHighRunwayHeading,
	FieldHighRunwayDisplaced,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "runways"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"airport_runways",
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

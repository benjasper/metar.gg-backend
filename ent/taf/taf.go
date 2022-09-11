// Code generated by ent, DO NOT EDIT.

package taf

const (
	// Label holds the string label denoting the taf type in the database.
	Label = "taf"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRawText holds the string denoting the raw_text field in the database.
	FieldRawText = "raw_text"
	// FieldIssueTime holds the string denoting the issue_time field in the database.
	FieldIssueTime = "issue_time"
	// FieldBulletinTime holds the string denoting the bulletin_time field in the database.
	FieldBulletinTime = "bulletin_time"
	// FieldValidFromTime holds the string denoting the valid_from_time field in the database.
	FieldValidFromTime = "valid_from_time"
	// FieldValidToTime holds the string denoting the valid_to_time field in the database.
	FieldValidToTime = "valid_to_time"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// EdgeStation holds the string denoting the station edge name in mutations.
	EdgeStation = "station"
	// EdgeSkyConditions holds the string denoting the sky_conditions edge name in mutations.
	EdgeSkyConditions = "sky_conditions"
	// EdgeForecast holds the string denoting the forecast edge name in mutations.
	EdgeForecast = "forecast"
	// Table holds the table name of the taf in the database.
	Table = "tafs"
	// StationTable is the table that holds the station relation/edge.
	StationTable = "tafs"
	// StationInverseTable is the table name for the Station entity.
	// It exists in this package in order to avoid circular dependency with the "station" package.
	StationInverseTable = "stations"
	// StationColumn is the table column denoting the station relation/edge.
	StationColumn = "station_tafs"
	// SkyConditionsTable is the table that holds the sky_conditions relation/edge.
	SkyConditionsTable = "sky_conditions"
	// SkyConditionsInverseTable is the table name for the SkyCondition entity.
	// It exists in this package in order to avoid circular dependency with the "skycondition" package.
	SkyConditionsInverseTable = "sky_conditions"
	// SkyConditionsColumn is the table column denoting the sky_conditions relation/edge.
	SkyConditionsColumn = "taf_sky_conditions"
	// ForecastTable is the table that holds the forecast relation/edge.
	ForecastTable = "forecasts"
	// ForecastInverseTable is the table name for the Forecast entity.
	// It exists in this package in order to avoid circular dependency with the "forecast" package.
	ForecastInverseTable = "forecasts"
	// ForecastColumn is the table column denoting the forecast relation/edge.
	ForecastColumn = "taf_forecast"
)

// Columns holds all SQL columns for taf fields.
var Columns = []string{
	FieldID,
	FieldRawText,
	FieldIssueTime,
	FieldBulletinTime,
	FieldValidFromTime,
	FieldValidToTime,
	FieldRemarks,
	FieldHash,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tafs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"station_tafs",
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
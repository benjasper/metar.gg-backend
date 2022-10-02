// Code generated by ent, DO NOT EDIT.

package metar

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the metar type in the database.
	Label = "metar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRawText holds the string denoting the raw_text field in the database.
	FieldRawText = "raw_text"
	// FieldObservationTime holds the string denoting the observation_time field in the database.
	FieldObservationTime = "observation_time"
	// FieldImportTime holds the string denoting the import_time field in the database.
	FieldImportTime = "import_time"
	// FieldNextImportTimePrediction holds the string denoting the next_import_time_prediction field in the database.
	FieldNextImportTimePrediction = "next_import_time_prediction"
	// FieldTemperature holds the string denoting the temperature field in the database.
	FieldTemperature = "temperature"
	// FieldDewpoint holds the string denoting the dewpoint field in the database.
	FieldDewpoint = "dewpoint"
	// FieldWindSpeed holds the string denoting the wind_speed field in the database.
	FieldWindSpeed = "wind_speed"
	// FieldWindGust holds the string denoting the wind_gust field in the database.
	FieldWindGust = "wind_gust"
	// FieldWindDirection holds the string denoting the wind_direction field in the database.
	FieldWindDirection = "wind_direction"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldAltimeter holds the string denoting the altimeter field in the database.
	FieldAltimeter = "altimeter"
	// FieldPresentWeather holds the string denoting the present_weather field in the database.
	FieldPresentWeather = "present_weather"
	// FieldFlightCategory holds the string denoting the flight_category field in the database.
	FieldFlightCategory = "flight_category"
	// FieldQualityControlCorrected holds the string denoting the quality_control_corrected field in the database.
	FieldQualityControlCorrected = "quality_control_corrected"
	// FieldQualityControlAutoStation holds the string denoting the quality_control_auto_station field in the database.
	FieldQualityControlAutoStation = "quality_control_auto_station"
	// FieldQualityControlMaintenanceIndicatorOn holds the string denoting the quality_control_maintenance_indicator_on field in the database.
	FieldQualityControlMaintenanceIndicatorOn = "quality_control_maintenance_indicator_on"
	// FieldQualityControlNoSignal holds the string denoting the quality_control_no_signal field in the database.
	FieldQualityControlNoSignal = "quality_control_no_signal"
	// FieldQualityControlLightningSensorOff holds the string denoting the quality_control_lightning_sensor_off field in the database.
	FieldQualityControlLightningSensorOff = "quality_control_lightning_sensor_off"
	// FieldQualityControlFreezingRainSensorOff holds the string denoting the quality_control_freezing_rain_sensor_off field in the database.
	FieldQualityControlFreezingRainSensorOff = "quality_control_freezing_rain_sensor_off"
	// FieldQualityControlPresentWeatherSensorOff holds the string denoting the quality_control_present_weather_sensor_off field in the database.
	FieldQualityControlPresentWeatherSensorOff = "quality_control_present_weather_sensor_off"
	// FieldSeaLevelPressure holds the string denoting the sea_level_pressure field in the database.
	FieldSeaLevelPressure = "sea_level_pressure"
	// FieldPressureTendency holds the string denoting the pressure_tendency field in the database.
	FieldPressureTendency = "pressure_tendency"
	// FieldMaxTemp6 holds the string denoting the max_temp_6 field in the database.
	FieldMaxTemp6 = "max_temp_6"
	// FieldMinTemp6 holds the string denoting the min_temp_6 field in the database.
	FieldMinTemp6 = "min_temp_6"
	// FieldMaxTemp24 holds the string denoting the max_temp_24 field in the database.
	FieldMaxTemp24 = "max_temp_24"
	// FieldMinTemp24 holds the string denoting the min_temp_24 field in the database.
	FieldMinTemp24 = "min_temp_24"
	// FieldPrecipitation holds the string denoting the precipitation field in the database.
	FieldPrecipitation = "precipitation"
	// FieldPrecipitation3 holds the string denoting the precipitation_3 field in the database.
	FieldPrecipitation3 = "precipitation_3"
	// FieldPrecipitation6 holds the string denoting the precipitation_6 field in the database.
	FieldPrecipitation6 = "precipitation_6"
	// FieldPrecipitation24 holds the string denoting the precipitation_24 field in the database.
	FieldPrecipitation24 = "precipitation_24"
	// FieldSnowDepth holds the string denoting the snow_depth field in the database.
	FieldSnowDepth = "snow_depth"
	// FieldVertVis holds the string denoting the vert_vis field in the database.
	FieldVertVis = "vert_vis"
	// FieldMetarType holds the string denoting the metar_type field in the database.
	FieldMetarType = "metar_type"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// EdgeStation holds the string denoting the station edge name in mutations.
	EdgeStation = "station"
	// EdgeSkyConditions holds the string denoting the sky_conditions edge name in mutations.
	EdgeSkyConditions = "sky_conditions"
	// Table holds the table name of the metar in the database.
	Table = "metars"
	// StationTable is the table that holds the station relation/edge.
	StationTable = "metars"
	// StationInverseTable is the table name for the WeatherStation entity.
	// It exists in this package in order to avoid circular dependency with the "weatherstation" package.
	StationInverseTable = "weather_stations"
	// StationColumn is the table column denoting the station relation/edge.
	StationColumn = "weather_station_metars"
	// SkyConditionsTable is the table that holds the sky_conditions relation/edge.
	SkyConditionsTable = "sky_conditions"
	// SkyConditionsInverseTable is the table name for the SkyCondition entity.
	// It exists in this package in order to avoid circular dependency with the "skycondition" package.
	SkyConditionsInverseTable = "sky_conditions"
	// SkyConditionsColumn is the table column denoting the sky_conditions relation/edge.
	SkyConditionsColumn = "metar_sky_conditions"
)

// Columns holds all SQL columns for metar fields.
var Columns = []string{
	FieldID,
	FieldRawText,
	FieldObservationTime,
	FieldImportTime,
	FieldNextImportTimePrediction,
	FieldTemperature,
	FieldDewpoint,
	FieldWindSpeed,
	FieldWindGust,
	FieldWindDirection,
	FieldVisibility,
	FieldAltimeter,
	FieldPresentWeather,
	FieldFlightCategory,
	FieldQualityControlCorrected,
	FieldQualityControlAutoStation,
	FieldQualityControlMaintenanceIndicatorOn,
	FieldQualityControlNoSignal,
	FieldQualityControlLightningSensorOff,
	FieldQualityControlFreezingRainSensorOff,
	FieldQualityControlPresentWeatherSensorOff,
	FieldSeaLevelPressure,
	FieldPressureTendency,
	FieldMaxTemp6,
	FieldMinTemp6,
	FieldMaxTemp24,
	FieldMinTemp24,
	FieldPrecipitation,
	FieldPrecipitation3,
	FieldPrecipitation6,
	FieldPrecipitation24,
	FieldSnowDepth,
	FieldVertVis,
	FieldMetarType,
	FieldHash,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "metars"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"weather_station_metars",
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
	// DefaultImportTime holds the default value on creation for the "import_time" field.
	DefaultImportTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// FlightCategory defines the type for the "flight_category" enum field.
type FlightCategory string

// FlightCategory values.
const (
	FlightCategoryVFR  FlightCategory = "VFR"
	FlightCategoryMVFR FlightCategory = "MVFR"
	FlightCategoryIFR  FlightCategory = "IFR"
	FlightCategoryLIFR FlightCategory = "LIFR"
)

func (fc FlightCategory) String() string {
	return string(fc)
}

// FlightCategoryValidator is a validator for the "flight_category" field enum values. It is called by the builders before save.
func FlightCategoryValidator(fc FlightCategory) error {
	switch fc {
	case FlightCategoryVFR, FlightCategoryMVFR, FlightCategoryIFR, FlightCategoryLIFR:
		return nil
	default:
		return fmt.Errorf("metar: invalid enum value for flight_category field: %q", fc)
	}
}

// MetarType defines the type for the "metar_type" enum field.
type MetarType string

// MetarType values.
const (
	MetarTypeMETAR MetarType = "METAR"
	MetarTypeSPECI MetarType = "SPECI"
)

func (mt MetarType) String() string {
	return string(mt)
}

// MetarTypeValidator is a validator for the "metar_type" field enum values. It is called by the builders before save.
func MetarTypeValidator(mt MetarType) error {
	switch mt {
	case MetarTypeMETAR, MetarTypeSPECI:
		return nil
	default:
		return fmt.Errorf("metar: invalid enum value for metar_type field: %q", mt)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e FlightCategory) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *FlightCategory) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = FlightCategory(str)
	if err := FlightCategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid FlightCategory", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e MetarType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *MetarType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = MetarType(str)
	if err := MetarTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid MetarType", str)
	}
	return nil
}

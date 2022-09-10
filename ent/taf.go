// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent/station"
	"metar.gg/ent/taf"
)

// Taf is the model entity for the Taf schema.
type Taf struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// The raw TAF text.
	RawText string `json:"raw_text,omitempty"`
	// The time the TAF was issued.
	IssueTime time.Time `json:"issue_time,omitempty"`
	// TAF bulletin time.
	BulletinTime time.Time `json:"bulletin_time,omitempty"`
	// The start time of the TAF validity period.
	ValidFromTime time.Time `json:"valid_from_time,omitempty"`
	// The end time of the TAF validity period.
	ValidToTime time.Time `json:"valid_to_time,omitempty"`
	// Remarks.
	Remarks string `json:"remarks,omitempty"`
	// The temperature in Celsius.
	Temperature float64 `json:"temperature,omitempty"`
	// The dewpoint in Celsius.
	Dewpoint float64 `json:"dewpoint,omitempty"`
	// The wind speed in knots, or 0 if calm.
	WindSpeed int `json:"wind_speed,omitempty"`
	// The wind gust in knots.
	WindGust int `json:"wind_gust,omitempty"`
	// The wind direction in degrees, or 0 if calm.
	WindDirection int `json:"wind_direction,omitempty"`
	// The visibility in statute miles.
	Visibility float64 `json:"visibility,omitempty"`
	// The altimeter setting in inches of mercury.
	Altimeter float64 `json:"altimeter,omitempty"`
	// FlightCategory holds the value of the "flight_category" field.
	FlightCategory *taf.FlightCategory `json:"flight_category,omitempty"`
	// Quality control corrected.
	QualityControlCorrected *bool `json:"quality_control_corrected,omitempty"`
	// Whether it's an automated station, of one of the following types A01|A01A|A02|A02A|AOA|AWOS.
	QualityControlAutoStation bool `json:"quality_control_auto_station,omitempty"`
	// Maintenance check indicator - maintenance is needed.
	QualityControlMaintenanceIndicatorOn bool `json:"quality_control_maintenance_indicator_on,omitempty"`
	// No signal.
	QualityControlNoSignal bool `json:"quality_control_no_signal,omitempty"`
	// Whether Lightning sensor is off.
	QualityControlLightningSensorOff bool `json:"quality_control_lightning_sensor_off,omitempty"`
	// Whether Freezing rain sensor is off.
	QualityControlFreezingRainSensorOff bool `json:"quality_control_freezing_rain_sensor_off,omitempty"`
	// Whether Present weather sensor is off.
	QualityControlPresentWeatherSensorOff bool `json:"quality_control_present_weather_sensor_off,omitempty"`
	// The sea level pressure in hectopascal.s
	SeaLevelPressure *float64 `json:"sea_level_pressure,omitempty"`
	// The pressur_6e tendency in hectopascals.
	PressureTendency *float64 `json:"pressure_tendency,omitempty"`
	// The maximum air temperature in Celsius from the past 6 hours.
	MaxTemp6 *float64 `json:"max_temp_6,omitempty"`
	// The minimum air temperature in Celsius from the past 6 hours.
	MinTemp6 *float64 `json:"min_temp_6,omitempty"`
	// The maximum air temperature in Celsius from the past 24 hours.
	MaxTemp24 *float64 `json:"max_temp_24,omitempty"`
	// The minimum air temperature in Celsius from the past 24 hours.
	MinTemp24 *float64 `json:"min_temp_24,omitempty"`
	// The precipitation in inches from since the last observation. 0.0005 in = trace precipitation
	Precipitation *float64 `json:"precipitation,omitempty"`
	// The precipitation in inches from the past 3 hours. 0.0005 in = trace precipitation
	Precipitation3 *float64 `json:"precipitation_3,omitempty"`
	// The precipitation in inches from the past 6 hours. 0.0005 in = trace precipitation
	Precipitation6 *float64 `json:"precipitation_6,omitempty"`
	// The precipitation in inches from the past 24 hours. 0.0005 in = trace precipitation
	Precipitation24 *float64 `json:"precipitation_24,omitempty"`
	// The snow depth in inches.
	SnowDepth *float64 `json:"snow_depth,omitempty"`
	// The vertical visibility in feet.
	VertVis *float64 `json:"vert_vis,omitempty"`
	// The type of TAF.
	MetarType taf.MetarType `json:"metar_type,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TafQuery when eager-loading is set.
	Edges        TafEdges `json:"edges"`
	station_tafs *int
}

// TafEdges holds the relations/edges for other nodes in the graph.
type TafEdges struct {
	// The station that issued this taf.
	Station *Station `json:"station,omitempty"`
	// The sky conditions.
	SkyConditions []*SkyCondition `json:"sky_conditions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedSkyConditions map[string][]*SkyCondition
}

// StationOrErr returns the Station value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TafEdges) StationOrErr() (*Station, error) {
	if e.loadedTypes[0] {
		if e.Station == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: station.Label}
		}
		return e.Station, nil
	}
	return nil, &NotLoadedError{edge: "station"}
}

// SkyConditionsOrErr returns the SkyConditions value or an error if the edge
// was not loaded in eager-loading.
func (e TafEdges) SkyConditionsOrErr() ([]*SkyCondition, error) {
	if e.loadedTypes[1] {
		return e.SkyConditions, nil
	}
	return nil, &NotLoadedError{edge: "sky_conditions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Taf) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case taf.FieldQualityControlCorrected, taf.FieldQualityControlAutoStation, taf.FieldQualityControlMaintenanceIndicatorOn, taf.FieldQualityControlNoSignal, taf.FieldQualityControlLightningSensorOff, taf.FieldQualityControlFreezingRainSensorOff, taf.FieldQualityControlPresentWeatherSensorOff:
			values[i] = new(sql.NullBool)
		case taf.FieldTemperature, taf.FieldDewpoint, taf.FieldVisibility, taf.FieldAltimeter, taf.FieldSeaLevelPressure, taf.FieldPressureTendency, taf.FieldMaxTemp6, taf.FieldMinTemp6, taf.FieldMaxTemp24, taf.FieldMinTemp24, taf.FieldPrecipitation, taf.FieldPrecipitation3, taf.FieldPrecipitation6, taf.FieldPrecipitation24, taf.FieldSnowDepth, taf.FieldVertVis:
			values[i] = new(sql.NullFloat64)
		case taf.FieldID, taf.FieldWindSpeed, taf.FieldWindGust, taf.FieldWindDirection:
			values[i] = new(sql.NullInt64)
		case taf.FieldRawText, taf.FieldRemarks, taf.FieldFlightCategory, taf.FieldMetarType, taf.FieldHash:
			values[i] = new(sql.NullString)
		case taf.FieldIssueTime, taf.FieldBulletinTime, taf.FieldValidFromTime, taf.FieldValidToTime:
			values[i] = new(sql.NullTime)
		case taf.ForeignKeys[0]: // station_tafs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Taf", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Taf fields.
func (t *Taf) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taf.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case taf.FieldRawText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_text", values[i])
			} else if value.Valid {
				t.RawText = value.String
			}
		case taf.FieldIssueTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field issue_time", values[i])
			} else if value.Valid {
				t.IssueTime = value.Time
			}
		case taf.FieldBulletinTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field bulletin_time", values[i])
			} else if value.Valid {
				t.BulletinTime = value.Time
			}
		case taf.FieldValidFromTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field valid_from_time", values[i])
			} else if value.Valid {
				t.ValidFromTime = value.Time
			}
		case taf.FieldValidToTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field valid_to_time", values[i])
			} else if value.Valid {
				t.ValidToTime = value.Time
			}
		case taf.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				t.Remarks = value.String
			}
		case taf.FieldTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temperature", values[i])
			} else if value.Valid {
				t.Temperature = value.Float64
			}
		case taf.FieldDewpoint:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field dewpoint", values[i])
			} else if value.Valid {
				t.Dewpoint = value.Float64
			}
		case taf.FieldWindSpeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_speed", values[i])
			} else if value.Valid {
				t.WindSpeed = int(value.Int64)
			}
		case taf.FieldWindGust:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_gust", values[i])
			} else if value.Valid {
				t.WindGust = int(value.Int64)
			}
		case taf.FieldWindDirection:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_direction", values[i])
			} else if value.Valid {
				t.WindDirection = int(value.Int64)
			}
		case taf.FieldVisibility:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				t.Visibility = value.Float64
			}
		case taf.FieldAltimeter:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field altimeter", values[i])
			} else if value.Valid {
				t.Altimeter = value.Float64
			}
		case taf.FieldFlightCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_category", values[i])
			} else if value.Valid {
				t.FlightCategory = new(taf.FlightCategory)
				*t.FlightCategory = taf.FlightCategory(value.String)
			}
		case taf.FieldQualityControlCorrected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_corrected", values[i])
			} else if value.Valid {
				t.QualityControlCorrected = new(bool)
				*t.QualityControlCorrected = value.Bool
			}
		case taf.FieldQualityControlAutoStation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_auto_station", values[i])
			} else if value.Valid {
				t.QualityControlAutoStation = value.Bool
			}
		case taf.FieldQualityControlMaintenanceIndicatorOn:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_maintenance_indicator_on", values[i])
			} else if value.Valid {
				t.QualityControlMaintenanceIndicatorOn = value.Bool
			}
		case taf.FieldQualityControlNoSignal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_no_signal", values[i])
			} else if value.Valid {
				t.QualityControlNoSignal = value.Bool
			}
		case taf.FieldQualityControlLightningSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_lightning_sensor_off", values[i])
			} else if value.Valid {
				t.QualityControlLightningSensorOff = value.Bool
			}
		case taf.FieldQualityControlFreezingRainSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_freezing_rain_sensor_off", values[i])
			} else if value.Valid {
				t.QualityControlFreezingRainSensorOff = value.Bool
			}
		case taf.FieldQualityControlPresentWeatherSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_present_weather_sensor_off", values[i])
			} else if value.Valid {
				t.QualityControlPresentWeatherSensorOff = value.Bool
			}
		case taf.FieldSeaLevelPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sea_level_pressure", values[i])
			} else if value.Valid {
				t.SeaLevelPressure = new(float64)
				*t.SeaLevelPressure = value.Float64
			}
		case taf.FieldPressureTendency:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pressure_tendency", values[i])
			} else if value.Valid {
				t.PressureTendency = new(float64)
				*t.PressureTendency = value.Float64
			}
		case taf.FieldMaxTemp6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp_6", values[i])
			} else if value.Valid {
				t.MaxTemp6 = new(float64)
				*t.MaxTemp6 = value.Float64
			}
		case taf.FieldMinTemp6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp_6", values[i])
			} else if value.Valid {
				t.MinTemp6 = new(float64)
				*t.MinTemp6 = value.Float64
			}
		case taf.FieldMaxTemp24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp_24", values[i])
			} else if value.Valid {
				t.MaxTemp24 = new(float64)
				*t.MaxTemp24 = value.Float64
			}
		case taf.FieldMinTemp24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp_24", values[i])
			} else if value.Valid {
				t.MinTemp24 = new(float64)
				*t.MinTemp24 = value.Float64
			}
		case taf.FieldPrecipitation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation", values[i])
			} else if value.Valid {
				t.Precipitation = new(float64)
				*t.Precipitation = value.Float64
			}
		case taf.FieldPrecipitation3:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_3", values[i])
			} else if value.Valid {
				t.Precipitation3 = new(float64)
				*t.Precipitation3 = value.Float64
			}
		case taf.FieldPrecipitation6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_6", values[i])
			} else if value.Valid {
				t.Precipitation6 = new(float64)
				*t.Precipitation6 = value.Float64
			}
		case taf.FieldPrecipitation24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_24", values[i])
			} else if value.Valid {
				t.Precipitation24 = new(float64)
				*t.Precipitation24 = value.Float64
			}
		case taf.FieldSnowDepth:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field snow_depth", values[i])
			} else if value.Valid {
				t.SnowDepth = new(float64)
				*t.SnowDepth = value.Float64
			}
		case taf.FieldVertVis:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vert_vis", values[i])
			} else if value.Valid {
				t.VertVis = new(float64)
				*t.VertVis = value.Float64
			}
		case taf.FieldMetarType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metar_type", values[i])
			} else if value.Valid {
				t.MetarType = taf.MetarType(value.String)
			}
		case taf.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				t.Hash = value.String
			}
		case taf.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field station_tafs", value)
			} else if value.Valid {
				t.station_tafs = new(int)
				*t.station_tafs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryStation queries the "station" edge of the Taf entity.
func (t *Taf) QueryStation() *StationQuery {
	return (&TafClient{config: t.config}).QueryStation(t)
}

// QuerySkyConditions queries the "sky_conditions" edge of the Taf entity.
func (t *Taf) QuerySkyConditions() *SkyConditionQuery {
	return (&TafClient{config: t.config}).QuerySkyConditions(t)
}

// Update returns a builder for updating this Taf.
// Note that you need to call Taf.Unwrap() before calling this method if this Taf
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Taf) Update() *TafUpdateOne {
	return (&TafClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Taf entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Taf) Unwrap() *Taf {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Taf is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Taf) String() string {
	var builder strings.Builder
	builder.WriteString("Taf(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("raw_text=")
	builder.WriteString(t.RawText)
	builder.WriteString(", ")
	builder.WriteString("issue_time=")
	builder.WriteString(t.IssueTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("bulletin_time=")
	builder.WriteString(t.BulletinTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("valid_from_time=")
	builder.WriteString(t.ValidFromTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("valid_to_time=")
	builder.WriteString(t.ValidToTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(t.Remarks)
	builder.WriteString(", ")
	builder.WriteString("temperature=")
	builder.WriteString(fmt.Sprintf("%v", t.Temperature))
	builder.WriteString(", ")
	builder.WriteString("dewpoint=")
	builder.WriteString(fmt.Sprintf("%v", t.Dewpoint))
	builder.WriteString(", ")
	builder.WriteString("wind_speed=")
	builder.WriteString(fmt.Sprintf("%v", t.WindSpeed))
	builder.WriteString(", ")
	builder.WriteString("wind_gust=")
	builder.WriteString(fmt.Sprintf("%v", t.WindGust))
	builder.WriteString(", ")
	builder.WriteString("wind_direction=")
	builder.WriteString(fmt.Sprintf("%v", t.WindDirection))
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", t.Visibility))
	builder.WriteString(", ")
	builder.WriteString("altimeter=")
	builder.WriteString(fmt.Sprintf("%v", t.Altimeter))
	builder.WriteString(", ")
	if v := t.FlightCategory; v != nil {
		builder.WriteString("flight_category=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.QualityControlCorrected; v != nil {
		builder.WriteString("quality_control_corrected=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("quality_control_auto_station=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlAutoStation))
	builder.WriteString(", ")
	builder.WriteString("quality_control_maintenance_indicator_on=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlMaintenanceIndicatorOn))
	builder.WriteString(", ")
	builder.WriteString("quality_control_no_signal=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlNoSignal))
	builder.WriteString(", ")
	builder.WriteString("quality_control_lightning_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlLightningSensorOff))
	builder.WriteString(", ")
	builder.WriteString("quality_control_freezing_rain_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlFreezingRainSensorOff))
	builder.WriteString(", ")
	builder.WriteString("quality_control_present_weather_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", t.QualityControlPresentWeatherSensorOff))
	builder.WriteString(", ")
	if v := t.SeaLevelPressure; v != nil {
		builder.WriteString("sea_level_pressure=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.PressureTendency; v != nil {
		builder.WriteString("pressure_tendency=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.MaxTemp6; v != nil {
		builder.WriteString("max_temp_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.MinTemp6; v != nil {
		builder.WriteString("min_temp_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.MaxTemp24; v != nil {
		builder.WriteString("max_temp_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.MinTemp24; v != nil {
		builder.WriteString("min_temp_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.Precipitation; v != nil {
		builder.WriteString("precipitation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.Precipitation3; v != nil {
		builder.WriteString("precipitation_3=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.Precipitation6; v != nil {
		builder.WriteString("precipitation_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.Precipitation24; v != nil {
		builder.WriteString("precipitation_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.SnowDepth; v != nil {
		builder.WriteString("snow_depth=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.VertVis; v != nil {
		builder.WriteString("vert_vis=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("metar_type=")
	builder.WriteString(fmt.Sprintf("%v", t.MetarType))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(t.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSkyConditions returns the SkyConditions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Taf) NamedSkyConditions(name string) ([]*SkyCondition, error) {
	if t.Edges.namedSkyConditions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedSkyConditions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Taf) appendNamedSkyConditions(name string, edges ...*SkyCondition) {
	if t.Edges.namedSkyConditions == nil {
		t.Edges.namedSkyConditions = make(map[string][]*SkyCondition)
	}
	if len(edges) == 0 {
		t.Edges.namedSkyConditions[name] = []*SkyCondition{}
	} else {
		t.Edges.namedSkyConditions[name] = append(t.Edges.namedSkyConditions[name], edges...)
	}
}

// Tafs is a parsable slice of Taf.
type Tafs []*Taf

func (t Tafs) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}

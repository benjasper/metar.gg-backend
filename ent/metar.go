// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
)

// Metar is the model entity for the Metar schema.
type Metar struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// The ICAO identifier of the station that provided the METAR or identifier of the weather station.
	StationID string `json:"station_id,omitempty"`
	// The raw METAR text.
	RawText string `json:"raw_text,omitempty"`
	// The time the METAR was observed.
	ObservationTime time.Time `json:"observation_time,omitempty"`
	// The latitude in decimal degrees of the station.
	Latitude *float64 `json:"latitude,omitempty"`
	// The longitude in decimal degrees of the station.
	Longitude *float64 `json:"longitude,omitempty"`
	// The elevation in meters of the station.
	Elevation *float64 `json:"elevation,omitempty"`
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
	// The present weather string.
	PresentWeather *string `json:"present_weather,omitempty"`
	// FlightCategory holds the value of the "flight_category" field.
	FlightCategory *metar.FlightCategory `json:"flight_category,omitempty"`
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
	// The type of METAR.
	MetarType metar.MetarType `json:"metar_type,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetarQuery when eager-loading is set.
	Edges          MetarEdges `json:"edges"`
	airport_metars *int
}

// MetarEdges holds the relations/edges for other nodes in the graph.
type MetarEdges struct {
	// The airport that reported this metar. This can also be empty if the metar is from a weather station.
	Airport *Airport `json:"airport,omitempty"`
	// The sky conditions.
	SkyConditions []*SkyCondition `json:"sky_conditions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedSkyConditions map[string][]*SkyCondition
}

// AirportOrErr returns the Airport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetarEdges) AirportOrErr() (*Airport, error) {
	if e.loadedTypes[0] {
		if e.Airport == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Airport, nil
	}
	return nil, &NotLoadedError{edge: "airport"}
}

// SkyConditionsOrErr returns the SkyConditions value or an error if the edge
// was not loaded in eager-loading.
func (e MetarEdges) SkyConditionsOrErr() ([]*SkyCondition, error) {
	if e.loadedTypes[1] {
		return e.SkyConditions, nil
	}
	return nil, &NotLoadedError{edge: "sky_conditions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metar) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case metar.FieldQualityControlCorrected, metar.FieldQualityControlAutoStation, metar.FieldQualityControlMaintenanceIndicatorOn, metar.FieldQualityControlNoSignal, metar.FieldQualityControlLightningSensorOff, metar.FieldQualityControlFreezingRainSensorOff, metar.FieldQualityControlPresentWeatherSensorOff:
			values[i] = new(sql.NullBool)
		case metar.FieldLatitude, metar.FieldLongitude, metar.FieldElevation, metar.FieldTemperature, metar.FieldDewpoint, metar.FieldVisibility, metar.FieldAltimeter, metar.FieldSeaLevelPressure, metar.FieldPressureTendency, metar.FieldMaxTemp6, metar.FieldMinTemp6, metar.FieldMaxTemp24, metar.FieldMinTemp24, metar.FieldPrecipitation, metar.FieldPrecipitation3, metar.FieldPrecipitation6, metar.FieldPrecipitation24, metar.FieldSnowDepth, metar.FieldVertVis:
			values[i] = new(sql.NullFloat64)
		case metar.FieldID, metar.FieldWindSpeed, metar.FieldWindGust, metar.FieldWindDirection:
			values[i] = new(sql.NullInt64)
		case metar.FieldStationID, metar.FieldRawText, metar.FieldPresentWeather, metar.FieldFlightCategory, metar.FieldMetarType, metar.FieldHash:
			values[i] = new(sql.NullString)
		case metar.FieldObservationTime:
			values[i] = new(sql.NullTime)
		case metar.ForeignKeys[0]: // airport_metars
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Metar", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metar fields.
func (m *Metar) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metar.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case metar.FieldStationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field station_id", values[i])
			} else if value.Valid {
				m.StationID = value.String
			}
		case metar.FieldRawText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_text", values[i])
			} else if value.Valid {
				m.RawText = value.String
			}
		case metar.FieldObservationTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field observation_time", values[i])
			} else if value.Valid {
				m.ObservationTime = value.Time
			}
		case metar.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				m.Latitude = new(float64)
				*m.Latitude = value.Float64
			}
		case metar.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				m.Longitude = new(float64)
				*m.Longitude = value.Float64
			}
		case metar.FieldElevation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field elevation", values[i])
			} else if value.Valid {
				m.Elevation = new(float64)
				*m.Elevation = value.Float64
			}
		case metar.FieldTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temperature", values[i])
			} else if value.Valid {
				m.Temperature = value.Float64
			}
		case metar.FieldDewpoint:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field dewpoint", values[i])
			} else if value.Valid {
				m.Dewpoint = value.Float64
			}
		case metar.FieldWindSpeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_speed", values[i])
			} else if value.Valid {
				m.WindSpeed = int(value.Int64)
			}
		case metar.FieldWindGust:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_gust", values[i])
			} else if value.Valid {
				m.WindGust = int(value.Int64)
			}
		case metar.FieldWindDirection:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wind_direction", values[i])
			} else if value.Valid {
				m.WindDirection = int(value.Int64)
			}
		case metar.FieldVisibility:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				m.Visibility = value.Float64
			}
		case metar.FieldAltimeter:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field altimeter", values[i])
			} else if value.Valid {
				m.Altimeter = value.Float64
			}
		case metar.FieldPresentWeather:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field present_weather", values[i])
			} else if value.Valid {
				m.PresentWeather = new(string)
				*m.PresentWeather = value.String
			}
		case metar.FieldFlightCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_category", values[i])
			} else if value.Valid {
				m.FlightCategory = new(metar.FlightCategory)
				*m.FlightCategory = metar.FlightCategory(value.String)
			}
		case metar.FieldQualityControlCorrected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_corrected", values[i])
			} else if value.Valid {
				m.QualityControlCorrected = new(bool)
				*m.QualityControlCorrected = value.Bool
			}
		case metar.FieldQualityControlAutoStation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_auto_station", values[i])
			} else if value.Valid {
				m.QualityControlAutoStation = value.Bool
			}
		case metar.FieldQualityControlMaintenanceIndicatorOn:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_maintenance_indicator_on", values[i])
			} else if value.Valid {
				m.QualityControlMaintenanceIndicatorOn = value.Bool
			}
		case metar.FieldQualityControlNoSignal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_no_signal", values[i])
			} else if value.Valid {
				m.QualityControlNoSignal = value.Bool
			}
		case metar.FieldQualityControlLightningSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_lightning_sensor_off", values[i])
			} else if value.Valid {
				m.QualityControlLightningSensorOff = value.Bool
			}
		case metar.FieldQualityControlFreezingRainSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_freezing_rain_sensor_off", values[i])
			} else if value.Valid {
				m.QualityControlFreezingRainSensorOff = value.Bool
			}
		case metar.FieldQualityControlPresentWeatherSensorOff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field quality_control_present_weather_sensor_off", values[i])
			} else if value.Valid {
				m.QualityControlPresentWeatherSensorOff = value.Bool
			}
		case metar.FieldSeaLevelPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sea_level_pressure", values[i])
			} else if value.Valid {
				m.SeaLevelPressure = new(float64)
				*m.SeaLevelPressure = value.Float64
			}
		case metar.FieldPressureTendency:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pressure_tendency", values[i])
			} else if value.Valid {
				m.PressureTendency = new(float64)
				*m.PressureTendency = value.Float64
			}
		case metar.FieldMaxTemp6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp_6", values[i])
			} else if value.Valid {
				m.MaxTemp6 = new(float64)
				*m.MaxTemp6 = value.Float64
			}
		case metar.FieldMinTemp6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp_6", values[i])
			} else if value.Valid {
				m.MinTemp6 = new(float64)
				*m.MinTemp6 = value.Float64
			}
		case metar.FieldMaxTemp24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp_24", values[i])
			} else if value.Valid {
				m.MaxTemp24 = new(float64)
				*m.MaxTemp24 = value.Float64
			}
		case metar.FieldMinTemp24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp_24", values[i])
			} else if value.Valid {
				m.MinTemp24 = new(float64)
				*m.MinTemp24 = value.Float64
			}
		case metar.FieldPrecipitation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation", values[i])
			} else if value.Valid {
				m.Precipitation = new(float64)
				*m.Precipitation = value.Float64
			}
		case metar.FieldPrecipitation3:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_3", values[i])
			} else if value.Valid {
				m.Precipitation3 = new(float64)
				*m.Precipitation3 = value.Float64
			}
		case metar.FieldPrecipitation6:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_6", values[i])
			} else if value.Valid {
				m.Precipitation6 = new(float64)
				*m.Precipitation6 = value.Float64
			}
		case metar.FieldPrecipitation24:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field precipitation_24", values[i])
			} else if value.Valid {
				m.Precipitation24 = new(float64)
				*m.Precipitation24 = value.Float64
			}
		case metar.FieldSnowDepth:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field snow_depth", values[i])
			} else if value.Valid {
				m.SnowDepth = new(float64)
				*m.SnowDepth = value.Float64
			}
		case metar.FieldVertVis:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vert_vis", values[i])
			} else if value.Valid {
				m.VertVis = new(float64)
				*m.VertVis = value.Float64
			}
		case metar.FieldMetarType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metar_type", values[i])
			} else if value.Valid {
				m.MetarType = metar.MetarType(value.String)
			}
		case metar.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				m.Hash = value.String
			}
		case metar.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field airport_metars", value)
			} else if value.Valid {
				m.airport_metars = new(int)
				*m.airport_metars = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAirport queries the "airport" edge of the Metar entity.
func (m *Metar) QueryAirport() *AirportQuery {
	return (&MetarClient{config: m.config}).QueryAirport(m)
}

// QuerySkyConditions queries the "sky_conditions" edge of the Metar entity.
func (m *Metar) QuerySkyConditions() *SkyConditionQuery {
	return (&MetarClient{config: m.config}).QuerySkyConditions(m)
}

// Update returns a builder for updating this Metar.
// Note that you need to call Metar.Unwrap() before calling this method if this Metar
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metar) Update() *MetarUpdateOne {
	return (&MetarClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Metar entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Metar) Unwrap() *Metar {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metar is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metar) String() string {
	var builder strings.Builder
	builder.WriteString("Metar(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("station_id=")
	builder.WriteString(m.StationID)
	builder.WriteString(", ")
	builder.WriteString("raw_text=")
	builder.WriteString(m.RawText)
	builder.WriteString(", ")
	builder.WriteString("observation_time=")
	builder.WriteString(m.ObservationTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := m.Latitude; v != nil {
		builder.WriteString("latitude=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Longitude; v != nil {
		builder.WriteString("longitude=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Elevation; v != nil {
		builder.WriteString("elevation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("temperature=")
	builder.WriteString(fmt.Sprintf("%v", m.Temperature))
	builder.WriteString(", ")
	builder.WriteString("dewpoint=")
	builder.WriteString(fmt.Sprintf("%v", m.Dewpoint))
	builder.WriteString(", ")
	builder.WriteString("wind_speed=")
	builder.WriteString(fmt.Sprintf("%v", m.WindSpeed))
	builder.WriteString(", ")
	builder.WriteString("wind_gust=")
	builder.WriteString(fmt.Sprintf("%v", m.WindGust))
	builder.WriteString(", ")
	builder.WriteString("wind_direction=")
	builder.WriteString(fmt.Sprintf("%v", m.WindDirection))
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", m.Visibility))
	builder.WriteString(", ")
	builder.WriteString("altimeter=")
	builder.WriteString(fmt.Sprintf("%v", m.Altimeter))
	builder.WriteString(", ")
	if v := m.PresentWeather; v != nil {
		builder.WriteString("present_weather=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.FlightCategory; v != nil {
		builder.WriteString("flight_category=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.QualityControlCorrected; v != nil {
		builder.WriteString("quality_control_corrected=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("quality_control_auto_station=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlAutoStation))
	builder.WriteString(", ")
	builder.WriteString("quality_control_maintenance_indicator_on=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlMaintenanceIndicatorOn))
	builder.WriteString(", ")
	builder.WriteString("quality_control_no_signal=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlNoSignal))
	builder.WriteString(", ")
	builder.WriteString("quality_control_lightning_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlLightningSensorOff))
	builder.WriteString(", ")
	builder.WriteString("quality_control_freezing_rain_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlFreezingRainSensorOff))
	builder.WriteString(", ")
	builder.WriteString("quality_control_present_weather_sensor_off=")
	builder.WriteString(fmt.Sprintf("%v", m.QualityControlPresentWeatherSensorOff))
	builder.WriteString(", ")
	if v := m.SeaLevelPressure; v != nil {
		builder.WriteString("sea_level_pressure=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.PressureTendency; v != nil {
		builder.WriteString("pressure_tendency=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.MaxTemp6; v != nil {
		builder.WriteString("max_temp_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.MinTemp6; v != nil {
		builder.WriteString("min_temp_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.MaxTemp24; v != nil {
		builder.WriteString("max_temp_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.MinTemp24; v != nil {
		builder.WriteString("min_temp_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Precipitation; v != nil {
		builder.WriteString("precipitation=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Precipitation3; v != nil {
		builder.WriteString("precipitation_3=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Precipitation6; v != nil {
		builder.WriteString("precipitation_6=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Precipitation24; v != nil {
		builder.WriteString("precipitation_24=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.SnowDepth; v != nil {
		builder.WriteString("snow_depth=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.VertVis; v != nil {
		builder.WriteString("vert_vis=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("metar_type=")
	builder.WriteString(fmt.Sprintf("%v", m.MetarType))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(m.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSkyConditions returns the SkyConditions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Metar) NamedSkyConditions(name string) ([]*SkyCondition, error) {
	if m.Edges.namedSkyConditions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedSkyConditions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Metar) appendNamedSkyConditions(name string, edges ...*SkyCondition) {
	if m.Edges.namedSkyConditions == nil {
		m.Edges.namedSkyConditions = make(map[string][]*SkyCondition)
	}
	if len(edges) == 0 {
		m.Edges.namedSkyConditions[name] = []*SkyCondition{}
	} else {
		m.Edges.namedSkyConditions[name] = append(m.Edges.namedSkyConditions[name], edges...)
	}
}

// Metars is a parsable slice of Metar.
type Metars []*Metar

func (m Metars) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}

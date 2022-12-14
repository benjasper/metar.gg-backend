// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"metar.gg/ent/airport"
	"metar.gg/ent/country"
	"metar.gg/ent/forecast"
	"metar.gg/ent/frequency"
	"metar.gg/ent/icingcondition"
	"metar.gg/ent/metar"
	"metar.gg/ent/region"
	"metar.gg/ent/runway"
	"metar.gg/ent/skycondition"
	"metar.gg/ent/taf"
	"metar.gg/ent/temperaturedata"
	"metar.gg/ent/turbulencecondition"
	"metar.gg/ent/weatherstation"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Airport) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Airport",
		Fields: make([]*Field, 19),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(a.ImportID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "import_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.LastUpdated); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "last_updated",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.IcaoCode); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "icao_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.IataCode); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "iata_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Identifier); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Type); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "airport.Type",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Importance); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "importance",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Name); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Latitude); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "float64",
		Name:  "latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Longitude); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "float64",
		Name:  "longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Timezone); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "timezone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Elevation); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "int",
		Name:  "elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Municipality); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "municipality",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.ScheduledService); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "bool",
		Name:  "scheduled_service",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.GpsCode); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "string",
		Name:  "gps_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.LocalCode); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "string",
		Name:  "local_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Website); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "string",
		Name:  "website",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Wikipedia); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "string",
		Name:  "wikipedia",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Keywords); err != nil {
		return nil, err
	}
	node.Fields[18] = &Field{
		Type:  "[]string",
		Name:  "keywords",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Region",
		Name: "region",
	}
	err = a.QueryRegion().
		Select(region.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Country",
		Name: "country",
	}
	err = a.QueryCountry().
		Select(country.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Frequency",
		Name: "frequencies",
	}
	err = a.QueryFrequencies().
		Select(frequency.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "WeatherStation",
		Name: "station",
	}
	err = a.QueryStation().
		Select(weatherstation.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Country) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Country",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(c.ImportID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "import_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.LastUpdated); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "last_updated",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Code); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Continent); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "country.Continent",
		Name:  "continent",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.WikipediaLink); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "wikipedia_link",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Keywords); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "[]string",
		Name:  "keywords",
		Value: string(buf),
	}
	return node, nil
}

func (f *Forecast) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Forecast",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(f.FromTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "from_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.ToTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "to_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.ChangeIndicator); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "forecast.ChangeIndicator",
		Name:  "change_indicator",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.ChangeTime); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "change_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.ChangeProbability); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "change_probability",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.WindDirection); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "wind_direction",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.WindShearDirection); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "wind_shear_direction",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Weather); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "weather",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.NotDecoded); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "not_decoded",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "SkyCondition",
		Name: "sky_conditions",
	}
	err = f.QuerySkyConditions().
		Select(skycondition.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "TurbulenceCondition",
		Name: "turbulence_conditions",
	}
	err = f.QueryTurbulenceConditions().
		Select(turbulencecondition.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "IcingCondition",
		Name: "icing_conditions",
	}
	err = f.QueryIcingConditions().
		Select(icingcondition.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "TemperatureData",
		Name: "temperature_data",
	}
	err = f.QueryTemperatureData().
		Select(temperaturedata.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (f *Frequency) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Frequency",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(f.ImportID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "import_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.LastUpdated); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "last_updated",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Type); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Frequency); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "float64",
		Name:  "frequency",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Airport",
		Name: "airport",
	}
	err = f.QueryAirport().
		Select(airport.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ic *IcingCondition) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ic.ID,
		Type:   "IcingCondition",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ic.Intensity); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "intensity",
		Value: string(buf),
	}
	return node, nil
}

func (m *Metar) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     m.ID,
		Type:   "Metar",
		Fields: make([]*Field, 23),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(m.RawText); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "raw_text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.ObservationTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "observation_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.ImportTime); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "import_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.NextImportTimePrediction); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "next_import_time_prediction",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.WindDirection); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "wind_direction",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.PresentWeather); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "present_weather",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.FlightCategory); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "metar.FlightCategory",
		Name:  "flight_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlCorrected); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "bool",
		Name:  "quality_control_corrected",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlAutoStation); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "bool",
		Name:  "quality_control_auto_station",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlMaintenanceIndicatorOn); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "bool",
		Name:  "quality_control_maintenance_indicator_on",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlNoSignal); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "bool",
		Name:  "quality_control_no_signal",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlLightningSensorOff); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "bool",
		Name:  "quality_control_lightning_sensor_off",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlFreezingRainSensorOff); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "bool",
		Name:  "quality_control_freezing_rain_sensor_off",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.QualityControlPresentWeatherSensorOff); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "bool",
		Name:  "quality_control_present_weather_sensor_off",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.MaxTemp6); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "float64",
		Name:  "max_temp_6",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.MinTemp6); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "float64",
		Name:  "min_temp_6",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.MaxTemp24); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "float64",
		Name:  "max_temp_24",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.MinTemp24); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "float64",
		Name:  "min_temp_24",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Precipitation); err != nil {
		return nil, err
	}
	node.Fields[18] = &Field{
		Type:  "float64",
		Name:  "precipitation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Precipitation3); err != nil {
		return nil, err
	}
	node.Fields[19] = &Field{
		Type:  "float64",
		Name:  "precipitation_3",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Precipitation6); err != nil {
		return nil, err
	}
	node.Fields[20] = &Field{
		Type:  "float64",
		Name:  "precipitation_6",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Precipitation24); err != nil {
		return nil, err
	}
	node.Fields[21] = &Field{
		Type:  "float64",
		Name:  "precipitation_24",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.MetarType); err != nil {
		return nil, err
	}
	node.Fields[22] = &Field{
		Type:  "metar.MetarType",
		Name:  "metar_type",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WeatherStation",
		Name: "station",
	}
	err = m.QueryStation().
		Select(weatherstation.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "SkyCondition",
		Name: "sky_conditions",
	}
	err = m.QuerySkyConditions().
		Select(skycondition.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (r *Region) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Region",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(r.ImportID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "import_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LastUpdated); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "last_updated",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Code); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LocalCode); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "local_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.WikipediaLink); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "wikipedia_link",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Keywords); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "[]string",
		Name:  "keywords",
		Value: string(buf),
	}
	return node, nil
}

func (r *Runway) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Runway",
		Fields: make([]*Field, 17),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(r.ImportID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "import_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LastUpdated); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "last_updated",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Surface); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "surface",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Lighted); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "bool",
		Name:  "lighted",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Closed); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "closed",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayIdentifier); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "low_runway_identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayLatitude); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "float64",
		Name:  "low_runway_latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayLongitude); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "float64",
		Name:  "low_runway_longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayElevation); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "int",
		Name:  "low_runway_elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayHeading); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "float64",
		Name:  "low_runway_heading",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayDisplacedThreshold); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "int",
		Name:  "low_runway_displaced_threshold",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayIdentifier); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "high_runway_identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayLatitude); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "float64",
		Name:  "high_runway_latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayLongitude); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "float64",
		Name:  "high_runway_longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayElevation); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "int",
		Name:  "high_runway_elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayHeading); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "float64",
		Name:  "high_runway_heading",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayDisplacedThreshold); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "int",
		Name:  "high_runway_displaced_threshold",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Airport",
		Name: "airport",
	}
	err = r.QueryAirport().
		Select(airport.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (sc *SkyCondition) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     sc.ID,
		Type:   "SkyCondition",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(sc.SkyCover); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "skycondition.SkyCover",
		Name:  "sky_cover",
		Value: string(buf),
	}
	if buf, err = json.Marshal(sc.CloudType); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "skycondition.CloudType",
		Name:  "cloud_type",
		Value: string(buf),
	}
	return node, nil
}

func (t *Taf) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Taf",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(t.RawText); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "raw_text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.IssueTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "issue_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.ImportTime); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "import_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.BulletinTime); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "bulletin_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.ValidFromTime); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "valid_from_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.ValidToTime); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "valid_to_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Remarks); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "remarks",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WeatherStation",
		Name: "station",
	}
	err = t.QueryStation().
		Select(weatherstation.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Forecast",
		Name: "forecast",
	}
	err = t.QueryForecast().
		Select(forecast.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (td *TemperatureData) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     td.ID,
		Type:   "TemperatureData",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(td.ValidTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "valid_time",
		Value: string(buf),
	}
	return node, nil
}

func (tc *TurbulenceCondition) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     tc.ID,
		Type:   "TurbulenceCondition",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(tc.Intensity); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "intensity",
		Value: string(buf),
	}
	return node, nil
}

func (ws *WeatherStation) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ws.ID,
		Type:   "WeatherStation",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ws.StationID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "station_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.Latitude); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "float64",
		Name:  "latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.Longitude); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "float64",
		Name:  "longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.Elevation); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "elevation",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Airport",
		Name: "airport",
	}
	err = ws.QueryAirport().
		Select(airport.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case airport.Table:
		query := c.Airport.Query().
			Where(airport.ID(id))
		query, err := query.CollectFields(ctx, "Airport")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case country.Table:
		query := c.Country.Query().
			Where(country.ID(id))
		query, err := query.CollectFields(ctx, "Country")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case forecast.Table:
		query := c.Forecast.Query().
			Where(forecast.ID(id))
		query, err := query.CollectFields(ctx, "Forecast")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case frequency.Table:
		query := c.Frequency.Query().
			Where(frequency.ID(id))
		query, err := query.CollectFields(ctx, "Frequency")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case icingcondition.Table:
		query := c.IcingCondition.Query().
			Where(icingcondition.ID(id))
		query, err := query.CollectFields(ctx, "IcingCondition")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case metar.Table:
		query := c.Metar.Query().
			Where(metar.ID(id))
		query, err := query.CollectFields(ctx, "Metar")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case region.Table:
		query := c.Region.Query().
			Where(region.ID(id))
		query, err := query.CollectFields(ctx, "Region")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case runway.Table:
		query := c.Runway.Query().
			Where(runway.ID(id))
		query, err := query.CollectFields(ctx, "Runway")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case skycondition.Table:
		query := c.SkyCondition.Query().
			Where(skycondition.ID(id))
		query, err := query.CollectFields(ctx, "SkyCondition")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case taf.Table:
		query := c.Taf.Query().
			Where(taf.ID(id))
		query, err := query.CollectFields(ctx, "Taf")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case temperaturedata.Table:
		query := c.TemperatureData.Query().
			Where(temperaturedata.ID(id))
		query, err := query.CollectFields(ctx, "TemperatureData")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case turbulencecondition.Table:
		query := c.TurbulenceCondition.Query().
			Where(turbulencecondition.ID(id))
		query, err := query.CollectFields(ctx, "TurbulenceCondition")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case weatherstation.Table:
		query := c.WeatherStation.Query().
			Where(weatherstation.ID(id))
		query, err := query.CollectFields(ctx, "WeatherStation")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case airport.Table:
		query := c.Airport.Query().
			Where(airport.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Airport")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case country.Table:
		query := c.Country.Query().
			Where(country.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Country")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case forecast.Table:
		query := c.Forecast.Query().
			Where(forecast.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Forecast")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case frequency.Table:
		query := c.Frequency.Query().
			Where(frequency.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Frequency")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case icingcondition.Table:
		query := c.IcingCondition.Query().
			Where(icingcondition.IDIn(ids...))
		query, err := query.CollectFields(ctx, "IcingCondition")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case metar.Table:
		query := c.Metar.Query().
			Where(metar.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Metar")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case region.Table:
		query := c.Region.Query().
			Where(region.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Region")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case runway.Table:
		query := c.Runway.Query().
			Where(runway.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Runway")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case skycondition.Table:
		query := c.SkyCondition.Query().
			Where(skycondition.IDIn(ids...))
		query, err := query.CollectFields(ctx, "SkyCondition")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case taf.Table:
		query := c.Taf.Query().
			Where(taf.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Taf")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case temperaturedata.Table:
		query := c.TemperatureData.Query().
			Where(temperaturedata.IDIn(ids...))
		query, err := query.CollectFields(ctx, "TemperatureData")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case turbulencecondition.Table:
		query := c.TurbulenceCondition.Query().
			Where(turbulencecondition.IDIn(ids...))
		query, err := query.CollectFields(ctx, "TurbulenceCondition")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case weatherstation.Table:
		query := c.WeatherStation.Query().
			Where(weatherstation.IDIn(ids...))
		query, err := query.CollectFields(ctx, "WeatherStation")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

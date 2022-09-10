package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Taf holds the schema definition for the Metar entity.
type Taf struct {
	ent.Schema
}

// Fields of the Metar.
func (Taf) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entgql.Skip()),
		field.String("raw_text").Comment("The raw TAF text."),
		field.Time("issue_time").Comment("The time the TAF was issued."),
		field.Time("bulletin_time").Comment("TAF bulletin time."),
		field.Time("valid_from_time").Comment("The start time of the TAF validity period.").Annotations(entgql.OrderField("valid_from_time")),
		field.Time("valid_to_time").Comment("The end time of the TAF validity period."),
		field.String("remarks").Comment("Remarks."),

		field.Float("temperature").Comment("The temperature in Celsius."),
		field.Float("dewpoint").Comment("The dewpoint in Celsius."),
		field.Int("wind_speed").Comment("The wind speed in knots, or 0 if calm."),
		field.Int("wind_gust").Comment("The wind gust in knots."),
		field.Int("wind_direction").Comment("The wind direction in degrees, or 0 if calm."),
		field.Float("visibility").Comment("The visibility in statute miles."),
		field.Float("altimeter").Comment("The altimeter setting in inches of mercury."),
		field.Enum("flight_category").Optional().Nillable().Values("VFR", "MVFR", "IFR", "LIFR"),
		field.Bool("quality_control_corrected").Optional().Nillable().Comment("Quality control corrected."),
		field.Bool("quality_control_auto_station").Comment("Whether it's an automated station, of one of the following types A01|A01A|A02|A02A|AOA|AWOS."),
		field.Bool("quality_control_maintenance_indicator_on").Comment("Maintenance check indicator - maintenance is needed."),
		field.Bool("quality_control_no_signal").Comment("No signal."),
		field.Bool("quality_control_lightning_sensor_off").Comment("Whether Lightning sensor is off."),
		field.Bool("quality_control_freezing_rain_sensor_off").Comment("Whether Freezing rain sensor is off."),
		field.Bool("quality_control_present_weather_sensor_off").Comment("Whether Present weather sensor is off."),
		field.Float("sea_level_pressure").Optional().Nillable().Comment("The sea level pressure in hectopascal.s"),
		field.Float("pressure_tendency").Optional().Nillable().Comment("The pressur_6e tendency in hectopascals."),
		field.Float("max_temp_6").Optional().Nillable().Comment("The maximum air temperature in Celsius from the past 6 hours."),
		field.Float("min_temp_6").Optional().Nillable().Comment("The minimum air temperature in Celsius from the past 6 hours."),
		field.Float("max_temp_24").Optional().Nillable().Comment("The maximum air temperature in Celsius from the past 24 hours."),
		field.Float("min_temp_24").Optional().Nillable().Comment("The minimum air temperature in Celsius from the past 24 hours."),
		field.Float("precipitation").Optional().Nillable().Comment("The precipitation in inches from since the last observation. 0.0005 in = trace precipitation"),
		field.Float("precipitation_3").Optional().Nillable().Comment("The precipitation in inches from the past 3 hours. 0.0005 in = trace precipitation"),
		field.Float("precipitation_6").Optional().Nillable().Comment("The precipitation in inches from the past 6 hours. 0.0005 in = trace precipitation"),
		field.Float("precipitation_24").Optional().Nillable().Comment("The precipitation in inches from the past 24 hours. 0.0005 in = trace precipitation"),
		field.Float("snow_depth").Optional().Nillable().Comment("The snow depth in inches."),
		field.Float("vert_vis").Optional().Nillable().Comment("The vertical visibility in feet."),
		field.Enum("metar_type").Values("TAF", "SPECI").Comment("The type of TAF."),
		field.String("hash").Annotations(entgql.Skip()),
	}
}

// Edges of the Metar.
func (Taf) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("station", Station.Type).Ref("tafs").Unique().Required().Comment("The station that issued this taf."),
		edge.To("sky_conditions", SkyCondition.Type).Comment("The sky conditions.").Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

// Indexes of the Metar.
func (Taf) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("issue_time"),
	}
}

// Mixin of the Metar.
func (Taf) Mixin() []ent.Mixin {
	return nil
}

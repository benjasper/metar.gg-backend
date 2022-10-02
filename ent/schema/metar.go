package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Metar holds the schema definition for the Metar entity.
type Metar struct {
	ent.Schema
}

// Fields of the Metar.
func (Metar) Fields() []ent.Field {
	return []ent.Field{
		field.Text("raw_text").Comment("The raw METAR text."),
		field.Time("observation_time").Comment("The time the METAR was observed."),
		field.Time("import_time").Comment("The time the METAR was imported.").Default(time.Now),
		field.Float("temperature").Comment("The temperature in Celsius.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Float("dewpoint").Comment("The dewpoint in Celsius.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Int("wind_speed").Comment("The wind speed in knots, or 0 if calm.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Int("wind_gust").Comment("The wind gust in knots.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Int("wind_direction").Comment("The wind direction in degrees, or 0 if calm."),
		field.Float("visibility").Comment("The visibility in statute miles.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Float("altimeter").Comment("The altimeter setting in inches of mercury.").Annotations(entgql.Skip(entgql.SkipType)),
		field.String("present_weather").Optional().Nillable().Comment("The present weather string."),
		field.Enum("flight_category").Optional().Nillable().Values("VFR", "MVFR", "IFR", "LIFR"),
		field.Bool("quality_control_corrected").Optional().Nillable().Comment("Quality control corrected."),
		field.Bool("quality_control_auto_station").Comment("Whether it's an automated station, of one of the following types A01|A01A|A02|A02A|AOA|AWOS."),
		field.Bool("quality_control_maintenance_indicator_on").Comment("Maintenance check indicator - maintenance is needed."),
		field.Bool("quality_control_no_signal").Comment("No signal."),
		field.Bool("quality_control_lightning_sensor_off").Comment("Whether Lightning sensor is off."),
		field.Bool("quality_control_freezing_rain_sensor_off").Comment("Whether Freezing rain sensor is off."),
		field.Bool("quality_control_present_weather_sensor_off").Comment("Whether Present weather sensor is off."),
		field.Float("sea_level_pressure").Optional().Nillable().Comment("The sea level pressure in hectopascals.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Float("pressure_tendency").Optional().Nillable().Comment("The pressure tendency in hectopascals.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Float("max_temp_6").Optional().Nillable().Comment("The maximum air temperature in Celsius from the past 6 hours."),
		field.Float("min_temp_6").Optional().Nillable().Comment("The minimum air temperature in Celsius from the past 6 hours."),
		field.Float("max_temp_24").Optional().Nillable().Comment("The maximum air temperature in Celsius from the past 24 hours."),
		field.Float("min_temp_24").Optional().Nillable().Comment("The minimum air temperature in Celsius from the past 24 hours."),
		field.Float("precipitation").Optional().Nillable().Comment("The precipitation in inches from since the last observation. 0.0005 in = trace precipitation."),
		field.Float("precipitation_3").Optional().Nillable().Comment("The precipitation in inches from the past 3 hours. 0.0005 in = trace precipitation."),
		field.Float("precipitation_6").Optional().Nillable().Comment("The precipitation in inches from the past 6 hours. 0.0005 in = trace precipitation."),
		field.Float("precipitation_24").Optional().Nillable().Comment("The precipitation in inches from the past 24 hours. 0.0005 in = trace precipitation."),
		field.Float("snow_depth").Optional().Nillable().Comment("The snow depth in inches.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Float("vert_vis").Optional().Nillable().Comment("The vertical visibility in feet.").Annotations(entgql.Skip(entgql.SkipType)),
		field.Enum("metar_type").Values("METAR", "SPECI").Comment("The type of METAR."),
		field.String("hash").Annotations(entgql.Skip()),
	}
}

// Edges of the Metar.
func (Metar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("station", WeatherStation.Type).Ref("metars").Unique().Required().Comment("The station that provided the METAR."),
		edge.To("sky_conditions", SkyCondition.Type).Comment("The sky conditions.").Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

// Indexes of the Metar.
func (Metar) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("observation_time"),
		index.Fields("hash"),
	}
}

// Mixin of the Metar.
func (Metar) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

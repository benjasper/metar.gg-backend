// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AirportsColumns holds the columns for the "airports" table.
	AirportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"large_airport", "medium_airport", "small_airport", "closed_airport", "heliport", "seaplane_base"}},
		{Name: "name", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "elevation", Type: field.TypeInt, Nullable: true},
		{Name: "continent", Type: field.TypeEnum, Enums: []string{"AF", "AN", "AS", "EU", "NA", "SA", "OC"}},
		{Name: "country", Type: field.TypeString},
		{Name: "region", Type: field.TypeString},
		{Name: "municipality", Type: field.TypeString, Nullable: true},
		{Name: "scheduled_service", Type: field.TypeBool},
		{Name: "gps_code", Type: field.TypeString, Nullable: true},
		{Name: "iata_code", Type: field.TypeString, Nullable: true},
		{Name: "local_code", Type: field.TypeString, Nullable: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "wikipedia", Type: field.TypeString, Nullable: true},
		{Name: "keywords", Type: field.TypeJSON},
	}
	// AirportsTable holds the schema information for the "airports" table.
	AirportsTable = &schema.Table{
		Name:       "airports",
		Columns:    AirportsColumns,
		PrimaryKey: []*schema.Column{AirportsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "airport_hash",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[1]},
			},
			{
				Name:    "airport_identifier",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[4]},
			},
		},
	}
	// FrequenciesColumns holds the columns for the "frequencies" table.
	FrequenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "frequency", Type: field.TypeFloat64},
		{Name: "airport_frequencies", Type: field.TypeInt, Nullable: true},
	}
	// FrequenciesTable holds the schema information for the "frequencies" table.
	FrequenciesTable = &schema.Table{
		Name:       "frequencies",
		Columns:    FrequenciesColumns,
		PrimaryKey: []*schema.Column{FrequenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "frequencies_airports_frequencies",
				Columns:    []*schema.Column{FrequenciesColumns[7]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "frequency_hash",
				Unique:  false,
				Columns: []*schema.Column{FrequenciesColumns[1]},
			},
		},
	}
	// MetarsColumns holds the columns for the "metars" table.
	MetarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "raw_text", Type: field.TypeString},
		{Name: "observation_time", Type: field.TypeTime},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "elevation", Type: field.TypeFloat64, Nullable: true},
		{Name: "temperature", Type: field.TypeFloat64},
		{Name: "dewpoint", Type: field.TypeFloat64},
		{Name: "wind_speed", Type: field.TypeInt},
		{Name: "wind_gust", Type: field.TypeInt},
		{Name: "wind_direction", Type: field.TypeInt},
		{Name: "visibility", Type: field.TypeFloat64},
		{Name: "altimeter", Type: field.TypeFloat64},
		{Name: "present_weather", Type: field.TypeString, Nullable: true},
		{Name: "flight_category", Type: field.TypeEnum, Nullable: true, Enums: []string{"VFR", "MVFR", "IFR", "LIFR"}},
		{Name: "quality_control_corrected", Type: field.TypeBool, Nullable: true},
		{Name: "quality_control_auto_station", Type: field.TypeBool},
		{Name: "quality_control_maintenance_indicator_on", Type: field.TypeBool},
		{Name: "quality_control_no_signal", Type: field.TypeBool},
		{Name: "quality_control_lightning_sensor_off", Type: field.TypeBool},
		{Name: "quality_control_freezing_rain_sensor_off", Type: field.TypeBool},
		{Name: "quality_control_present_weather_sensor_off", Type: field.TypeBool},
		{Name: "sea_level_pressure", Type: field.TypeFloat64, Nullable: true},
		{Name: "pressure_tendency", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temp_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "min_temp_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temp_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "min_temp_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_3", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "snow_depth", Type: field.TypeFloat64, Nullable: true},
		{Name: "vert_vis", Type: field.TypeFloat64, Nullable: true},
		{Name: "metar_type", Type: field.TypeEnum, Enums: []string{"METAR", "SPECI"}},
		{Name: "hash", Type: field.TypeString},
		{Name: "station_metars", Type: field.TypeInt},
	}
	// MetarsTable holds the schema information for the "metars" table.
	MetarsTable = &schema.Table{
		Name:       "metars",
		Columns:    MetarsColumns,
		PrimaryKey: []*schema.Column{MetarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metars_stations_metars",
				Columns:    []*schema.Column{MetarsColumns[36]},
				RefColumns: []*schema.Column{StationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "metar_observation_time",
				Unique:  false,
				Columns: []*schema.Column{MetarsColumns[2]},
			},
		},
	}
	// RunwaysColumns holds the columns for the "runways" table.
	RunwaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "length", Type: field.TypeInt},
		{Name: "width", Type: field.TypeInt},
		{Name: "surface", Type: field.TypeString, Nullable: true},
		{Name: "lighted", Type: field.TypeBool},
		{Name: "closed", Type: field.TypeBool},
		{Name: "low_runway_identifier", Type: field.TypeString},
		{Name: "low_runway_latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "low_runway_longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "low_runway_elevation", Type: field.TypeInt, Nullable: true},
		{Name: "low_runway_heading", Type: field.TypeFloat64, Nullable: true},
		{Name: "low_runway_displaced_threshold", Type: field.TypeInt, Nullable: true},
		{Name: "high_runway_identifier", Type: field.TypeString},
		{Name: "high_runway_latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "high_runway_longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "high_runway_elevation", Type: field.TypeInt, Nullable: true},
		{Name: "high_runway_heading", Type: field.TypeFloat64, Nullable: true},
		{Name: "high_runway_displaced_threshold", Type: field.TypeInt, Nullable: true},
		{Name: "airport_runways", Type: field.TypeInt, Nullable: true},
	}
	// RunwaysTable holds the schema information for the "runways" table.
	RunwaysTable = &schema.Table{
		Name:       "runways",
		Columns:    RunwaysColumns,
		PrimaryKey: []*schema.Column{RunwaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "runways_airports_runways",
				Columns:    []*schema.Column{RunwaysColumns[21]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "runway_hash",
				Unique:  false,
				Columns: []*schema.Column{RunwaysColumns[1]},
			},
		},
	}
	// SkyConditionsColumns holds the columns for the "sky_conditions" table.
	SkyConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sky_cover", Type: field.TypeEnum, Enums: []string{"SKC", "FEW", "SCT", "CLR", "BKN", "OVC", "OVX", "CAVOK"}},
		{Name: "cloud_base", Type: field.TypeInt, Nullable: true},
		{Name: "metar_sky_conditions", Type: field.TypeInt},
		{Name: "taf_sky_conditions", Type: field.TypeInt, Nullable: true},
	}
	// SkyConditionsTable holds the schema information for the "sky_conditions" table.
	SkyConditionsTable = &schema.Table{
		Name:       "sky_conditions",
		Columns:    SkyConditionsColumns,
		PrimaryKey: []*schema.Column{SkyConditionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sky_conditions_metars_sky_conditions",
				Columns:    []*schema.Column{SkyConditionsColumns[3]},
				RefColumns: []*schema.Column{MetarsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sky_conditions_tafs_sky_conditions",
				Columns:    []*schema.Column{SkyConditionsColumns[4]},
				RefColumns: []*schema.Column{TafsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StationsColumns holds the columns for the "stations" table.
	StationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "station_id", Type: field.TypeString, Unique: true},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "elevation", Type: field.TypeFloat64, Nullable: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "airport_station", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// StationsTable holds the schema information for the "stations" table.
	StationsTable = &schema.Table{
		Name:       "stations",
		Columns:    StationsColumns,
		PrimaryKey: []*schema.Column{StationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stations_airports_station",
				Columns:    []*schema.Column{StationsColumns[6]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "station_station_id",
				Unique:  true,
				Columns: []*schema.Column{StationsColumns[1]},
			},
			{
				Name:    "station_latitude",
				Unique:  false,
				Columns: []*schema.Column{StationsColumns[2]},
			},
			{
				Name:    "station_longitude",
				Unique:  false,
				Columns: []*schema.Column{StationsColumns[3]},
			},
		},
	}
	// TafsColumns holds the columns for the "tafs" table.
	TafsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "raw_text", Type: field.TypeString},
		{Name: "issue_time", Type: field.TypeTime},
		{Name: "bulletin_time", Type: field.TypeTime},
		{Name: "valid_from_time", Type: field.TypeTime},
		{Name: "valid_to_time", Type: field.TypeTime},
		{Name: "remarks", Type: field.TypeString},
		{Name: "temperature", Type: field.TypeFloat64},
		{Name: "dewpoint", Type: field.TypeFloat64},
		{Name: "wind_speed", Type: field.TypeInt},
		{Name: "wind_gust", Type: field.TypeInt},
		{Name: "wind_direction", Type: field.TypeInt},
		{Name: "visibility", Type: field.TypeFloat64},
		{Name: "altimeter", Type: field.TypeFloat64},
		{Name: "flight_category", Type: field.TypeEnum, Nullable: true, Enums: []string{"VFR", "MVFR", "IFR", "LIFR"}},
		{Name: "quality_control_corrected", Type: field.TypeBool, Nullable: true},
		{Name: "quality_control_auto_station", Type: field.TypeBool},
		{Name: "quality_control_maintenance_indicator_on", Type: field.TypeBool},
		{Name: "quality_control_no_signal", Type: field.TypeBool},
		{Name: "quality_control_lightning_sensor_off", Type: field.TypeBool},
		{Name: "quality_control_freezing_rain_sensor_off", Type: field.TypeBool},
		{Name: "quality_control_present_weather_sensor_off", Type: field.TypeBool},
		{Name: "sea_level_pressure", Type: field.TypeFloat64, Nullable: true},
		{Name: "pressure_tendency", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temp_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "min_temp_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temp_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "min_temp_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_3", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_6", Type: field.TypeFloat64, Nullable: true},
		{Name: "precipitation_24", Type: field.TypeFloat64, Nullable: true},
		{Name: "snow_depth", Type: field.TypeFloat64, Nullable: true},
		{Name: "vert_vis", Type: field.TypeFloat64, Nullable: true},
		{Name: "metar_type", Type: field.TypeEnum, Enums: []string{"TAF", "SPECI"}},
		{Name: "hash", Type: field.TypeString},
		{Name: "station_tafs", Type: field.TypeInt},
	}
	// TafsTable holds the schema information for the "tafs" table.
	TafsTable = &schema.Table{
		Name:       "tafs",
		Columns:    TafsColumns,
		PrimaryKey: []*schema.Column{TafsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tafs_stations_tafs",
				Columns:    []*schema.Column{TafsColumns[36]},
				RefColumns: []*schema.Column{StationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "taf_issue_time",
				Unique:  false,
				Columns: []*schema.Column{TafsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AirportsTable,
		FrequenciesTable,
		MetarsTable,
		RunwaysTable,
		SkyConditionsTable,
		StationsTable,
		TafsTable,
	}
)

func init() {
	FrequenciesTable.ForeignKeys[0].RefTable = AirportsTable
	MetarsTable.ForeignKeys[0].RefTable = StationsTable
	RunwaysTable.ForeignKeys[0].RefTable = AirportsTable
	SkyConditionsTable.ForeignKeys[0].RefTable = MetarsTable
	SkyConditionsTable.ForeignKeys[1].RefTable = TafsTable
	StationsTable.ForeignKeys[0].RefTable = AirportsTable
	TafsTable.ForeignKeys[0].RefTable = StationsTable
}

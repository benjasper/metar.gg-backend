// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AirportsColumns holds the columns for the "airports" table.
	AirportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "import_id", Type: field.TypeInt},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "icao_code", Type: field.TypeString, Nullable: true, Size: 4},
		{Name: "iata_code", Type: field.TypeString, Nullable: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"large_airport", "medium_airport", "small_airport", "closed_airport", "heliport", "seaplane_base"}},
		{Name: "importance", Type: field.TypeInt, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "timezone", Type: field.TypeString, Nullable: true},
		{Name: "elevation", Type: field.TypeInt, Nullable: true},
		{Name: "municipality", Type: field.TypeString, Nullable: true},
		{Name: "scheduled_service", Type: field.TypeBool},
		{Name: "gps_code", Type: field.TypeString, Nullable: true},
		{Name: "local_code", Type: field.TypeString, Nullable: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "wikipedia", Type: field.TypeString, Nullable: true},
		{Name: "keywords", Type: field.TypeJSON},
		{Name: "country_airports", Type: field.TypeUUID, Nullable: true},
		{Name: "region_airports", Type: field.TypeUUID, Nullable: true},
	}
	// AirportsTable holds the schema information for the "airports" table.
	AirportsTable = &schema.Table{
		Name:       "airports",
		Columns:    AirportsColumns,
		PrimaryKey: []*schema.Column{AirportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "airports_countries_airports",
				Columns:    []*schema.Column{AirportsColumns[22]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "airports_regions_airports",
				Columns:    []*schema.Column{AirportsColumns[23]},
				RefColumns: []*schema.Column{RegionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "airport_hash",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[2]},
			},
			{
				Name:    "airport_import_id",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[1]},
			},
			{
				Name:    "fulltext",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[10], AirportsColumns[15], AirportsColumns[5], AirportsColumns[6], AirportsColumns[18], AirportsColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"mysql": "FULLTEXT",
					},
				},
			},
			{
				Name:    "airport_identifier",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[7]},
			},
			{
				Name:    "airport_name",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[10]},
			},
			{
				Name:    "airport_municipality",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[15]},
			},
			{
				Name:    "airport_local_code",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[18]},
			},
			{
				Name:    "airport_icao_code",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[5]},
			},
			{
				Name:    "airport_iata_code",
				Unique:  false,
				Columns: []*schema.Column{AirportsColumns[6]},
			},
		},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "import_id", Type: field.TypeInt},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "continent", Type: field.TypeEnum, Enums: []string{"AF", "AN", "AS", "EU", "NA", "SA", "OC"}},
		{Name: "wikipedia_link", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeJSON},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "country_hash",
				Unique:  false,
				Columns: []*schema.Column{CountriesColumns[2]},
			},
			{
				Name:    "country_import_id",
				Unique:  false,
				Columns: []*schema.Column{CountriesColumns[1]},
			},
			{
				Name:    "country_name",
				Unique:  false,
				Columns: []*schema.Column{CountriesColumns[6]},
			},
		},
	}
	// ForecastsColumns holds the columns for the "forecasts" table.
	ForecastsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "from_time", Type: field.TypeTime},
		{Name: "to_time", Type: field.TypeTime},
		{Name: "change_indicator", Type: field.TypeEnum, Nullable: true, Enums: []string{"BECMG", "FM", "TEMPO", "PROB"}},
		{Name: "change_time", Type: field.TypeTime, Nullable: true},
		{Name: "change_probability", Type: field.TypeInt, Nullable: true},
		{Name: "wind_direction", Type: field.TypeInt, Nullable: true},
		{Name: "wind_direction_variable", Type: field.TypeBool},
		{Name: "wind_speed", Type: field.TypeInt, Nullable: true},
		{Name: "wind_gust", Type: field.TypeInt, Nullable: true},
		{Name: "wind_shear_height", Type: field.TypeInt, Nullable: true},
		{Name: "wind_shear_direction", Type: field.TypeInt, Nullable: true},
		{Name: "wind_shear_speed", Type: field.TypeInt, Nullable: true},
		{Name: "visibility_horizontal", Type: field.TypeFloat64, Nullable: true},
		{Name: "visibility_horizontal_is_more_than", Type: field.TypeBool},
		{Name: "visibility_vertical", Type: field.TypeInt, Nullable: true},
		{Name: "altimeter", Type: field.TypeFloat64, Nullable: true},
		{Name: "weather", Type: field.TypeString, Nullable: true},
		{Name: "not_decoded", Type: field.TypeString, Nullable: true},
		{Name: "taf_forecast", Type: field.TypeUUID, Nullable: true},
	}
	// ForecastsTable holds the schema information for the "forecasts" table.
	ForecastsTable = &schema.Table{
		Name:       "forecasts",
		Columns:    ForecastsColumns,
		PrimaryKey: []*schema.Column{ForecastsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "forecasts_tafs_forecast",
				Columns:    []*schema.Column{ForecastsColumns[19]},
				RefColumns: []*schema.Column{TafsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FrequenciesColumns holds the columns for the "frequencies" table.
	FrequenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "import_id", Type: field.TypeInt},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "frequency", Type: field.TypeFloat64},
		{Name: "airport_frequencies", Type: field.TypeUUID, Nullable: true},
	}
	// FrequenciesTable holds the schema information for the "frequencies" table.
	FrequenciesTable = &schema.Table{
		Name:       "frequencies",
		Columns:    FrequenciesColumns,
		PrimaryKey: []*schema.Column{FrequenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "frequencies_airports_frequencies",
				Columns:    []*schema.Column{FrequenciesColumns[8]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "frequency_hash",
				Unique:  false,
				Columns: []*schema.Column{FrequenciesColumns[2]},
			},
			{
				Name:    "frequency_import_id",
				Unique:  false,
				Columns: []*schema.Column{FrequenciesColumns[1]},
			},
		},
	}
	// IcingConditionsColumns holds the columns for the "icing_conditions" table.
	IcingConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "intensity", Type: field.TypeString},
		{Name: "min_altitude", Type: field.TypeInt, Nullable: true},
		{Name: "max_altitude", Type: field.TypeInt, Nullable: true},
		{Name: "forecast_icing_conditions", Type: field.TypeUUID, Nullable: true},
	}
	// IcingConditionsTable holds the schema information for the "icing_conditions" table.
	IcingConditionsTable = &schema.Table{
		Name:       "icing_conditions",
		Columns:    IcingConditionsColumns,
		PrimaryKey: []*schema.Column{IcingConditionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "icing_conditions_forecasts_icing_conditions",
				Columns:    []*schema.Column{IcingConditionsColumns[4]},
				RefColumns: []*schema.Column{ForecastsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MetarsColumns holds the columns for the "metars" table.
	MetarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "raw_text", Type: field.TypeString, Size: 2147483647},
		{Name: "observation_time", Type: field.TypeTime},
		{Name: "import_time", Type: field.TypeTime},
		{Name: "next_import_time_prediction", Type: field.TypeTime, Nullable: true},
		{Name: "temperature", Type: field.TypeFloat64, Nullable: true},
		{Name: "dewpoint", Type: field.TypeFloat64, Nullable: true},
		{Name: "wind_speed", Type: field.TypeInt, Nullable: true},
		{Name: "wind_gust", Type: field.TypeInt, Nullable: true},
		{Name: "wind_direction", Type: field.TypeInt, Nullable: true},
		{Name: "wind_direction_variable", Type: field.TypeBool},
		{Name: "visibility", Type: field.TypeFloat64, Nullable: true},
		{Name: "visibility_is_more_than", Type: field.TypeBool},
		{Name: "altimeter", Type: field.TypeFloat64, Nullable: true},
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
		{Name: "weather_station_metars", Type: field.TypeUUID},
	}
	// MetarsTable holds the schema information for the "metars" table.
	MetarsTable = &schema.Table{
		Name:       "metars",
		Columns:    MetarsColumns,
		PrimaryKey: []*schema.Column{MetarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metars_weather_stations_metars",
				Columns:    []*schema.Column{MetarsColumns[37]},
				RefColumns: []*schema.Column{WeatherStationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "metar_observation_time",
				Unique:  false,
				Columns: []*schema.Column{MetarsColumns[2]},
			},
			{
				Name:    "metar_hash",
				Unique:  false,
				Columns: []*schema.Column{MetarsColumns[36]},
			},
			{
				Name:    "metar_import_time",
				Unique:  false,
				Columns: []*schema.Column{MetarsColumns[3]},
			},
		},
	}
	// RegionsColumns holds the columns for the "regions" table.
	RegionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "import_id", Type: field.TypeInt},
		{Name: "hash", Type: field.TypeString},
		{Name: "import_flag", Type: field.TypeBool, Default: false},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString},
		{Name: "local_code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "wikipedia_link", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeJSON},
	}
	// RegionsTable holds the schema information for the "regions" table.
	RegionsTable = &schema.Table{
		Name:       "regions",
		Columns:    RegionsColumns,
		PrimaryKey: []*schema.Column{RegionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "region_hash",
				Unique:  false,
				Columns: []*schema.Column{RegionsColumns[2]},
			},
			{
				Name:    "region_import_id",
				Unique:  false,
				Columns: []*schema.Column{RegionsColumns[1]},
			},
			{
				Name:    "region_name",
				Unique:  false,
				Columns: []*schema.Column{RegionsColumns[7]},
			},
		},
	}
	// RunwaysColumns holds the columns for the "runways" table.
	RunwaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "import_id", Type: field.TypeInt},
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
		{Name: "airport_runways", Type: field.TypeUUID, Nullable: true},
	}
	// RunwaysTable holds the schema information for the "runways" table.
	RunwaysTable = &schema.Table{
		Name:       "runways",
		Columns:    RunwaysColumns,
		PrimaryKey: []*schema.Column{RunwaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "runways_airports_runways",
				Columns:    []*schema.Column{RunwaysColumns[22]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "runway_hash",
				Unique:  false,
				Columns: []*schema.Column{RunwaysColumns[2]},
			},
			{
				Name:    "runway_import_id",
				Unique:  false,
				Columns: []*schema.Column{RunwaysColumns[1]},
			},
		},
	}
	// SkyConditionsColumns holds the columns for the "sky_conditions" table.
	SkyConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sky_cover", Type: field.TypeEnum, Enums: []string{"SKC", "FEW", "SCT", "CLR", "NSC", "BKN", "OVC", "OVCX", "OVX", "CAVOK"}},
		{Name: "cloud_base", Type: field.TypeInt, Nullable: true},
		{Name: "cloud_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"CB", "CU", "TCU"}},
		{Name: "forecast_sky_conditions", Type: field.TypeUUID, Nullable: true},
		{Name: "metar_sky_conditions", Type: field.TypeUUID, Nullable: true},
	}
	// SkyConditionsTable holds the schema information for the "sky_conditions" table.
	SkyConditionsTable = &schema.Table{
		Name:       "sky_conditions",
		Columns:    SkyConditionsColumns,
		PrimaryKey: []*schema.Column{SkyConditionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sky_conditions_forecasts_sky_conditions",
				Columns:    []*schema.Column{SkyConditionsColumns[4]},
				RefColumns: []*schema.Column{ForecastsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sky_conditions_metars_sky_conditions",
				Columns:    []*schema.Column{SkyConditionsColumns[5]},
				RefColumns: []*schema.Column{MetarsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TafsColumns holds the columns for the "tafs" table.
	TafsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "raw_text", Type: field.TypeString, Size: 2147483647},
		{Name: "issue_time", Type: field.TypeTime},
		{Name: "import_time", Type: field.TypeTime},
		{Name: "bulletin_time", Type: field.TypeTime},
		{Name: "valid_from_time", Type: field.TypeTime},
		{Name: "valid_to_time", Type: field.TypeTime},
		{Name: "remarks", Type: field.TypeString},
		{Name: "hash", Type: field.TypeString},
		{Name: "weather_station_tafs", Type: field.TypeUUID},
	}
	// TafsTable holds the schema information for the "tafs" table.
	TafsTable = &schema.Table{
		Name:       "tafs",
		Columns:    TafsColumns,
		PrimaryKey: []*schema.Column{TafsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tafs_weather_stations_tafs",
				Columns:    []*schema.Column{TafsColumns[9]},
				RefColumns: []*schema.Column{WeatherStationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "taf_issue_time",
				Unique:  false,
				Columns: []*schema.Column{TafsColumns[2]},
			},
			{
				Name:    "taf_hash",
				Unique:  false,
				Columns: []*schema.Column{TafsColumns[8]},
			},
		},
	}
	// TemperatureDataColumns holds the columns for the "temperature_data" table.
	TemperatureDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "valid_time", Type: field.TypeTime},
		{Name: "temperature", Type: field.TypeFloat64},
		{Name: "min_temperature", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_temperature", Type: field.TypeFloat64, Nullable: true},
		{Name: "forecast_temperature_data", Type: field.TypeUUID, Nullable: true},
	}
	// TemperatureDataTable holds the schema information for the "temperature_data" table.
	TemperatureDataTable = &schema.Table{
		Name:       "temperature_data",
		Columns:    TemperatureDataColumns,
		PrimaryKey: []*schema.Column{TemperatureDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "temperature_data_forecasts_temperature_data",
				Columns:    []*schema.Column{TemperatureDataColumns[5]},
				RefColumns: []*schema.Column{ForecastsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TurbulenceConditionsColumns holds the columns for the "turbulence_conditions" table.
	TurbulenceConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "intensity", Type: field.TypeString},
		{Name: "min_altitude", Type: field.TypeInt},
		{Name: "max_altitude", Type: field.TypeInt},
		{Name: "forecast_turbulence_conditions", Type: field.TypeUUID, Nullable: true},
	}
	// TurbulenceConditionsTable holds the schema information for the "turbulence_conditions" table.
	TurbulenceConditionsTable = &schema.Table{
		Name:       "turbulence_conditions",
		Columns:    TurbulenceConditionsColumns,
		PrimaryKey: []*schema.Column{TurbulenceConditionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "turbulence_conditions_forecasts_turbulence_conditions",
				Columns:    []*schema.Column{TurbulenceConditionsColumns[4]},
				RefColumns: []*schema.Column{ForecastsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WeatherStationsColumns holds the columns for the "weather_stations" table.
	WeatherStationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "station_id", Type: field.TypeString, Unique: true},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "elevation", Type: field.TypeFloat64, Nullable: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "airport_station", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// WeatherStationsTable holds the schema information for the "weather_stations" table.
	WeatherStationsTable = &schema.Table{
		Name:       "weather_stations",
		Columns:    WeatherStationsColumns,
		PrimaryKey: []*schema.Column{WeatherStationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weather_stations_airports_station",
				Columns:    []*schema.Column{WeatherStationsColumns[6]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "weatherstation_station_id",
				Unique:  true,
				Columns: []*schema.Column{WeatherStationsColumns[1]},
			},
			{
				Name:    "weatherstation_latitude",
				Unique:  false,
				Columns: []*schema.Column{WeatherStationsColumns[2]},
			},
			{
				Name:    "weatherstation_longitude",
				Unique:  false,
				Columns: []*schema.Column{WeatherStationsColumns[3]},
			},
			{
				Name:    "weatherstation_hash",
				Unique:  false,
				Columns: []*schema.Column{WeatherStationsColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AirportsTable,
		CountriesTable,
		ForecastsTable,
		FrequenciesTable,
		IcingConditionsTable,
		MetarsTable,
		RegionsTable,
		RunwaysTable,
		SkyConditionsTable,
		TafsTable,
		TemperatureDataTable,
		TurbulenceConditionsTable,
		WeatherStationsTable,
	}
)

func init() {
	AirportsTable.ForeignKeys[0].RefTable = CountriesTable
	AirportsTable.ForeignKeys[1].RefTable = RegionsTable
	AirportsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_520_ci",
	}
	ForecastsTable.ForeignKeys[0].RefTable = TafsTable
	FrequenciesTable.ForeignKeys[0].RefTable = AirportsTable
	IcingConditionsTable.ForeignKeys[0].RefTable = ForecastsTable
	MetarsTable.ForeignKeys[0].RefTable = WeatherStationsTable
	RunwaysTable.ForeignKeys[0].RefTable = AirportsTable
	SkyConditionsTable.ForeignKeys[0].RefTable = ForecastsTable
	SkyConditionsTable.ForeignKeys[1].RefTable = MetarsTable
	TafsTable.ForeignKeys[0].RefTable = WeatherStationsTable
	TemperatureDataTable.ForeignKeys[0].RefTable = ForecastsTable
	TurbulenceConditionsTable.ForeignKeys[0].RefTable = ForecastsTable
	WeatherStationsTable.ForeignKeys[0].RefTable = AirportsTable
}

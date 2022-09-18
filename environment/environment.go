package environment

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"strconv"
	"strings"
)

type Environment struct {
	AdminSecret                 string `mapstructure:"ADMIN_SECRET"`
	Database                    string `mapstructure:"DATABASE"`
	Port                        string `mapstructure:"PORT"`
	MaxConcurrentImports        int    `mapstructure:"MAX_CONCURRENT_IMPORTS"`
	AxiomDataset                string `mapstructure:"AXIOM_DATASET"`
	KeepDataForDays             int    `mapstructure:"KEEP_DATA_FOR_DAYS"`
	GraphQLQueryComplexityLimit int    `mapstructure:"GRAPHQL_QUERY_COMPLEXITY_LIMIT"`
}

var Global Environment

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		// We don't want to panic here, because we might be running in a production environment
	}

	data := make(map[string]interface{})

	// Load all environment variables into a map
	for _, s := range os.Environ() {
		split := strings.SplitN(s, "=", 2)

		if split[0] == "MAX_CONCURRENT_IMPORTS" {
			// Convert to int
			data[split[0]], err = strconv.Atoi(split[1])
			if err != nil {
				log.Fatal("Could not convert MAX_CONCURRENT_IMPORTS to int")
			}

			continue
		}

		if split[0] == "KEEP_DATA_FOR_DAYS" {
			// Convert to int
			data[split[0]], err = strconv.Atoi(split[1])
			if err != nil || data[split[0]].(int) < 1 {
				log.Println("Did not receive a valid value for KEEP_DATA_FOR_DAYS, defaulting to 1 day")
				data[split[0]] = 1
			}

			continue
		}

		if split[0] == "GRAPHQL_QUERY_COMPLEXITY_LIMIT" {
			// Convert to int
			data[split[0]], err = strconv.Atoi(split[1])
			if err != nil || data[split[0]].(int) < 1 {
				log.Println("Did not receive a valid value for GRAPHQL_QUERY_COMPLEXITY_LIMIT, defaulting to 80")
				data[split[0]] = 80
			}

			continue
		}

		data[split[0]] = split[1]
	}

	// Decode the map into the Global variable
	err = mapstructure.Decode(data, &Global)
	if err != nil {
		log.Fatal(err)
	}

	// Set default values
	if Global.MaxConcurrentImports == 0 {
		Global.MaxConcurrentImports = 1
	}

	if Global.KeepDataForDays == 0 {
		Global.KeepDataForDays = 1
	}

	if Global.GraphQLQueryComplexityLimit == 0 {
		Global.GraphQLQueryComplexityLimit = 80
	}

	if Global.Port == "" {
		Global.Port = "80"
	}
}

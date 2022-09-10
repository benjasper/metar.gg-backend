package environment

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"metar.gg/logging"
	"os"
	"strconv"
	"strings"
)

type Environment struct {
	AdminSecret          string `mapstructure:"ADMIN_SECRET"`
	Database             string `mapstructure:"DATABASE"`
	Port                 string `mapstructure:"PORT"`
	MaxConcurrentImports int    `mapstructure:"MAX_CONCURRENT_IMPORTS"`
}

var Global Environment

func Initialize(logger *logging.Logger) {
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
				logger.Error("Could not convert MAX_CONCURRENT_IMPORTS to int")
				continue
			}

			continue
		}

		data[split[0]] = split[1]
	}

	// Decode the map into the Global variable
	err = mapstructure.Decode(data, &Global)
	if err != nil {
		logger.Fatal(err)
	}
}

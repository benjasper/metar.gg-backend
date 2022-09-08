package environment

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"metar.gg/logging"
	"os"
	"strings"
)

type Environment struct {
	AdminSecret string `mapstructure:"ADMIN_SECRET"`
	Database    string `mapstructure:"DATABASE"`
	Port        string `mapstructure:"PORT"`
}

var Global Environment

func Initialize(logger *logging.Logger) {
	err := godotenv.Load()
	if err != nil {
		// We don't want to panic here, because we might be running in a production environment
	}

	data := make(map[string]string)

	// Load all environment variables into a map
	for _, s := range os.Environ() {
		split := strings.SplitN(s, "=", 2)
		data[split[0]] = split[1]
	}

	// Decode the map into the Global variable
	err = mapstructure.Decode(data, &Global)
	if err != nil {
		logger.Fatal(err)
	}
}

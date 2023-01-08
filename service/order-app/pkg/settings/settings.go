package settings

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Settings struct {
	App  AppSettings        `mapstructure:"app"`
	Conn ConnectionSettings `mapstructure:"conn"`
}

// Load global application settings, specify altPaths to add alternative config search paths
func Load(altPaths ...string) (*Settings, error) {
	appName := filepath.Base(os.Args[0])

	// remove extension
	appNames := strings.Split(appName, ".")
	if appNames[len(appNames)-1] == "exe" {
		appNames = appNames[:len(appNames)-1]
	}

	appName = ""
	for i, a := range appNames {
		if a == "exe" {
			continue
		}

		appName += a
		if i != len(appNames)-1 {
			appName += "."
		}
	}

	//config paths
	viper.SetConfigType("yaml")
	viper.SetConfigName(appName)
	viper.AddConfigPath(".")

	for _, path := range altPaths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	s := new(Settings)
	if err := viper.Unmarshal(&s); err != nil {
		return nil, err
	}

	return s, nil
}

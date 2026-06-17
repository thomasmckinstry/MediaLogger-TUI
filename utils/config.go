package utils

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type theme struct {
	Focus   string `yaml:"Focus"`
	Unfocus string `yaml:"Unfocus"`
}

type ConfigModel struct {
	Theme theme `yaml:"Theme"`

	Database string `yaml:"Database"`

	MediaOptions  []string `yaml:"MediaOptions"`
	StatusOptions []string `yaml:"StatusOptions"`
}

// Starts with default values. Overwritten if a config file exists.
var Config ConfigModel = ConfigModel{
	Theme: theme{
		Focus:   "#D17600",
		Unfocus: "#6E3F00",
	},

	Database: "media.db",

	MediaOptions:  []string{"Movie", "Book", "TV Show"},
	StatusOptions: []string{"Pending", "Started", "Completed"},
}

func ReadConfig(filepath string) {
	DebugLog("Config Filepath: ", filepath)
	file, err := os.ReadFile(filepath)
	DebugLog("Read Config Error: ", err)
	if !errors.Is(err, os.ErrNotExist) && err != nil {
		CheckError("Failed to read config: ", err)
	}

	err = yaml.Unmarshal(file, &Config)
	CheckError("Failed to unmarshal config: ", err)
	SetTheme(Config.Theme.Focus, Config.Theme.Unfocus)
	DebugLog("Config: ", Config)
}

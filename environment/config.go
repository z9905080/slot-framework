package environment

import (
	"os"
	"slot-framework/pkg/encoder"
)

type ConfigPathType string

var ConfigPath ConfigPathType

// New init a config env
func New(path ConfigPathType) (Config, error) {
	var config Config
	// load from path
	data, openErr := os.ReadFile(string(path))
	if openErr != nil {
		return Config{}, openErr
	}

	unmarshalErr := encoder.Json.Unmarshal(data, &config)
	if unmarshalErr != nil {
		return Config{}, unmarshalErr
	}

	return config, nil
}

type Config struct {
	LogSetting struct {
		Level  string `json:"level"`
		Output string `json:"output"`
		Depth  int    `json:"depth"`
	} `json:"log_setting"`
}

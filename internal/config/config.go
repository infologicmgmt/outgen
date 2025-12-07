/*
  Filename: config.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Handles loading and parsing of configuration files.
*/

package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadConfig reads a configuration file and returns a map of values.
// It supports both YAML and JSON file formats.
func LoadConfig(configFile string) (map[string]interface{}, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var values map[string]interface{}
	ext := filepath.Ext(configFile)

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &values)
	case ".json":
		err = json.Unmarshal(data, &values)
	}

	return values, err
}

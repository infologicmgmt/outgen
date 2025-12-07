/*
  Filename: config_test.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Unit tests for the config package.
*/

package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test with a YAML file
	yamlContent := `
name: outgen
version: 1.0
`
	yamlFile, err := os.CreateTemp("", "*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(yamlFile.Name())
	if _, err := yamlFile.WriteString(yamlContent); err != nil {
		t.Fatal(err)
	}
	yamlFile.Close()

	yamlConfig, err := LoadConfig(yamlFile.Name())
	if err != nil {
		t.Errorf("Failed to load YAML config: %v", err)
	}
	if yamlConfig["name"] != "outgen" {
		t.Errorf("Expected name to be 'outgen', but got '%s'", yamlConfig["name"])
	}

	// Test with a JSON file
	jsonContent := `
{
  "name": "outgen",
  "version": 1.0
}
`
	jsonFile, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(jsonFile.Name())
	if _, err := jsonFile.WriteString(jsonContent); err != nil {
		t.Fatal(err)
	}
	jsonFile.Close()

	jsonConfig, err := LoadConfig(jsonFile.Name())
	if err != nil {
		t.Errorf("Failed to load JSON config: %v", err)
	}
	if jsonConfig["name"] != "outgen" {
		t.Errorf("Expected name to be 'outgen', but got '%s'", jsonConfig["name"])
	}
}

/*
  Filename: processor_test.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Unit tests for the processor package.
*/

package processor

import (
	"os"
	"testing"
)

func TestGetOutputFile(t *testing.T) {
	testCases := []struct {
		inputFile string
		expected  string
	}{
		{"test.txt.j2", "test.txt"},
		{"Dockerfile.mustache", "Dockerfile"},
		{"config.json.hbs", "config.json"},
	}

	for _, tc := range testCases {
		actual := GetOutputFile(tc.inputFile)
		if actual != tc.expected {
			t.Errorf("For inputFile '%s', expected output file '%s', but got '%s'", tc.inputFile, tc.expected, actual)
		}
	}
}

func TestJobProcess(t *testing.T) {
	// Create a temporary template file
	templateContent := "Hello, {{.name}}!"
	templateFile, err := os.CreateTemp("", "*.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(templateFile.Name())
	if _, err := templateFile.WriteString(templateContent); err != nil {
		t.Fatal(err)
	}
	templateFile.Close()

	// Create a job to process the template file
	outputFile := GetOutputFile(templateFile.Name())
	defer os.Remove(outputFile)
	job := &Job{
		InputFile:  templateFile.Name(),
		OutputFile: outputFile,
		Format:     "go",
		Overwrite:  true,
		Config:     map[string]interface{}{"name": "World"},
	}

	// Process the job
	if err := job.Process(); err != nil {
		t.Errorf("Job processing failed: %v", err)
	}

	// Check the output file
	outputContent, err := os.ReadFile(outputFile)
	if err != nil {
		t.Errorf("Failed to read output file: %v", err)
	}
	expectedContent := "Hello, World!"
	if string(outputContent) != expectedContent {
		t.Errorf("Expected output content '%s', but got '%s'", expectedContent, string(outputContent))
	}
}

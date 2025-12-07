/*
  Filename: processor.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Handles concurrent file processing.
*/

package processor

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/infologicmgmt/outgen/internal/log"
	"github.com/infologicmgmt/outgen/internal/template"
)

// Job represents a file to be processed.
type Job struct {
	InputFile  string
	OutputFile string
	Format     string
	Overwrite  bool
	Config     map[string]interface{}
}

// Result holds the outcome of a job.
type Result struct {
	Job   Job
	Error error
}

// Process handles the processing of a single file.
func (j *Job) Process() error {
	log.Logger.Info().Str("inputFile", j.InputFile).Msg("Processing file")

	// Read the input file
	var inputContent []byte
	var err error
	if j.InputFile == "stdin" {
		inputContent, err = io.ReadAll(os.Stdin)
	} else {
		inputContent, err = os.ReadFile(j.InputFile)
	}
	if err != nil {
		return err
	}

	// Determine the template format
	format := j.Format
	if format == "" {
		format = template.DetectFormatFromExtension(j.InputFile)
	}
	if format == "" {
		return fmt.Errorf("could not detect template format for file: %s", j.InputFile)
	}

	// Get the template engine
	engine, err := template.NewTemplateEngine(format)
	if err != nil {
		return err
	}

	// Render the template
	outputContent, err := engine.Render(string(inputContent), j.Config)
	if err != nil {
		return err
	}

	// Write the output file
	if j.OutputFile == "stdout" {
		_, err = os.Stdout.WriteString(outputContent)
		return err
	}

	if !j.Overwrite {
		if _, err := os.Stat(j.OutputFile); !os.IsNotExist(err) {
			return fmt.Errorf("output file already exists: %s. Use --overwrite to replace it", j.OutputFile)
		}
	}

	return os.WriteFile(j.OutputFile, []byte(outputContent), 0644)
}

// GetOutputFile returns the output filename for a given input file.
func GetOutputFile(inputFile string) string {
	ext := filepath.Ext(inputFile)
	return strings.TrimSuffix(inputFile, ext)
}

/*
  Filename: engine.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Defines the interface for template engines.
*/

package template

// TemplateEngine is the interface that all template engines must implement.
type TemplateEngine interface {
	// Render processes the template with the given data and returns the output.
	Render(templateContent string, data map[string]interface{}) (string, error)
}

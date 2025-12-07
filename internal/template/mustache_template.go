/*
  Filename: mustache_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the Mustache template engine.
*/

package template

import (
	"github.com/cbroglie/mustache"
)

// MustacheTemplateEngine is an adapter for the cbroglie/mustache implementation.
type MustacheTemplateEngine struct{}

// Render processes the template using the cbroglie/mustache engine.
func (e *MustacheTemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	return mustache.Render(templateContent, data)
}

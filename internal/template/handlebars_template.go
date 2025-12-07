/*
  Filename: handlebars_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the Handlebars template engine.
*/

package template

import (
	"github.com/aymerick/raymond"
)

// HandlebarsTemplateEngine is an adapter for the aymerick/raymond Handlebars implementation.
type HandlebarsTemplateEngine struct{}

// Render processes the template using the aymerick/raymond Handlebars engine.
func (e *HandlebarsTemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	tpl, err := raymond.Parse(templateContent)
	if err != nil {
		return "", err
	}
	return tpl.Exec(data)
}

/*
  Filename: jinja_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the Jinja2 template engine.
*/

package template

import (
	"github.com/flosch/pongo2/v6"
)

// JinjaTemplateEngine is an adapter for the pongo2 Jinja2 implementation.
type JinjaTemplateEngine struct{}

// Render processes the template using the pongo2 Jinja2 engine.
func (e *JinjaTemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	tpl, err := pongo2.FromString(templateContent)
	if err != nil {
		return "", err
	}
	return tpl.Execute(pongo2.Context(data))
}

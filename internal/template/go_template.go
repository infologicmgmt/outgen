/*
  Filename: go_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the Go text/template engine.
*/

package template

import (
	"bytes"
	"text/template"
)

// GoTemplateEngine is an adapter for Go's built-in text/template engine.
type GoTemplateEngine struct{}

// Render processes the template using Go's text/template engine.
func (e *GoTemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	tmpl, err := template.New("go-template").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

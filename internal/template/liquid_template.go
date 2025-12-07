/*
  Filename: liquid_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the Liquid template engine.
*/

package template

import (
	"github.com/osteele/liquid"
)

// LiquidTemplateEngine is an adapter for the osteele/liquid implementation.
type LiquidTemplateEngine struct{}

// Render processes the template using the osteele/liquid engine.
func (e *LiquidTemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	engine := liquid.NewEngine()
	return engine.ParseAndRenderString(templateContent, data)
}

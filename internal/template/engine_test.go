/*
  Filename: engine_test.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Unit tests for the template engines.
*/

package template

import (
	"testing"
)

func TestGoTemplateEngine(t *testing.T) {
	engine := &GoTemplateEngine{}
	template := "Hello, {{.name}}!"
	data := map[string]interface{}{"name": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("Go template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestJinjaTemplateEngine(t *testing.T) {
	engine := &JinjaTemplateEngine{}
	template := "Hello, {{name}}!"
	data := map[string]interface{}{"name": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("Jinja template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestMustacheTemplateEngine(t *testing.T) {
	engine := &MustacheTemplateEngine{}
	template := "Hello, {{name}}!"
	data := map[string]interface{}{"name": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("Mustache template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestHandlebarsTemplateEngine(t *testing.T) {
	engine := &HandlebarsTemplateEngine{}
	template := "Hello, {{name}}!"
	data := map[string]interface{}{"name": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("Handlebars template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestLiquidTemplateEngine(t *testing.T) {
	engine := &LiquidTemplateEngine{}
	template := "Hello, {{name}}!"
	data := map[string]interface{}{"name": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("Liquid template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestM4TemplateEngine(t *testing.T) {
	engine := &M4TemplateEngine{}
	template := "Hello, NAME!"
	data := map[string]interface{}{"NAME": "World"}
	expected := "Hello, World!"
	actual, err := engine.Render(template, data)
	if err != nil {
		t.Errorf("M4 template render failed: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestDetectFormatFromExtension(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{"test.j2", "jinja"},
		{"test.mustache", "mustache"},
		{"test.hbs", "handlebars"},
		{"test.liquid", "liquid"},
		{"test.tmpl", "go"},
		{"test.m4", "m4"},
		{"test.txt", ""},
	}

	for _, tc := range testCases {
		actual := DetectFormatFromExtension(tc.filename)
		if actual != tc.expected {
			t.Errorf("For filename '%s', expected format '%s', but got '%s'", tc.filename, tc.expected, actual)
		}
	}
}

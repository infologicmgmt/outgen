/*
  Filename: factory.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Factory for creating template engine instances.
*/

package template

import (
	"fmt"
	"path/filepath"
	"strings"
)

// NewTemplateEngine returns a TemplateEngine instance based on the format.
func NewTemplateEngine(format string) (TemplateEngine, error) {
	switch format {
	case "go", "gotemplate":
		return &GoTemplateEngine{}, nil
	case "jinja":
		return &JinjaTemplateEngine{}, nil
	case "mustache":
		return &MustacheTemplateEngine{}, nil
	case "handlebars":
		return &HandlebarsTemplateEngine{}, nil
	case "liquid":
		return &LiquidTemplateEngine{}, nil
	case "m4":
		return &M4TemplateEngine{}, nil
	default:
		return nil, fmt.Errorf("unsupported template format: %s", format)
	}
}

// DetectFormatFromExtension determines the template format from a file extension.
func DetectFormatFromExtension(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".j2", ".jinja2":
		return "jinja"
	case ".mustache":
		return "mustache"
	case ".hbs":
		return "handlebars"
	case ".liquid":
		return "liquid"
	case ".tmpl", ".tpl":
		return "go"
	case ".m4":
		return "m4"
	default:
		return ""
	}
}

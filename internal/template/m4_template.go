/*
  Filename: m4_template.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements a wrapper for the m4 preprocessor.
*/

package template

import (
	"bytes"
	"fmt"
	"os/exec"
)

// M4TemplateEngine is a wrapper for the system's m4 preprocessor.
type M4TemplateEngine struct{}

// Render processes the template by invoking the m4 preprocessor.
func (e *M4TemplateEngine) Render(templateContent string, data map[string]interface{}) (string, error) {
	args := []string{}
	for key, value := range data {
		args = append(args, "-D", fmt.Sprintf("%s=%v", key, value))
	}

	cmd := exec.Command("m4", args...)
	cmd.Stdin = bytes.NewBufferString(templateContent)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("m4 execution failed: %v\nstderr: %s", err, stderr.String())
	}

	return out.String(), nil
}

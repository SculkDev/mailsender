package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var templateVarRegex = regexp.MustCompile(`\{\{\s*(\w+)\s*\}\}`)

func RenderTemplate(templateType TemplateType, vars map[string]string) (html, text, subject string, err error) {
	baseDir := filepath.Join("templates", string(templateType))

	html, err = renderFile(filepath.Join(baseDir, "html.template"), vars)
	if err != nil {
		return "", "", "", fmt.Errorf("render html template: %w", err)
	}

	text, err = renderFile(filepath.Join(baseDir, "text.template"), vars)
	if err != nil {
		return "", "", "", fmt.Errorf("render text template: %w", err)
	}

	subject, err = renderFile(filepath.Join(baseDir, "subject.template"), vars)
	if err != nil {
		return "", "", "", fmt.Errorf("render subject template: %w", err)
	}

	return html, text, subject, nil
}

func renderFile(path string, vars map[string]string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}

	result := templateVarRegex.ReplaceAllStringFunc(string(data), func(match string) string {
		submatch := templateVarRegex.FindStringSubmatch(match)
		if len(submatch) < 2 {
			return match
		}
		key := submatch[1]
		if val, ok := vars[key]; ok {
			return val
		}
		return match
	})

	return result, nil
}

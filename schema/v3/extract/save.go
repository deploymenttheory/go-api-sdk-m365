package extract

import (
	"fmt"
	"os"
	"text/template"
)

// pathsTemplate is the template for the generated Go file with paths
const pathsTemplate = `package msgraphpaths

var Paths = []string{
{{- range . }}
    "{{ . }}",
{{- end }}
}
`

// SavePathsToFile saves the paths to a file using the template
func SavePathsToFile(paths []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("paths").Parse(pathsTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	err = tmpl.Execute(file, paths)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

package extract

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// SaveURLPathsToFile saves the paths to a file using a template that groups paths by their first segment
func SaveURLPathsToFile(paths []string, path string) error {
	// Group paths by their first segment, treating parameterized segments uniformly
	groupedPaths := groupPathsByFirstSegment(paths)

	// Define the template for the grouped paths
	const pathsTemplate = `package msgraphpaths

{{- range $group, $paths := . }}

var {{ $group }} = []string{
{{- range $paths }}
    "{{ . }}",
{{- end }}
}
{{- end }}
`

	// Create and open the file for writing
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Parse the template
	tmpl, err := template.New("paths").Parse(pathsTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the grouped paths
	err = tmpl.Execute(file, groupedPaths)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

// groupPathsByFirstSegment groups the paths by their first segment, normalizing parameterized segments
func groupPathsByFirstSegment(paths []string) map[string][]string {
	groupedPaths := make(map[string][]string)
	segmentRegex := regexp.MustCompile(`^\w+`)

	for _, path := range paths {
		segments := strings.Split(strings.TrimPrefix(path, "/"), "/")
		if len(segments) > 0 {
			// Use regex to capture the first segment before any parameterized part
			firstSegment := segmentRegex.FindString(segments[0])
			if firstSegment != "" {
				groupedPaths[firstSegment] = append(groupedPaths[firstSegment], path)
			}
		}
	}
	return groupedPaths
}

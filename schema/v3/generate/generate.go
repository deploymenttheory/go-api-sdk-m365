package generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const structTemplate = `
{{- range $exampleName, $example := .Examples }}
type {{ formatName $exampleName }} struct {
{{- range $fieldName, $fieldValue := $example.value }}
    {{ formatFieldName $fieldName }} {{ formatType $fieldValue }}
{{- end }}
}
{{- end }}
`

func formatName(name string) string {
	caser := cases.Title(language.English)
	return strings.ReplaceAll(caser.String(strings.ReplaceAll(name, ".", " ")), " ", "")
}

func formatFieldName(name string) string {
	caser := cases.Title(language.English)
	return caser.String(name)
}

func formatType(value interface{}) string {
	switch v := value.(type) {
	case string:
		if strings.HasPrefix(v, "microsoft.graph.") {
			return "*" + formatName(v)
		}
		return "string"
	case map[string]interface{}:
		return "*" + formatName(v["@odata.type"].(string))
	case []interface{}:
		elemType := formatType(v[0])
		return "[]" + elemType
	default:
		return "interface{}"
	}
}

func GenerateStructs(examples map[string]interface{}) (string, error) {
	funcMap := template.FuncMap{
		"formatName":      formatName,
		"formatFieldName": formatFieldName,
		"formatType":      formatType,
	}

	tmpl, err := template.New("structs").Funcs(funcMap).Parse(structTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]interface{}{"Examples": examples})
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

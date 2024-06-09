package template

const structTemplate = `
{{- range $exampleName, $example := .Examples }}
type {{ formatName $exampleName }} struct {
{{- range $fieldName, $fieldValue := $example.value }}
    {{ formatFieldName $fieldName }} {{ formatType $fieldValue }}
{{- end }}
}
{{- end }}
`

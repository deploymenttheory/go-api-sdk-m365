package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Edmx represents the root element of the CSDL file
type Edmx struct {
	XMLName     xml.Name     `xml:"Edmx"`
	Schemas     []Schema     `xml:"DataServices>Schema"`
	Annotations []Annotation `xml:"Annotations>Annotation"`
}

// Schema represents a schema in the CSDL file
type Schema struct {
	XMLName      xml.Name      `xml:"Schema"`
	EntityTypes  []EntityType  `xml:"EntityType"`
	ComplexTypes []ComplexType `xml:"ComplexType"`
	EnumTypes    []EnumType    `xml:"EnumType"`
}

// EntityType represents an entity type in the CSDL schema
type EntityType struct {
	Name        string       `xml:"Name,attr"`
	Properties  []Property   `xml:"Property"`
	Annotations []Annotation `xml:"Annotation"`
}

// ComplexType represents a complex type in the CSDL schema
type ComplexType struct {
	Name        string       `xml:"Name,attr"`
	Properties  []Property   `xml:"Property"`
	Annotations []Annotation `xml:"Annotation"`
}

// Property represents a property of an entity or complex type
type Property struct {
	Name        string       `xml:"Name,attr"`
	Type        string       `xml:"Type,attr"`
	Annotations []Annotation `xml:"Annotation"`
}

// EnumType represents an enumeration type in the CSDL schema
type EnumType struct {
	Name        string       `xml:"Name,attr"`
	Members     []EnumMember `xml:"Member"`
	Annotations []Annotation `xml:"Annotation"`
}

// EnumMember represents a member of an enumeration type
type EnumMember struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

// Annotation represents an annotation in the CSDL schema
type Annotation struct {
	Target      string                     `xml:"Target,attr"`
	Term        string                     `xml:"Term,attr"`
	StringValue string                     `xml:"String,attr"`
	Collection  []AnnotationCollectionItem `xml:"Collection>Record"`
}

// AnnotationCollectionItem represents an item in a collection annotation
type AnnotationCollectionItem struct {
	PropertyValues []PropertyValue `xml:"PropertyValue"`
}

// PropertyValue represents a property value in an annotation
type PropertyValue struct {
	Property   string `xml:"Property,attr"`
	Value      string `xml:"String,attr"`
	Date       string `xml:"Date,attr"`
	EnumMember string `xml:"EnumMember"`
}

const StructTemplate = `{{- if .Annotations }}
// {{.Name}} 
{{- range .Annotations}}
// {{.Term}}: {{.StringValue}}
{{- range .Collection}}
// {{range .PropertyValues}}
// {{.Property}}: {{if .Value}}{{.Value}}{{else if .Date}}{{.Date}}{{else if .EnumMember}}{{.EnumMember}}{{end}}
{{- end}}{{end}}{{end}}{{end}}
type {{.Name}} struct {
{{- range .Properties}}
    {{- if .Annotations}}{{range .Annotations}}
// {{.Term}}: {{.StringValue}}
    {{- range .Collection}}
// {{range .PropertyValues}}
// {{.Property}}: {{if .Value}}{{.Value}}{{else if .Date}}{{.Date}}{{else if .EnumMember}}{{.EnumMember}}{{end}}
{{- end}}{{end}}{{end}}{{end}}
    {{.Name}} {{.Type}} ` + "`json:\"{{.JSONName}},omitempty\"`" + `
{{- end}}
}
`

const EnumTemplate = `{{- if .Annotations }}
// {{.Name}} 
{{- range .Annotations}}
// {{.Term}}: {{.StringValue}}
{{- range .Collection}}
// {{range .PropertyValues}}
// {{.Property}}: {{if .Value}}{{.Value}}{{else if .Date}}{{.Date}}{{else if .EnumMember}}{{.EnumMember}}{{end}}
{{- end}}{{end}}{{end}}{{end}}
type {{.Name}} int

const (
{{- range .Members}}
    {{$.Name}}{{.Name}} {{$.Name}} = {{.Value}} // {{range .Annotations}}{{.StringValue}} {{end}}
{{- end}}
)
`

func main() {
	// Define flags for input and output file paths
	inputPath := flag.String("input", "schema.csdl", "Path to the CSDL file")
	outputPath := flag.String("output", "output/output.go", "Path to save the generated Go structs")
	flag.Parse()

	log.Println("Starting the CSDL processing...")

	// Read the CSDL file
	data, err := os.ReadFile(*inputPath)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	log.Println("Successfully read the input file.")

	var edmx Edmx

	// Unmarshal the XML into the Edmx struct
	err = xml.Unmarshal(data, &edmx)
	if err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}
	log.Println("Successfully unmarshalled the XML data.")

	// Ensure the output directory exists
	outputDir := filepath.Dir(*outputPath)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
	log.Println("Output directory ensured.")

	// Open the output file
	outputFile, err := os.Create(*outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Write the package name at the top of the output file
	fmt.Fprintln(outputFile, "package graphmodels")
	fmt.Fprintln(outputFile)
	fmt.Fprintln(outputFile, "import \"time\"")
	fmt.Fprintln(outputFile)

	// Map to track generated structs to avoid duplication
	generatedStructs := make(map[string]bool)

	// Generate Go structs from the CSDL schemas
	log.Println("Generating Go structs...")
	for _, schema := range edmx.Schemas {
		for _, complexType := range schema.ComplexTypes {
			err := GenerateStruct(outputFile, complexType.Name, complexType.Properties, complexType.Annotations, edmx.Annotations, generatedStructs)
			if err != nil {
				log.Fatalf("Error generating struct: %v", err)
			}
			log.Printf("Generated struct for %s\n", complexType.Name)
		}
	}

	// Generate Go enums from the CSDL schemas
	log.Println("Generating Go enums...")
	for _, schema := range edmx.Schemas {
		for _, enumType := range schema.EnumTypes {
			err := GenerateEnum(outputFile, enumType, edmx.Annotations)
			if err != nil {
				log.Fatalf("Error generating enum: %v", err)
			}
			log.Printf("Generated enum for %s\n", enumType.Name)
		}
	}

	log.Printf("Go structs and enums generated and saved to %s\n", *outputPath)
}

// GenerateStruct generates a Go struct from a complex type
func GenerateStruct(outputFile *os.File, structName string, properties []Property, localAnnotations []Annotation, globalAnnotations []Annotation, generatedStructs map[string]bool) error {
	if generatedStructs[structName] {
		return nil
	}
	generatedStructs[structName] = true

	// Collect all relevant annotations
	annotations := append(localAnnotations, findGlobalAnnotations(structName, globalAnnotations)...)

	tmpl, err := template.New("struct").Parse(StructTemplate)
	if err != nil {
		return fmt.Errorf("error parsing struct template: %w", err)
	}

	var buf bytes.Buffer
	data := struct {
		Name       string
		Properties []struct {
			Name        string
			Type        string
			JSONName    string
			Annotations []Annotation
		}
		Annotations []Annotation
	}{
		Name:        capitalize(structName),
		Annotations: annotations,
	}

	for _, prop := range properties {
		goType, _ := mapType(prop.Type)
		data.Properties = append(data.Properties, struct {
			Name        string
			Type        string
			JSONName    string
			Annotations []Annotation
		}{
			Name:        capitalize(prop.Name),
			Type:        goType,
			JSONName:    prop.Name,
			Annotations: prop.Annotations,
		})
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("error executing struct template: %w", err)
	}

	_, err = outputFile.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}

	return nil
}

// GenerateEnum generates a Go enum from an enum type
func GenerateEnum(outputFile *os.File, enumType EnumType, globalAnnotations []Annotation) error {
	tmpl, err := template.New("enum").Parse(EnumTemplate)
	if err != nil {
		return fmt.Errorf("error parsing enum template: %w", err)
	}

	var buf bytes.Buffer
	data := struct {
		Name    string
		Members []struct {
			Name        string
			Value       string
			Annotations []Annotation
		}
		Annotations []Annotation
	}{
		Name:        capitalize(enumType.Name),
		Annotations: findGlobalAnnotations(enumType.Name, globalAnnotations),
	}

	for _, member := range enumType.Members {
		// Attach annotations specific to this member
		memberAnnotations := findGlobalAnnotations(fmt.Sprintf("%s/%s", enumType.Name, member.Name), globalAnnotations)
		log.Printf("Enum member: %s/%s has %d annotations\n", enumType.Name, member.Name, len(memberAnnotations))
		for _, ann := range memberAnnotations {
			log.Printf("Annotation for %s: %s - %s\n", member.Name, ann.Term, ann.StringValue)
		}
		data.Members = append(data.Members, struct {
			Name        string
			Value       string
			Annotations []Annotation
		}{
			Name:        capitalize(member.Name),
			Value:       member.Value,
			Annotations: memberAnnotations,
		})
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("error executing enum template: %w", err)
	}

	_, err = outputFile.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}

	return nil
}

// mapType maps a CSDL type to a Go type and indicates if it's a complex type
func mapType(csdlType string) (string, bool) {
	switch csdlType {
	case "Edm.String":
		return "string", false
	case "Edm.Int32":
		return "int", false
	case "Edm.Int64":
		return "int64", false
	case "Edm.Int16":
		return "int16", false
	case "Edm.Boolean":
		return "*bool", false
	case "Edm.DateTimeOffset":
		return "time.Time", false
	case "Edm.Duration":
		return "time.Duration", false
	case "Edm.Binary":
		return "[]byte", false
	case "Edm.Guid":
		return "string", false
	case "Edm.TimeOfDay":
		return "string", false // no direct equivalent, usually handled as string or a custom type
	case "Edm.Date":
		return "string", false // no direct equivalent, usually handled as string or a custom type
	case "Edm.Double":
		return "float64", false
	case "Edm.Single":
		return "float32", false
	case "Edm.Decimal":
		return "string", false // Decimal precision can vary, often handled as string or a custom type
	default:
		if strings.HasPrefix(csdlType, "Collection(") && strings.HasSuffix(csdlType, ")") {
			elementType := csdlType[len("Collection(") : len(csdlType)-1]
			goType, isComplex := mapType(elementType)
			return "[]*" + goType, isComplex
		}

		// Split the type by '.' to extract the last part (type name)
		parts := strings.Split(csdlType, ".")
		typeName := parts[len(parts)-1]

		// Capitalize the type name directly
		return capitalize(typeName), true
	}
}

// capitalize capitalizes the first letter of a string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// findGlobalAnnotations finds global annotations for a given target
func findGlobalAnnotations(target string, annotations []Annotation) []Annotation {
	var result []Annotation
	for _, annotation := range annotations {
		if strings.HasSuffix(annotation.Target, target) {
			log.Printf("Found annotation for target %s: %s - %s\n", target, annotation.Term, annotation.StringValue)
			result = append(result, annotation)
		}
	}
	return result
}

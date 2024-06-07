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

// Edmx represents the root element of the CSDL schema
type Edmx struct {
	XMLName xml.Name `xml:"Edmx"`
	Schemas []Schema `xml:"DataServices>Schema"`
}

// Schema represents a schema in the CSDL schema
type Schema struct {
	XMLName      xml.Name      `xml:"Schema"`
	Namespace    string        `xml:"Namespace,attr"`
	Alias        string        `xml:"Alias,attr"`
	EntityTypes  []EntityType  `xml:"EntityType"`
	ComplexTypes []ComplexType `xml:"ComplexType"`
	EnumTypes    []EnumType    `xml:"EnumType"`
	Annotations  []Annotation  `xml:"Annotations>Annotation"`
}

// EntityType represents an entity type in the CSDL schema
type EntityType struct {
	Name                 string               `xml:"Name,attr"`
	BaseType             string               `xml:"BaseType,attr,omitempty"`
	Abstract             bool                 `xml:"Abstract,attr,omitempty"`
	OpenType             bool                 `xml:"OpenType,attr,omitempty"`
	HasStream            bool                 `xml:"HasStream,attr,omitempty"`
	Key                  *EntityKey           `xml:"Key,omitempty"`
	Properties           []Property           `xml:"Property"`
	NavigationProperties []NavigationProperty `xml:"NavigationProperty"`
	Annotations          []Annotation         `xml:"Annotation"`
}

// EntityKey represents the key of an entity type in the CSDL schema
type EntityKey struct {
	PropertyRefs []PropertyRef `xml:"PropertyRef"`
}

// PropertyRef represents a property reference in an entity key
type PropertyRef struct {
	Name string `xml:"Name,attr"`
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
	Nullable    bool         `xml:"Nullable,attr,omitempty"`
	Annotations []Annotation `xml:"Annotation"`
}

// NavigationProperty represents a navigation property in the CSDL schema
type NavigationProperty struct {
	Name           string       `xml:"Name,attr"`
	Type           string       `xml:"Type,attr"`
	Nullable       bool         `xml:"Nullable,attr,omitempty"`
	Partner        string       `xml:"Partner,attr,omitempty"`
	ContainsTarget bool         `xml:"ContainsTarget,attr,omitempty"`
	Annotations    []Annotation `xml:"Annotation"`
}

// EnumType represents an enumeration type in the CSDL schema
type EnumType struct {
	Name        string       `xml:"Name,attr"`
	Members     []EnumMember `xml:"Member"`
	Annotations []Annotation `xml:"Annotation"`
}

// EnumMember represents a member of an enumeration type
type EnumMember struct {
	Name        string       `xml:"Name,attr"`
	Value       string       `xml:"Value,attr"`
	Annotations []Annotation `xml:"Annotation"`
}

// Annotation represents an annotation in the CSDL schema
type Annotation struct {
	Target      string                     `xml:"Target,attr"`
	Term        string                     `xml:"Term,attr"`
	StringValue string                     `xml:"String,attr,omitempty"`
	BoolValue   bool                       `xml:"Bool,attr,omitempty"`
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
{{- end}}
{{- end}}
type {{.Name}} struct {
{{- range .Properties}}
    // {{.Name}}: {{range .Annotations}}{{.Term}}: {{.StringValue}} {{end}}
    {{.Name}} {{.Type}} ` + "`json:\"{{.JSONName}},omitempty\"`" + `
{{- end}}
{{- range .NavigationProperties}}
    // {{.Name}}: {{range .Annotations}}{{.Term}}: {{.StringValue}} {{end}}
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
type {{.Name}} string

const (
{{- range .Members}}
    {{.Name}} {{$.Name}} = "{{.Name}}" // {{range .Annotations}} {{.Term}}: {{.StringValue}} {{end}}
{{- end}}
)
`

func main() {
	inputPath, outputPath := parseFlags()
	log.Println("Starting the CSDL processing...")

	data := readFile(*inputPath)
	edmx := unmarshalXML(data)

	ensureOutputDir(*outputPath)
	outputFile := createOutputFile(*outputPath)
	defer outputFile.Close()

	writePackageHeader(outputFile)
	generatedStructs := make(map[string]bool)

	generateGoStructs(edmx, outputFile, generatedStructs)
	generateGoEnums(edmx, outputFile)

	log.Printf("Go structs and enums generated and saved to %s\n", *outputPath)
}

func parseFlags() (*string, *string) {
	inputPath := flag.String("input", "schema.csdl", "Path to the CSDL file")
	outputPath := flag.String("output", "output/output.go", "Path to save the generated Go structs")
	flag.Parse()
	return inputPath, outputPath
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	log.Println("Successfully read the input file.")
	return data
}

func unmarshalXML(data []byte) Edmx {
	var edmx Edmx
	err := xml.Unmarshal(data, &edmx)
	if err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	// Adding detailed logs for the unmarshalling process
	log.Println("Successfully unmarshalled the XML data.")
	log.Printf("Edmx Version: %s", edmx.XMLName.Space)
	log.Printf("Number of Schemas: %d", len(edmx.Schemas))
	for i, schema := range edmx.Schemas {
		log.Printf("Schema %d: Namespace: %s, Alias: %s", i, schema.Namespace, schema.Alias)
		log.Printf("Number of EntityTypes: %d", len(schema.EntityTypes))
		for j, entityType := range schema.EntityTypes {
			log.Printf("EntityType %d: Name: %s, BaseType: %s", j, entityType.Name, entityType.BaseType)
			log.Printf("Number of Properties: %d", len(entityType.Properties))
			for k, property := range entityType.Properties {
				log.Printf("Property %d: Name: %s, Type: %s, Nullable: %t", k, property.Name, property.Type, property.Nullable)
				for l, annotation := range property.Annotations {
					log.Printf("Property %d, Annotation %d: Term: %s, StringValue: %s, Target: %s", k, l, annotation.Term, annotation.StringValue, annotation.Target)
				}
			}
			log.Printf("Number of NavigationProperties: %d", len(entityType.NavigationProperties))
			for k, navProp := range entityType.NavigationProperties {
				log.Printf("NavigationProperty %d: Name: %s, Type: %s", k, navProp.Name, navProp.Type)
				for l, annotation := range navProp.Annotations {
					log.Printf("NavigationProperty %d, Annotation %d: Term: %s, StringValue: %s, Target: %s", k, l, annotation.Term, annotation.StringValue, annotation.Target)
				}
			}
			log.Printf("Number of Annotations on EntityType: %d", len(entityType.Annotations))
			for k, annotation := range entityType.Annotations {
				log.Printf("EntityType Annotation %d: Term: %s, StringValue: %s, Target: %s", k, annotation.Term, annotation.StringValue, annotation.Target)
			}
		}
		log.Printf("Number of ComplexTypes: %d", len(schema.ComplexTypes))
		for j, complexType := range schema.ComplexTypes {
			log.Printf("ComplexType %d: Name: %s", j, complexType.Name)
			log.Printf("Number of Properties: %d", len(complexType.Properties))
			for k, property := range complexType.Properties {
				log.Printf("Property %d: Name: %s, Type: %s, Nullable: %t", k, property.Name, property.Type, property.Nullable)
				for l, annotation := range property.Annotations {
					log.Printf("Property %d, Annotation %d: Term: %s, StringValue: %s, Target: %s", k, l, annotation.Term, annotation.StringValue, annotation.Target)
				}
			}
			log.Printf("Number of Annotations on ComplexType: %d", len(complexType.Annotations))
			for k, annotation := range complexType.Annotations {
				log.Printf("ComplexType Annotation %d: Term: %s, StringValue: %s, Target: %s", k, annotation.Term, annotation.StringValue, annotation.Target)
			}
		}
		log.Printf("Number of EnumTypes: %d", len(schema.EnumTypes))
		for j, enumType := range schema.EnumTypes {
			log.Printf("EnumType %d: Name: %s", j, enumType.Name)
			log.Printf("Number of Members: %d", len(enumType.Members))
			for k, member := range enumType.Members {
				log.Printf("Member %d: Name: %s, Value: %s", k, member.Name, member.Value)
				for l, annotation := range member.Annotations {
					log.Printf("Member %d, Annotation %d: Term: %s, StringValue: %s, Target: %s", k, l, annotation.Term, annotation.StringValue, annotation.Target)
				}
			}
			log.Printf("Number of Annotations on EnumType: %d", len(enumType.Annotations))
			for k, annotation := range enumType.Annotations {
				log.Printf("EnumType Annotation %d: Term: %s, StringValue: %s, Target: %s", k, annotation.Term, annotation.StringValue, annotation.Target)
			}
		}
		log.Printf("Number of Annotations on Schema: %d", len(schema.Annotations))
		for k, annotation := range schema.Annotations {
			log.Printf("Schema Annotation %d: Target: %s, Term: %s, StringValue: %s", k, annotation.Target, annotation.Term, annotation.StringValue)
		}
	}

	return edmx
}

func ensureOutputDir(path string) {
	outputDir := filepath.Dir(path)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
	log.Println("Output directory ensured.")
}

func createOutputFile(path string) *os.File {
	outputFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	return outputFile
}

func writePackageHeader(outputFile *os.File) {
	fmt.Fprintln(outputFile, "package graphmodels")
	fmt.Fprintln(outputFile)
	fmt.Fprintln(outputFile, "import \"time\"")
	fmt.Fprintln(outputFile)
}

func generateGoStructs(edmx Edmx, outputFile *os.File, generatedStructs map[string]bool) {
	log.Println("Generating Go structs...")
	for _, schema := range edmx.Schemas {
		for _, entityType := range schema.EntityTypes {
			generateStruct(outputFile, entityType.Name, entityType.Properties, entityType.NavigationProperties, entityType.Annotations, schema.Annotations, generatedStructs)
		}
		for _, complexType := range schema.ComplexTypes {
			generateStruct(outputFile, complexType.Name, complexType.Properties, nil, complexType.Annotations, schema.Annotations, generatedStructs)
		}
	}
}

func generateStruct(outputFile *os.File, structName string, properties []Property, navigationProperties []NavigationProperty, localAnnotations []Annotation, globalAnnotations []Annotation, generatedStructs map[string]bool) {
	if generatedStructs[structName] {
		return
	}
	generatedStructs[structName] = true

	entityTarget := fmt.Sprintf("microsoft.graph.%s", structName)
	annotations := collectAnnotations(entityTarget, localAnnotations, globalAnnotations)
	log.Printf("Collected %d annotations for struct %s", len(annotations), structName)

	data := prepareStructData(structName, properties, navigationProperties, annotations, globalAnnotations)
	tmpl := parseTemplate(StructTemplate)

	executeTemplate(tmpl, data, outputFile)
	log.Printf("Successfully generated struct for %s", structName)
}

func collectAnnotations(target string, localAnnotations, globalAnnotations []Annotation) []Annotation {
	var result []Annotation
	for _, annotation := range localAnnotations {
		if annotation.Target == target {
			result = append(result, annotation)
		}
	}
	for _, annotation := range globalAnnotations {
		if annotation.Target == target {
			result = append(result, annotation)
		}
	}
	return result
}

func prepareStructData(structName string, properties []Property, navigationProperties []NavigationProperty, annotations, globalAnnotations []Annotation) struct {
	Name       string
	Properties []struct {
		Name, Type, JSONName string
		Annotations          []Annotation
	}
	NavigationProperties []struct {
		Name, Type, JSONName string
		Annotations          []Annotation
	}
	Annotations []Annotation
} {
	data := struct {
		Name       string
		Properties []struct {
			Name, Type, JSONName string
			Annotations          []Annotation
		}
		NavigationProperties []struct {
			Name, Type, JSONName string
			Annotations          []Annotation
		}
		Annotations []Annotation
	}{
		Name:        capitalize(structName),
		Annotations: annotations,
	}

	for _, prop := range properties {
		goType, _ := mapType(prop.Type)
		propTarget := fmt.Sprintf("microsoft.graph.%s/%s", structName, prop.Name)
		propAnnotations := collectAnnotations(propTarget, prop.Annotations, globalAnnotations)
		log.Printf("Property: %s, Type: %s, Annotations: %v", prop.Name, goType, propAnnotations)
		data.Properties = append(data.Properties, struct {
			Name        string
			Type        string
			JSONName    string
			Annotations []Annotation
		}{
			Name:        capitalize(prop.Name),
			Type:        goType,
			JSONName:    prop.Name,
			Annotations: propAnnotations,
		})
	}

	for _, navProp := range navigationProperties {
		goType, _ := mapType(navProp.Type)
		navPropTarget := fmt.Sprintf("microsoft.graph.%s/%s", structName, navProp.Name)
		navPropAnnotations := collectAnnotations(navPropTarget, navProp.Annotations, globalAnnotations)
		log.Printf("NavigationProperty: %s, Type: %s, Annotations: %v", navProp.Name, goType, navPropAnnotations)
		data.NavigationProperties = append(data.NavigationProperties, struct {
			Name        string
			Type        string
			JSONName    string
			Annotations []Annotation
		}{
			Name:        capitalize(navProp.Name),
			Type:        goType,
			JSONName:    navProp.Name,
			Annotations: navPropAnnotations,
		})
	}

	return data
}

func parseTemplate(templateString string) *template.Template {
	tmpl, err := template.New("template").Parse(templateString)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	return tmpl
}

func executeTemplate(tmpl *template.Template, data interface{}, outputFile *os.File) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	_, err = outputFile.Write(buf.Bytes())
	if err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}
}

func generateGoEnums(edmx Edmx, outputFile *os.File) {
	log.Println("Generating Go enums...")
	for _, schema := range edmx.Schemas {
		for _, enumType := range schema.EnumTypes {
			generateEnum(outputFile, enumType, schema.Annotations)
		}
	}
}

func generateEnum(outputFile *os.File, enumType EnumType, globalAnnotations []Annotation) {
	tmpl := parseTemplate(EnumTemplate)

	annotations := findGlobalAnnotations(fmt.Sprintf("microsoft.graph.%s", enumType.Name), globalAnnotations)
	log.Printf("Generating enum for %s with %d global annotations", enumType.Name, len(annotations))

	data := prepareEnumData(enumType, annotations, globalAnnotations)
	executeTemplate(tmpl, data, outputFile)

	log.Printf("Successfully generated enum for %s", enumType.Name)
}

func prepareEnumData(enumType EnumType, annotations, globalAnnotations []Annotation) struct {
	Name    string
	Members []struct {
		Name, Value string
		Annotations []Annotation
	}
	Annotations []Annotation
} {
	data := struct {
		Name    string
		Members []struct {
			Name, Value string
			Annotations []Annotation
		}
		Annotations []Annotation
	}{
		Name:        capitalize(enumType.Name),
		Annotations: annotations,
	}

	for _, member := range enumType.Members {
		memberAnnotations := append(member.Annotations, findGlobalAnnotations(fmt.Sprintf("microsoft.graph.%s/%s", enumType.Name, member.Name), globalAnnotations)...)
		log.Printf("EnumMember: %s, Value: %s, Annotations: %v", member.Name, member.Value, memberAnnotations)

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

	return data
}

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
		parts := strings.Split(csdlType, ".")
		typeName := parts[len(parts)-1]
		return capitalize(typeName), true
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func findGlobalAnnotations(target string, annotations []Annotation) []Annotation {
	var result []Annotation
	log.Printf("Finding global annotations for target: %s", target)
	for _, annotation := range annotations {
		if annotation.Target == target || annotation.Target == fmt.Sprintf("microsoft.graph.%s", target) {
			log.Printf("Found annotation: %s - %s: %s", annotation.Target, annotation.Term, annotation.StringValue)
			result = append(result, annotation)
		}
	}
	log.Printf("Total annotations found for target %s: %d", target, len(result))
	return result
}

package openapi3

type OpenAPISpec struct {
	OpenAPI      string                   `yaml:"openapi" mapstructure:"openapi"`
	Info         map[string]interface{}   `yaml:"info" mapstructure:"info"`
	Servers      []map[string]interface{} `yaml:"servers,omitempty" mapstructure:"servers"`
	Paths        map[string]interface{}   `yaml:"paths" mapstructure:"paths"`
	Components   map[string]interface{}   `yaml:"components,omitempty" mapstructure:"components"`
	Security     []map[string]interface{} `yaml:"security,omitempty" mapstructure:"security"`
	Tags         []map[string]interface{} `yaml:"tags,omitempty" mapstructure:"tags"`
	ExternalDocs map[string]interface{}   `yaml:"externalDocs,omitempty" mapstructure:"externalDocs"`
	Examples     map[string]interface{}   `yaml:"examples,omitempty" mapstructure:"examples"`
}

package openapi3

type OpenAPISpec struct {
	OpenAPI      string                 `yaml:"openapi"`
	Info         Info                   `yaml:"info"`
	Servers      []Server               `yaml:"servers,omitempty"`
	Paths        map[string]PathItem    `yaml:"paths"`
	Components   Components             `yaml:"components,omitempty"`
	Security     []SecurityRequirement  `yaml:"security,omitempty"`
	Tags         []Tag                  `yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs,omitempty"`
}

type Info struct {
	Title          string   `yaml:"title"`
	Description    string   `yaml:"description,omitempty"`
	TermsOfService string   `yaml:"termsOfService,omitempty"`
	Contact        *Contact `yaml:"contact,omitempty"`
	License        *License `yaml:"license,omitempty"`
	Version        string   `yaml:"version"`
}

type Contact struct {
	Name  string `yaml:"name,omitempty"`
	URL   string `yaml:"url,omitempty"`
	Email string `yaml:"email,omitempty"`
}

type License struct {
	Name string `yaml:"name,omitempty"`
	URL  string `yaml:"url,omitempty"`
}

type Server struct {
	URL         string                    `yaml:"url"`
	Description string                    `yaml:"description,omitempty"`
	Variables   map[string]ServerVariable `yaml:"variables,omitempty"`
}

type ServerVariable struct {
	Enum        []string `yaml:"enum,omitempty"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description,omitempty"`
}

type PathItem struct {
	Ref         string      `yaml:"$ref,omitempty"`
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	Get         *Operation  `yaml:"get,omitempty"`
	Put         *Operation  `yaml:"put,omitempty"`
	Post        *Operation  `yaml:"post,omitempty"`
	Delete      *Operation  `yaml:"delete,omitempty"`
	Options     *Operation  `yaml:"options,omitempty"`
	Head        *Operation  `yaml:"head,omitempty"`
	Patch       *Operation  `yaml:"patch,omitempty"`
	Trace       *Operation  `yaml:"trace,omitempty"`
	Parameters  []Parameter `yaml:"parameters,omitempty"`
}

type Operation struct {
	Tags         []string               `yaml:"tags,omitempty"`
	Summary      string                 `yaml:"summary,omitempty"`
	Description  string                 `yaml:"description,omitempty"`
	OperationID  string                 `yaml:"operationId,omitempty"`
	Parameters   []Parameter            `yaml:"parameters,omitempty"`
	RequestBody  *RequestBody           `yaml:"requestBody,omitempty"`
	Responses    map[string]Response    `yaml:"responses"`
	Deprecated   bool                   `yaml:"deprecated,omitempty"`
	Security     []SecurityRequirement  `yaml:"security,omitempty"`
	Servers      []Server               `yaml:"servers,omitempty"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs,omitempty"`
}

type Parameter struct {
	Name            string               `yaml:"name"`
	In              string               `yaml:"in"`
	Description     string               `yaml:"description,omitempty"`
	Required        bool                 `yaml:"required"`
	Deprecated      bool                 `yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                 `yaml:"allowEmptyValue,omitempty"`
	Style           string               `yaml:"style,omitempty"`
	Explode         *bool                `yaml:"explode,omitempty"`
	AllowReserved   bool                 `yaml:"allowReserved,omitempty"`
	Schema          *SchemaRef           `yaml:"schema,omitempty"`
	Example         interface{}          `yaml:"example,omitempty"`
	Examples        map[string]Example   `yaml:"examples,omitempty"`
	Content         map[string]MediaType `yaml:"content,omitempty"`
}

type RequestBody struct {
	Description string               `yaml:"description,omitempty"`
	Content     map[string]MediaType `yaml:"content"`
	Required    bool                 `yaml:"required,omitempty"`
}

type Response struct {
	Description string               `yaml:"description"`
	Headers     map[string]Header    `yaml:"headers,omitempty"`
	Content     map[string]MediaType `yaml:"content,omitempty"`
	Links       map[string]Link      `yaml:"links,omitempty"`
}

type MediaType struct {
	Schema   *SchemaRef          `yaml:"schema,omitempty"`
	Example  interface{}         `yaml:"example,omitempty"`
	Examples map[string]Example  `yaml:"examples,omitempty"`
	Encoding map[string]Encoding `yaml:"encoding,omitempty"`
}

type Components struct {
	Schemas         map[string]*Schema         `yaml:"schemas,omitempty"`
	Responses       map[string]*Response       `yaml:"responses,omitempty"`
	Parameters      map[string]*Parameter      `yaml:"parameters,omitempty"`
	Examples        map[string]*Example        `yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `yaml:"requestBodies,omitempty"`
	Headers         map[string]*Header         `yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `yaml:"securitySchemes,omitempty"`
	Links           map[string]*Link           `yaml:"links,omitempty"`
	Callbacks       map[string]*Callback       `yaml:"callbacks,omitempty"`
}

type SchemaRef struct {
	Ref   string  `yaml:"$ref,omitempty"`
	Value *Schema `yaml:"-"` // Avoid recursion in JSON/YAML
}

type Schema struct {
	Title                string                 `yaml:"title,omitempty"`
	MultipleOf           float64                `yaml:"multipleOf,omitempty"`
	Maximum              *float64               `yaml:"maximum,omitempty"`
	ExclusiveMaximum     bool                   `yaml:"exclusiveMaximum,omitempty"`
	Minimum              *float64               `yaml:"minimum,omitempty"`
	ExclusiveMinimum     bool                   `yaml:"exclusiveMinimum,omitempty"`
	MaxLength            *int                   `yaml:"maxLength,omitempty"`
	MinLength            *int                   `yaml:"minLength,omitempty"`
	Pattern              string                 `yaml:"pattern,omitempty"`
	MaxItems             *int                   `yaml:"maxItems,omitempty"`
	MinItems             *int                   `yaml:"minItems,omitempty"`
	UniqueItems          bool                   `yaml:"uniqueItems,omitempty"`
	MaxProperties        *int                   `yaml:"maxProperties,omitempty"`
	MinProperties        *int                   `yaml:"minProperties,omitempty"`
	Required             []string               `yaml:"required,omitempty"`
	Enum                 []interface{}          `yaml:"enum,omitempty"`
	Type                 string                 `yaml:"type,omitempty"`
	AllOf                []*SchemaRef           `yaml:"allOf,omitempty"`
	OneOf                []*SchemaRef           `yaml:"oneOf,omitempty"`
	AnyOf                []*SchemaRef           `yaml:"anyOf,omitempty"`
	Not                  *SchemaRef             `yaml:"not,omitempty"`
	Items                *SchemaRef             `yaml:"items,omitempty"`
	Properties           map[string]*SchemaRef  `yaml:"properties,omitempty"`
	AdditionalProperties *SchemaRef             `yaml:"additionalProperties,omitempty"`
	Description          string                 `yaml:"description,omitempty"`
	Format               string                 `yaml:"format,omitempty"`
	Default              interface{}            `yaml:"default,omitempty"`
	Nullable             bool                   `yaml:"nullable,omitempty"`
	Discriminator        *Discriminator         `yaml:"discriminator,omitempty"`
	ReadOnly             bool                   `yaml:"readOnly,omitempty"`
	WriteOnly            bool                   `yaml:"writeOnly,omitempty"`
	Xml                  *XML                   `yaml:"xml,omitempty"`
	ExternalDocs         *ExternalDocumentation `yaml:"externalDocs,omitempty"`
	Example              interface{}            `yaml:"example,omitempty"`
	Deprecated           bool                   `yaml:"deprecated,omitempty"`
}

type Example struct {
	Summary       string      `yaml:"summary,omitempty"`
	Description   string      `yaml:"description,omitempty"`
	Value         interface{} `yaml:"value,omitempty"`
	ExternalValue string      `yaml:"externalValue,omitempty"`
}

type Encoding struct {
	ContentType   string             `yaml:"contentType,omitempty"`
	Headers       map[string]*Header `yaml:"headers,omitempty"`
	Style         string             `yaml:"style,omitempty"`
	Explode       bool               `yaml:"explode,omitempty"`
	AllowReserved bool               `yaml:"allowReserved,omitempty"`
}

type Header struct {
	Description     string                `yaml:"description,omitempty"`
	Required        bool                  `yaml:"required,omitempty"`
	Deprecated      bool                  `yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                  `yaml:"allowEmptyValue,omitempty"`
	Style           string                `yaml:"style,omitempty"`
	Explode         *bool                 `yaml:"explode,omitempty"`
	AllowReserved   bool                  `yaml:"allowReserved,omitempty"`
	Schema          *SchemaRef            `yaml:"schema,omitempty"`
	Example         interface{}           `yaml:"example,omitempty"`
	Examples        map[string]*Example   `yaml:"examples,omitempty"`
	Content         map[string]*MediaType `yaml:"content,omitempty"`
}

type Link struct {
	OperationRef string                 `yaml:"operationRef,omitempty"`
	OperationID  string                 `yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `yaml:"parameters,omitempty"`
	RequestBody  interface{}            `yaml:"requestBody,omitempty"`
	Description  string                 `yaml:"description,omitempty"`
	Server       *Server                `yaml:"server,omitempty"`
}

type Callback map[string]*PathItem

type SecurityRequirement map[string][]string

type SecurityScheme struct {
	Type             string      `yaml:"type"`
	Description      string      `yaml:"description,omitempty"`
	Name             string      `yaml:"name,omitempty"`
	In               string      `yaml:"in,omitempty"`
	Scheme           string      `yaml:"scheme,omitempty"`
	BearerFormat     string      `yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `yaml:"flows,omitempty"`
	OpenIDConnectUrl string      `yaml:"openIdConnectUrl,omitempty"`
}

type OAuthFlows struct {
	Implicit          *OAuthFlow `yaml:"implicit,omitempty"`
	Password          *OAuthFlow `yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode,omitempty"`
}

type OAuthFlow struct {
	AuthorizationUrl string            `yaml:"authorizationUrl,omitempty"`
	TokenUrl         string            `yaml:"tokenUrl,omitempty"`
	RefreshUrl       string            `yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `yaml:"scopes,omitempty"`
}

type Tag struct {
	Name         string                 `yaml:"name"`
	Description  string                 `yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs,omitempty"`
}

type ExternalDocumentation struct {
	Description string `yaml:"description,omitempty"`
	URL         string `yaml:"url"`
}

type Discriminator struct {
	PropertyName string            `yaml:"propertyName"`
	Mapping      map[string]string `yaml:"mapping,omitempty"`
}

type XML struct {
	Name      string `yaml:"name,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	Prefix    string `yaml:"prefix,omitempty"`
	Attribute bool   `yaml:"attribute,omitempty"`
	Wrapped   bool   `yaml:"wrapped,omitempty"`
}

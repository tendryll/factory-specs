package v3_1

import (
	"encoding/json"
)

// OpenAPI represents the root OpenAPI 3.1 document
type OpenAPI struct {
	OpenAPI           string                 `json:"openapi" yaml:"openapi"`
	Info              Info                   `json:"info" yaml:"info"`
	JSONSchemaDialect *string                `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             *Paths                 `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          map[string]*PathItem   `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Security          []SecurityRequirement  `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              []Tag                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Extensions        map[string]interface{} `json:"-" yaml:",inline"`
}

// Info represents the metadata about the API
type Info struct {
	Title          string                 `json:"title" yaml:"title"`
	Summary        *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    *string                `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService *string                `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact               `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License               `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string                 `json:"version" yaml:"version"`
	Extensions     map[string]interface{} `json:"-" yaml:",inline"`
}

// Contact represents the contact information for the API
type Contact struct {
	Name       *string                `json:"name,omitempty" yaml:"name,omitempty"`
	URL        *string                `json:"url,omitempty" yaml:"url,omitempty"`
	Email      *string                `json:"email,omitempty" yaml:"email,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// License represents the license information for the API
type License struct {
	Name       string                 `json:"name" yaml:"name"`
	Identifier *string                `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        *string                `json:"url,omitempty" yaml:"url,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// Server represents a server configuration
type Server struct {
	URL         string                     `json:"url" yaml:"url"`
	Description *string                    `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
	Extensions  map[string]interface{}     `json:"-" yaml:",inline"`
}

// ServerVariable represents a server variable for templating
type ServerVariable struct {
	Enum        []string               `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string                 `json:"default" yaml:"default"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// Components holds reusable objects for different aspects of the OAS
type Components struct {
	Schemas         map[string]*Schema            `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*ResponseRef       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*ParameterRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*ExampleRef        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBodyRef    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*HeaderRef         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeRef `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*LinkRef           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*CallbacksRef      `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       map[string]*PathItem          `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
	Extensions      map[string]interface{}        `json:"-" yaml:",inline"`
}

// Paths holds the relative paths to the individual endpoints
type Paths struct {
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
	Paths      map[string]*PathItem   `json:"-" yaml:",inline"`
}

// PathItem describes the operations available on a single path
type PathItem struct {
	Ref         *string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Servers     []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []*ParameterRef        `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Get         *Operation             `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation             `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation             `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation             `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation             `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation             `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation             `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation             `json:"trace,omitempty" yaml:"trace,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// Operation describes a single API operation on a path
type Operation struct {
	Tags         []string                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      *string                  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  *string                  `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  *string                  `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*ParameterRef          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBodyRef          `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    *Responses               `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    map[string]*CallbacksRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   *bool                    `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []SecurityRequirement    `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []Server                 `json:"servers,omitempty" yaml:"servers,omitempty"`
	Extensions   map[string]interface{}   `json:"-" yaml:",inline"`
}

// ExternalDocumentation allows referencing an external resource for extended documentation
type ExternalDocumentation struct {
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string                 `json:"url" yaml:"url"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// Parameter describes a single operation parameter
type Parameter struct {
	Name            string                 `json:"name" yaml:"name"`
	In              string                 `json:"in" yaml:"in"` // "query", "header", "path", "cookie"
	Description     *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Required        *bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      *bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue *bool                  `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           *string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   *bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *Schema                `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content         map[string]*MediaType  `json:"content,omitempty" yaml:"content,omitempty"`
	Example         interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Extensions      map[string]interface{} `json:"-" yaml:",inline"`
}

// ParameterRef represents either a Parameter or a Reference
type ParameterRef struct {
	Ref   *string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Parameter `json:"-" yaml:"-"`
}

// RequestBody describes a single request body
type RequestBody struct {
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType  `json:"content" yaml:"content"`
	Required    *bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// RequestBodyRef represents either a RequestBody or a Reference
type RequestBodyRef struct {
	Ref   *string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *RequestBody `json:"-" yaml:"-"`
}

// MediaType provides schema and examples for the media type identified by its key
type MediaType struct {
	Schema     *Schema                `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example    interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples   map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding   map[string]*Encoding   `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// Encoding defines encoding information for request body media types
type Encoding struct {
	ContentType   *string                `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderRef  `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         *string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved *bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Extensions    map[string]interface{} `json:"-" yaml:",inline"`
}

// Responses is a container for the expected responses of an operation
type Responses struct {
	Default    *ResponseRef            `json:"default,omitempty" yaml:"default,omitempty"`
	Responses  map[string]*ResponseRef `json:"-" yaml:",inline"`
	Extensions map[string]interface{}  `json:"-" yaml:",inline"`
}

// Response describes a single response from an API Operation
type Response struct {
	Description string                 `json:"description" yaml:"description"`
	Headers     map[string]*HeaderRef  `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType  `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef    `json:"links,omitempty" yaml:"links,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// ResponseRef represents either a Response or a Reference
type ResponseRef struct {
	Ref   *string   `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Response `json:"-" yaml:"-"`
}

// Callbacks is a map of possible out-of band callbacks related to the parent operation
type Callbacks map[string]*PathItem

// CallbacksRef represents either Callbacks or a Reference
type CallbacksRef struct {
	Ref   *string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Callbacks `json:"-" yaml:"-"`
}

// Example object
type Example struct {
	Summary       *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{}            `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue *string                `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
	Extensions    map[string]interface{} `json:"-" yaml:",inline"`
}

// ExampleRef represents either an Example or a Reference
type ExampleRef struct {
	Ref   *string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Example `json:"-" yaml:"-"`
}

// Link represents a possible design-time link for a response
type Link struct {
	OperationRef *string                `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  *string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]string      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

// LinkRef represents either a Link or a Reference
type LinkRef struct {
	Ref   *string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Link   `json:"-" yaml:"-"`
}

// Header represents a header parameter
type Header struct {
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Required    *bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated  *bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Style       *string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode     *bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	Schema      *Schema                `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content     map[string]*MediaType  `json:"content,omitempty" yaml:"content,omitempty"`
	Example     interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples    map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// HeaderRef represents either a Header or a Reference
type HeaderRef struct {
	Ref   *string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Header `json:"-" yaml:"-"`
}

// Tag adds metadata to a single tag that is used by the Operation Object
type Tag struct {
	Name         string                 `json:"name" yaml:"name"`
	Description  *string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

// Reference represents a simple object to allow referencing other components
type Reference struct {
	Ref         string  `json:"$ref" yaml:"$ref"`
	Summary     *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

// Schema represents a schema object (can be boolean or object in OpenAPI 3.1)
type Schema struct {
	// Core schema properties
	Type        interface{}   `json:"type,omitempty" yaml:"type,omitempty"` // can be string or array of strings
	Format      *string       `json:"format,omitempty" yaml:"format,omitempty"`
	Title       *string       `json:"title,omitempty" yaml:"title,omitempty"`
	Description *string       `json:"description,omitempty" yaml:"description,omitempty"`
	Default     interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Example     interface{}   `json:"example,omitempty" yaml:"example,omitempty"`
	Examples    []interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`

	// Validation properties
	MultipleOf       *float64      `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          *float64      `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum interface{}   `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum interface{}   `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        *int          `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        *int          `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          *string       `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         *int          `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         *int          `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      *bool         `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties    *int          `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties    *int          `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required         []string      `json:"required,omitempty" yaml:"required,omitempty"`
	Enum             []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`

	// Object properties
	Properties           map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties interface{}        `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	PatternProperties    map[string]*Schema `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`

	// Array properties
	Items       interface{} `json:"items,omitempty" yaml:"items,omitempty"`
	PrefixItems []*Schema   `json:"prefixItems,omitempty" yaml:"prefixItems,omitempty"`
	Contains    *Schema     `json:"contains,omitempty" yaml:"contains,omitempty"`

	// Composition
	AllOf []*Schema `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf []*Schema `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf []*Schema `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not   *Schema   `json:"not,omitempty" yaml:"not,omitempty"`

	// Conditional
	If   *Schema `json:"if,omitempty" yaml:"if,omitempty"`
	Then *Schema `json:"then,omitempty" yaml:"then,omitempty"`
	Else *Schema `json:"else,omitempty" yaml:"else,omitempty"`

	// OpenAPI specific
	Discriminator *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	XML           *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Deprecated    *bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	ReadOnly      *bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     *bool                  `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`

	// Extensions
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// Discriminator is used to differentiate between schemas
type Discriminator struct {
	PropertyName string                 `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string      `json:"mapping,omitempty" yaml:"mapping,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

// XML provides additional metadata for XML serialization
type XML struct {
	Name       *string                `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace  *string                `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix     *string                `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute  *bool                  `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped    *bool                  `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// SecurityScheme defines a security scheme that can be used by the operations
type SecurityScheme struct {
	Type             string                 `json:"type" yaml:"type"` // "apiKey", "http", "mutualTLS", "oauth2", "openIdConnect"
	Description      *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Name             *string                `json:"name,omitempty" yaml:"name,omitempty"`
	In               *string                `json:"in,omitempty" yaml:"in,omitempty"` // "query", "header", "cookie"
	Scheme           *string                `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     *string                `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows            `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectURL *string                `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
	Extensions       map[string]interface{} `json:"-" yaml:",inline"`
}

// SecuritySchemeRef represents either a SecurityScheme or a Reference
type SecuritySchemeRef struct {
	Ref   *string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *SecurityScheme `json:"-" yaml:"-"`
}

// OAuthFlows allows configuration of the supported OAuth Flows
type OAuthFlows struct {
	Implicit          *OAuthFlow             `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow             `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow             `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow             `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
	Extensions        map[string]interface{} `json:"-" yaml:",inline"`
}

// OAuthFlow configuration details for a supported OAuth Flow
type OAuthFlow struct {
	AuthorizationURL *string                `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         *string                `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       *string                `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes" yaml:"scopes"`
	Extensions       map[string]interface{} `json:"-" yaml:",inline"`
}

// SecurityRequirement lists the required security schemes to execute this operation
type SecurityRequirement map[string][]string

// Custom JSON marshaling/unmarshaling methods for handling extensions and references

func (o *OpenAPI) MarshalJSON() ([]byte, error) {
	type Alias OpenAPI
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}

	b, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}

	if o.Extensions == nil {
		return b, nil
	}

	var base map[string]interface{}
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}

	for k, v := range o.Extensions {
		base[k] = v
	}

	return json.Marshal(base)
}

func (p *ParameterRef) UnmarshalJSON(data []byte) error {
	var ref struct {
		Ref string `json:"$ref"`
	}
	if err := json.Unmarshal(data, &ref); err == nil && ref.Ref != "" {
		p.Ref = &ref.Ref
		return nil
	}

	var param Parameter
	if err := json.Unmarshal(data, &param); err != nil {
		return err
	}
	p.Value = &param
	return nil
}

func (p *ParameterRef) MarshalJSON() ([]byte, error) {
	if p.Ref != nil {
		return json.Marshal(struct {
			Ref string `json:"$ref"`
		}{Ref: *p.Ref})
	}
	return json.Marshal(p.Value)
}

// Similar marshal/unmarshal methods would be needed for other *Ref types
// (ResponseRef, RequestBodyRef, ExampleRef, HeaderRef, SecuritySchemeRef, LinkRef, CallbacksRef)
// Following the same pattern as ParameterRef

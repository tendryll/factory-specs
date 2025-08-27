package v3_0

// OpenAPI represents the root OpenAPI 3.0 document
type OpenAPI struct {
	OpenAPI      string                 `json:"openapi" yaml:"openapi"`
	Info         Info                   `json:"info" yaml:"info"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Servers      []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Security     []SecurityRequirement  `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []Tag                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	Paths        Paths                  `json:"paths" yaml:"paths"`
	Components   *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

// Info represents the metadata about the API
type Info struct {
	Title          string                 `json:"title" yaml:"title"`
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
	Schemas         map[string]*SchemaOrRef         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*ResponseOrRef       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*ParameterOrRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*ExampleOrRef        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBodyOrRef    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*HeaderOrRef         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeOrRef `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*LinkOrRef           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*CallbackOrRef       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Extensions      map[string]interface{}          `json:"-" yaml:",inline"`
}

// Paths holds the relative paths to the individual endpoints
type Paths struct {
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
	Items      map[string]*PathItem   `json:"paths" yaml:"paths"`
}

// PathItem describes the operations available on a single path
type PathItem struct {
	Ref         *string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation             `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation             `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation             `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation             `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation             `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation             `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation             `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation             `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []*ParameterOrRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// Operation describes a single API operation on a path
type Operation struct {
	Tags         []string                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      *string                   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  *string                   `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation    `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  *string                   `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*ParameterOrRef         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBodyOrRef         `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    Responses                 `json:"responses" yaml:"responses"`
	Callbacks    map[string]*CallbackOrRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   *bool                     `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []SecurityRequirement     `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []Server                  `json:"servers,omitempty" yaml:"servers,omitempty"`
	Extensions   map[string]interface{}    `json:"-" yaml:",inline"`
}

// ExternalDocumentation allows referencing an external resource for extended documentation
type ExternalDocumentation struct {
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string                 `json:"url" yaml:"url"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// Parameter describes a single operation parameter
type Parameter struct {
	Name            string                   `json:"name" yaml:"name"`
	In              string                   `json:"in" yaml:"in"` // "query", "header", "path", "cookie"
	Description     *string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Required        *bool                    `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      *bool                    `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue *bool                    `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           *string                  `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                    `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   *bool                    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content         map[string]*MediaType    `json:"content,omitempty" yaml:"content,omitempty"`
	Example         interface{}              `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Extensions      map[string]interface{}   `json:"-" yaml:",inline"`
}

// RequestBody describes a single request body
type RequestBody struct {
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType  `json:"content" yaml:"content"`
	Required    *bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

// MediaType provides schema and examples for the media type identified by its key
type MediaType struct {
	Schema     *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example    interface{}              `json:"example,omitempty" yaml:"example,omitempty"`
	Examples   map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding   map[string]*Encoding     `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	Extensions map[string]interface{}   `json:"-" yaml:",inline"`
}

// Encoding defines encoding information for request body media types
type Encoding struct {
	ContentType   *string                 `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderOrRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         *string                 `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                   `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved *bool                   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Extensions    map[string]interface{}  `json:"-" yaml:",inline"`
}

// Responses is a container for the expected responses of an operation
type Responses struct {
	Default    *ResponseOrRef            `json:"default,omitempty" yaml:"default,omitempty"`
	Responses  map[string]*ResponseOrRef `json:"-" yaml:",inline"`
	Extensions map[string]interface{}    `json:"-" yaml:",inline"`
}

// Response describes a single response from an API Operation
type Response struct {
	Description string                  `json:"description" yaml:"description"`
	Headers     map[string]*HeaderOrRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType   `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkOrRef   `json:"links,omitempty" yaml:"links,omitempty"`
	Extensions  map[string]interface{}  `json:"-" yaml:",inline"`
}

// Example object
type Example struct {
	Summary       *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{}            `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue *string                `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
	Extensions    map[string]interface{} `json:"-" yaml:",inline"`
}

// Link represents a possible design-time link for a response
type Link struct {
	OperationRef *string                `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  *string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

// Header represents a header parameter
type Header struct {
	Description     *string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Required        *bool                    `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      *bool                    `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue *bool                    `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           *string                  `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                    `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   *bool                    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content         map[string]*MediaType    `json:"content,omitempty" yaml:"content,omitempty"`
	Example         interface{}              `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Extensions      map[string]interface{}   `json:"-" yaml:",inline"`
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
	Ref string `json:"$ref" yaml:"$ref"`
}

// Schema represents a schema object for OpenAPI 3.0 (JSON Schema subset)
type Schema struct {
	// Core schema properties
	Title       *string     `json:"title,omitempty" yaml:"title,omitempty"`
	Type        *string     `json:"type,omitempty" yaml:"type,omitempty"` // "array", "boolean", "integer", "number", "object", "string"
	Format      *string     `json:"format,omitempty" yaml:"format,omitempty"`
	Description *string     `json:"description,omitempty" yaml:"description,omitempty"`
	Default     interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Example     interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Validation properties
	MultipleOf       *float64      `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          *float64      `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum *bool         `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum *bool         `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
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
	Properties           map[string]*SchemaOrRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties interface{}             `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"` // can be bool or Schema

	// Array properties
	Items *SchemaOrRef `json:"items,omitempty" yaml:"items,omitempty"`

	// Composition
	AllOf []*SchemaOrRef `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf []*SchemaOrRef `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf []*SchemaOrRef `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not   *SchemaOrRef   `json:"not,omitempty" yaml:"not,omitempty"`

	// OpenAPI 3.0 specific
	Nullable      *bool                  `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Discriminator *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly      *bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     *bool                  `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	XML           *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Deprecated    *bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Extensions
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// Discriminator is used to differentiate between schemas
type Discriminator struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// XML provides additional metadata for XML serialization
type XML struct {
	Name      *string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    *string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute *bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   *bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

// SecurityScheme defines a security scheme that can be used by the operations
type SecurityScheme interface {
	SecuritySchemeType() string
}

// APIKeySecurityScheme represents API Key security
type APIKeySecurityScheme struct {
	Type        string                 `json:"type" yaml:"type"` // "apiKey"
	Name        string                 `json:"name" yaml:"name"`
	In          string                 `json:"in" yaml:"in"` // "header", "query", "cookie"
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

func (a *APIKeySecurityScheme) SecuritySchemeType() string { return "apiKey" }

// HTTPSecurityScheme represents HTTP security
type HTTPSecurityScheme struct {
	Type         string                 `json:"type" yaml:"type"` // "http"
	Scheme       string                 `json:"scheme" yaml:"scheme"`
	BearerFormat *string                `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Description  *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:",inline"`
}

func (h *HTTPSecurityScheme) SecuritySchemeType() string { return "http" }

// OAuth2SecurityScheme represents OAuth2 security
type OAuth2SecurityScheme struct {
	Type        string                 `json:"type" yaml:"type"` // "oauth2"
	Flows       OAuthFlows             `json:"flows" yaml:"flows"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions  map[string]interface{} `json:"-" yaml:",inline"`
}

func (o *OAuth2SecurityScheme) SecuritySchemeType() string { return "oauth2" }

// OpenIDConnectSecurityScheme represents OpenID Connect security
type OpenIDConnectSecurityScheme struct {
	Type             string                 `json:"type" yaml:"type"` // "openIdConnect"
	OpenIDConnectURL string                 `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
	Description      *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions       map[string]interface{} `json:"-" yaml:",inline"`
}

func (o *OpenIDConnectSecurityScheme) SecuritySchemeType() string { return "openIdConnect" }

// OAuthFlows allows configuration of the supported OAuth Flows
type OAuthFlows struct {
	Implicit          *ImplicitOAuthFlow          `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *PasswordOAuthFlow          `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *ClientCredentialsOAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *AuthorizationCodeOAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
	Extensions        map[string]interface{}      `json:"-" yaml:",inline"`
}

// ImplicitOAuthFlow configuration for implicit OAuth flow
type ImplicitOAuthFlow struct {
	AuthorizationURL string                 `json:"authorizationUrl" yaml:"authorizationUrl"`
	RefreshURL       *string                `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes" yaml:"scopes"`
	Extensions       map[string]interface{} `json:"-" yaml:",inline"`
}

// PasswordOAuthFlow configuration for password OAuth flow
type PasswordOAuthFlow struct {
	TokenURL   string                 `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL *string                `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes     map[string]string      `json:"scopes" yaml:"scopes"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// ClientCredentialsOAuthFlow configuration for client credentials OAuth flow
type ClientCredentialsOAuthFlow struct {
	TokenURL   string                 `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL *string                `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes     map[string]string      `json:"scopes" yaml:"scopes"`
	Extensions map[string]interface{} `json:"-" yaml:",inline"`
}

// AuthorizationCodeOAuthFlow configuration for authorization code OAuth flow
type AuthorizationCodeOAuthFlow struct {
	AuthorizationURL string                 `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenURL         string                 `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL       *string                `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes" yaml:"scopes"`
	Extensions       map[string]interface{} `json:"-" yaml:",inline"`
}

// SecurityRequirement lists the required security schemes to execute this operation
type SecurityRequirement map[string][]string

// Callback is a map of possible out-of band callbacks related to the parent operation
type Callback map[string]*PathItem

// Reference wrapper types for handling oneOf relationships

// SchemaOrRef represents either a Schema or a Reference
type SchemaOrRef struct {
	Schema *Schema    `json:"-"`
	Ref    *Reference `json:"-"`
}

// ParameterOrRef represents either a Parameter or a Reference
type ParameterOrRef struct {
	Parameter *Parameter `json:"-"`
	Ref       *Reference `json:"-"`
}

// ResponseOrRef represents either a Response or a Reference
type ResponseOrRef struct {
	Response *Response  `json:"-"`
	Ref      *Reference `json:"-"`
}

// RequestBodyOrRef represents either a RequestBody or a Reference
type RequestBodyOrRef struct {
	RequestBody *RequestBody `json:"-"`
	Ref         *Reference   `json:"-"`
}

// ExampleOrRef represents either an Example or a Reference
type ExampleOrRef struct {
	Example *Example   `json:"-"`
	Ref     *Reference `json:"-"`
}

// HeaderOrRef represents either a Header or a Reference
type HeaderOrRef struct {
	Header *Header    `json:"-"`
	Ref    *Reference `json:"-"`
}

// SecuritySchemeOrRef represents either a SecurityScheme or a Reference
type SecuritySchemeOrRef struct {
	SecurityScheme SecurityScheme `json:"-"`
	Ref            *Reference     `json:"-"`
}

// LinkOrRef represents either a Link or a Reference
type LinkOrRef struct {
	Link *Link      `json:"-"`
	Ref  *Reference `json:"-"`
}

// CallbackOrRef represents either a Callback or a Reference
type CallbackOrRef struct {
	Callback *Callback  `json:"-"`
	Ref      *Reference `json:"-"`
}

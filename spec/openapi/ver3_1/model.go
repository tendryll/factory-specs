// Code generated from OpenAPI v3.1.1 JSON Schema (v3.1.1.json). DO NOT EDIT.

package ver3_1

//import (
//	"time"
//)

// Callbacks - generated from $defs/callbacks

// Document wraps the root OpenAPI struct generated from v3.1.1 schema.
type Document struct {
	Spec OpenAPI
}

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

type Callbacks struct {
	AdditionalProperties map[string]PathItem    `json:"-"`
	Extensions           map[string]interface{} `json:"-"`
}

// CallbacksOrReference - fallback to interface{}

type CallbacksOrReference interface{}

// Components - generated from $defs/components

type Components struct {
	Callbacks       map[string]CallbacksOrReference      `json:"callbacks,omitempty"`
	Examples        map[string]ExampleOrReference        `json:"examples,omitempty"`
	Headers         map[string]HeaderOrReference         `json:"headers,omitempty"`
	Links           map[string]LinkOrReference           `json:"links,omitempty"`
	Parameters      map[string]ParameterOrReference      `json:"parameters,omitempty"`
	Pathitems       map[string]PathItem                  `json:"pathItems,omitempty"`
	Requestbodies   map[string]RequestBodyOrReference    `json:"requestBodies,omitempty"`
	Responses       map[string]ResponseOrReference       `json:"responses,omitempty"`
	Schemas         map[string]interface{}               `json:"schemas,omitempty"`
	Securityschemes map[string]SecuritySchemeOrReference `json:"securitySchemes,omitempty"`
	Extensions      map[string]interface{}               `json:"-"`
}

// Contact - generated from $defs/contact

type Contact struct {
	Email      *string                `json:"email,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Url        *string                `json:"url,omitempty"`
	Extensions map[string]interface{} `json:"-"`
}

// Content - generated from $defs/content

type Content struct {
	AdditionalProperties map[string]MediaType   `json:"-"`
	Extensions           map[string]interface{} `json:"-"`
}

// Encoding - generated from $defs/encoding

type Encoding struct {
	Allowreserved *bool                        `json:"allowReserved,omitempty"`
	Contenttype   *string                      `json:"contentType,omitempty"`
	Explode       *bool                        `json:"explode,omitempty"`
	Headers       map[string]HeaderOrReference `json:"headers,omitempty"`
	Style         *string                      `json:"style,omitempty"`
	Extensions    map[string]interface{}       `json:"-"`
}

// Example - generated from $defs/example

type Example struct {
	Description   *string                `json:"description,omitempty"`
	Externalvalue *string                `json:"externalValue,omitempty"`
	Summary       *string                `json:"summary,omitempty"`
	Value         interface{}            `json:"value,omitempty"`
	Extensions    map[string]interface{} `json:"-"`
}

// ExampleOrReference - fallback to interface{}

type ExampleOrReference interface{}

// Examples - generated from $defs/examples

type Examples struct {
	Example    interface{}                   `json:"example,omitempty"`
	Examples   map[string]ExampleOrReference `json:"examples,omitempty"`
	Extensions map[string]interface{}        `json:"-"`
}

// ExternalDocumentation - generated from $defs/external-documentation

type ExternalDocumentation struct {
	Description *string                `json:"description,omitempty"`
	Url         string                 `json:"url"`
	Extensions  map[string]interface{} `json:"-"`
}

// Header - generated from $defs/header

type Header struct {
	Content     *Content               `json:"content,omitempty"`
	Deprecated  *bool                  `json:"deprecated,omitempty"`
	Description *string                `json:"description,omitempty"`
	Required    *bool                  `json:"required,omitempty"`
	Schema      interface{}            `json:"schema,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// HeaderOrReference - fallback to interface{}

type HeaderOrReference interface{}

// Info - generated from $defs/info

type Info struct {
	Contact        *Contact               `json:"contact,omitempty"`
	Description    *string                `json:"description,omitempty"`
	License        *License               `json:"license,omitempty"`
	Summary        *string                `json:"summary,omitempty"`
	Termsofservice *string                `json:"termsOfService,omitempty"`
	Title          string                 `json:"title"`
	Version        string                 `json:"version"`
	Extensions     map[string]interface{} `json:"-"`
}

// License - generated from $defs/license

type License struct {
	Identifier *string                `json:"identifier,omitempty"`
	Name       string                 `json:"name"`
	Url        *string                `json:"url,omitempty"`
	Extensions map[string]interface{} `json:"-"`
}

// Link - generated from $defs/link

type Link struct {
	Body         *Server                `json:"body,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Operationid  *string                `json:"operationId,omitempty"`
	Operationref *string                `json:"operationRef,omitempty"`
	Parameters   *MapOfStrings          `json:"parameters,omitempty"`
	Requestbody  interface{}            `json:"requestBody,omitempty"`
	Extensions   map[string]interface{} `json:"-"`
}

// LinkOrReference - fallback to interface{}

type LinkOrReference interface{}

// MapOfStrings - generated from $defs/map-of-strings

type MapOfStrings struct {
	AdditionalProperties map[string]string      `json:"-"`
	Extensions           map[string]interface{} `json:"-"`
}

// MediaType - generated from $defs/media-type

type MediaType struct {
	Encoding   map[string]Encoding    `json:"encoding,omitempty"`
	Schema     interface{}            `json:"schema,omitempty"`
	Extensions map[string]interface{} `json:"-"`
}

// OauthFlows - generated from $defs/oauth-flows

type OauthFlows struct {
	Authorizationcode *OauthFlowsAuthorizationCode `json:"authorizationCode,omitempty"`
	Clientcredentials *OauthFlowsClientCredentials `json:"clientCredentials,omitempty"`
	Implicit          *OauthFlowsImplicit          `json:"implicit,omitempty"`
	Password          *OauthFlowsPassword          `json:"password,omitempty"`
	Extensions        map[string]interface{}       `json:"-"`
}

// OauthFlowsAuthorizationCode - generated from $defs/oauth-flows/$defs/authorization-code

type OauthFlowsAuthorizationCode struct {
	Authorizationurl string                 `json:"authorizationUrl"`
	Refreshurl       *string                `json:"refreshUrl,omitempty"`
	Scopes           MapOfStrings           `json:"scopes"`
	Tokenurl         string                 `json:"tokenUrl"`
	Extensions       map[string]interface{} `json:"-"`
}

// OauthFlowsClientCredentials - generated from $defs/oauth-flows/$defs/client-credentials

type OauthFlowsClientCredentials struct {
	Refreshurl *string                `json:"refreshUrl,omitempty"`
	Scopes     MapOfStrings           `json:"scopes"`
	Tokenurl   string                 `json:"tokenUrl"`
	Extensions map[string]interface{} `json:"-"`
}

// OauthFlowsImplicit - generated from $defs/oauth-flows/$defs/implicit

type OauthFlowsImplicit struct {
	Authorizationurl string                 `json:"authorizationUrl"`
	Refreshurl       *string                `json:"refreshUrl,omitempty"`
	Scopes           MapOfStrings           `json:"scopes"`
	Extensions       map[string]interface{} `json:"-"`
}

// OauthFlowsPassword - generated from $defs/oauth-flows/$defs/password

type OauthFlowsPassword struct {
	Refreshurl *string                `json:"refreshUrl,omitempty"`
	Scopes     MapOfStrings           `json:"scopes"`
	Tokenurl   string                 `json:"tokenUrl"`
	Extensions map[string]interface{} `json:"-"`
}

// Operation - generated from $defs/operation

type Operation struct {
	Callbacks    map[string]CallbacksOrReference `json:"callbacks,omitempty"`
	Deprecated   *bool                           `json:"deprecated,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	Externaldocs *ExternalDocumentation          `json:"externalDocs,omitempty"`
	Operationid  *string                         `json:"operationId,omitempty"`
	Parameters   []ParameterOrReference          `json:"parameters,omitempty"`
	Requestbody  *RequestBodyOrReference         `json:"requestBody,omitempty"`
	Responses    *Responses                      `json:"responses,omitempty"`
	Security     []SecurityRequirement           `json:"security,omitempty"`
	Servers      []Server                        `json:"servers,omitempty"`
	Summary      *string                         `json:"summary,omitempty"`
	Tags         []string                        `json:"tags,omitempty"`
	Extensions   map[string]interface{}          `json:"-"`
}

// Parameter - generated from $defs/parameter

type Parameter struct {
	Content     *Content               `json:"content,omitempty"`
	Deprecated  *bool                  `json:"deprecated,omitempty"`
	Description *string                `json:"description,omitempty"`
	In          string                 `json:"in"`
	Name        string                 `json:"name"`
	Required    *bool                  `json:"required,omitempty"`
	Schema      interface{}            `json:"schema,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// ParameterOrReference - fallback to interface{}

type ParameterOrReference interface{}

// PathItem - generated from $defs/path-item

type PathItem struct {
	Ref         *string                `json:"$ref,omitempty"`
	Delete      *Operation             `json:"delete,omitempty"`
	Description *string                `json:"description,omitempty"`
	Get         *Operation             `json:"get,omitempty"`
	Head        *Operation             `json:"head,omitempty"`
	Options     *Operation             `json:"options,omitempty"`
	Parameters  []ParameterOrReference `json:"parameters,omitempty"`
	Patch       *Operation             `json:"patch,omitempty"`
	Post        *Operation             `json:"post,omitempty"`
	Put         *Operation             `json:"put,omitempty"`
	Servers     []Server               `json:"servers,omitempty"`
	Summary     *string                `json:"summary,omitempty"`
	Trace       *Operation             `json:"trace,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// Paths - generated from $defs/paths

type Paths struct {
	Extensions map[string]interface{} `json:"-"`
}

// Reference - generated from $defs/reference

type Reference struct {
	Ref         *string                `json:"$ref,omitempty"`
	Description *string                `json:"description,omitempty"`
	Summary     *string                `json:"summary,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// RequestBody - generated from $defs/request-body

type RequestBody struct {
	Content     Content                `json:"content"`
	Description *string                `json:"description,omitempty"`
	Required    *bool                  `json:"required,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// RequestBodyOrReference - fallback to interface{}

type RequestBodyOrReference interface{}

// Response - generated from $defs/response

type Response struct {
	Content     *Content                     `json:"content,omitempty"`
	Description string                       `json:"description"`
	Headers     map[string]HeaderOrReference `json:"headers,omitempty"`
	Links       map[string]LinkOrReference   `json:"links,omitempty"`
	Extensions  map[string]interface{}       `json:"-"`
}

// ResponseOrReference - fallback to interface{}

type ResponseOrReference interface{}

// Responses - generated from $defs/responses

type Responses struct {
	Default    *ResponseOrReference   `json:"default,omitempty"`
	Extensions map[string]interface{} `json:"-"`
}

// Schema - fallback to interface{}

type Schema interface{}

// SecurityRequirement - generated from $defs/security-requirement

type SecurityRequirement struct {
	AdditionalProperties map[string][]string    `json:"-"`
	Extensions           map[string]interface{} `json:"-"`
}

// SecurityScheme - generated from $defs/security-scheme

type SecurityScheme struct {
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Extensions  map[string]interface{} `json:"-"`
}

// SecuritySchemeOrReference - fallback to interface{}

type SecuritySchemeOrReference interface{}

// SecuritySchemeTypeApikey - fallback to interface{}

type SecuritySchemeTypeApikey interface{}

// SecuritySchemeTypeHttp - fallback to interface{}

type SecuritySchemeTypeHttp interface{}

// SecuritySchemeTypeHttpBearer - fallback to interface{}

type SecuritySchemeTypeHttpBearer interface{}

// SecuritySchemeTypeOauth2 - fallback to interface{}

type SecuritySchemeTypeOauth2 interface{}

// SecuritySchemeTypeOidc - fallback to interface{}

type SecuritySchemeTypeOidc interface{}

// Server - generated from $defs/server

type Server struct {
	Description *string                   `json:"description,omitempty"`
	Url         string                    `json:"url"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
	Extensions  map[string]interface{}    `json:"-"`
}

// ServerVariable - generated from $defs/server-variable

type ServerVariable struct {
	Default     string                 `json:"default"`
	Description *string                `json:"description,omitempty"`
	Enum        []string               `json:"enum,omitempty"`
	Extensions  map[string]interface{} `json:"-"`
}

// SpecificationExtensions - fallback to interface{}

type SpecificationExtensions interface{}

// StylesForForm - fallback to interface{}

type StylesForForm interface{}

// Tag - generated from $defs/tag

type Tag struct {
	Description  *string                `json:"description,omitempty"`
	Externaldocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	Name         string                 `json:"name"`
	Extensions   map[string]interface{} `json:"-"`
}

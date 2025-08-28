package ver3_0

// Document wraps an OpenAPI 3.0.4 specification document.
type Document struct {
	Spec OpenAPI
}

type OpenAPI struct {
	OpenAPI      string                `json:"openapi"`
	Info         Info                  `json:"info"`
	Servers      []Server              `json:"servers,omitempty"`
	Paths        *Paths                `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   Components            `json:"components,omitempty"`
	Security     []Securityrequirement `json:"security,omitempty"`
	Tags         []Tag                 `json:"tags,omitempty"`
	ExternalDocs Externaldocumentation `json:"externalDocs,omitempty"`
}

type Reference struct {
}

type Info struct {
	Title          string  `json:"title,omitempty"`
	Description    string  `json:"description,omitempty"`
	Termsofservice string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Server struct {
	Url         string                    `json:"url,omitempty"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]Servervariable `json:"variables,omitempty"`
}

type Servervariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Components struct {
	Schemas         map[string]interface{} `json:"schemas,omitempty"`
	Responses       map[string]interface{} `json:"responses,omitempty"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
	Examples        map[string]interface{} `json:"examples,omitempty"`
	Requestbodies   map[string]interface{} `json:"requestBodies,omitempty"`
	Headers         map[string]interface{} `json:"headers,omitempty"`
	Securityschemes map[string]interface{} `json:"securitySchemes,omitempty"`
	Links           map[string]interface{} `json:"links,omitempty"`
	Callbacks       map[string]interface{} `json:"callbacks,omitempty"`
}

type Schema struct {
	Title                string                 `json:"title,omitempty"`
	Multipleof           float64                `json:"multipleOf,omitempty"`
	Maximum              float64                `json:"maximum,omitempty"`
	Exclusivemaximum     bool                   `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                `json:"minimum,omitempty"`
	Exclusiveminimum     bool                   `json:"exclusiveMinimum,omitempty"`
	Maxlength            int                    `json:"maxLength,omitempty"`
	Minlength            int                    `json:"minLength,omitempty"`
	Pattern              string                 `json:"pattern,omitempty"`
	Maxitems             int                    `json:"maxItems,omitempty"`
	Minitems             int                    `json:"minItems,omitempty"`
	Uniqueitems          bool                   `json:"uniqueItems,omitempty"`
	Maxproperties        int                    `json:"maxProperties,omitempty"`
	Minproperties        int                    `json:"minProperties,omitempty"`
	Required             []string               `json:"required,omitempty"`
	Enum                 []interface{}          `json:"enum,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	Not                  interface{}            `json:"not,omitempty"`
	Allof                []interface{}          `json:"allOf,omitempty"`
	Oneof                []interface{}          `json:"oneOf,omitempty"`
	Anyof                []interface{}          `json:"anyOf,omitempty"`
	Items                interface{}            `json:"items,omitempty"`
	Properties           map[string]interface{} `json:"properties,omitempty"`
	Additionalproperties interface{}            `json:"additionalProperties,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Format               string                 `json:"format,omitempty"`
	Default              interface{}            `json:"default,omitempty"`
	Nullable             bool                   `json:"nullable,omitempty"`
	Discriminator        Discriminator          `json:"discriminator,omitempty"`
	Readonly             bool                   `json:"readOnly,omitempty"`
	Writeonly            bool                   `json:"writeOnly,omitempty"`
	Example              interface{}            `json:"example,omitempty"`
	Externaldocs         Externaldocumentation  `json:"externalDocs,omitempty"`
	Deprecated           bool                   `json:"deprecated,omitempty"`
	Xml                  Xml                    `json:"xml,omitempty"`
}

type Discriminator struct {
	Propertyname string            `json:"propertyName,omitempty"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

type Xml struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

type Response struct {
	Description string                 `json:"description,omitempty"`
	Headers     map[string]interface{} `json:"headers,omitempty"`
	Content     map[string]Mediatype   `json:"content,omitempty"`
	Links       map[string]interface{} `json:"links,omitempty"`
}

type Mediatype struct {
	Schema   interface{}            `json:"schema,omitempty"`
	Example  interface{}            `json:"example,omitempty"`
	Examples map[string]interface{} `json:"examples,omitempty"`
	Encoding map[string]Encoding    `json:"encoding,omitempty"`
}

type Example struct {
	Summary       string      `json:"summary,omitempty"`
	Description   string      `json:"description,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	Externalvalue string      `json:"externalValue,omitempty"`
}

type Header struct {
	Description     string                 `json:"description,omitempty"`
	Required        bool                   `json:"required,omitempty"`
	Deprecated      bool                   `json:"deprecated,omitempty"`
	Allowemptyvalue bool                   `json:"allowEmptyValue,omitempty"`
	Style           string                 `json:"style,omitempty"`
	Explode         bool                   `json:"explode,omitempty"`
	Allowreserved   bool                   `json:"allowReserved,omitempty"`
	Schema          interface{}            `json:"schema,omitempty"`
	Content         map[string]Mediatype   `json:"content,omitempty"`
	Example         interface{}            `json:"example,omitempty"`
	Examples        map[string]interface{} `json:"examples,omitempty"`
}

type Paths struct {
}

type Pathitem struct {
	Ref         *string       `json:"$ref,omitempty"`
	Summary     string        `json:"summary,omitempty"`
	Description string        `json:"description,omitempty"`
	Get         Operation     `json:"get,omitempty"`
	Put         Operation     `json:"put,omitempty"`
	Post        Operation     `json:"post,omitempty"`
	Delete      Operation     `json:"delete,omitempty"`
	Options     Operation     `json:"options,omitempty"`
	Head        Operation     `json:"head,omitempty"`
	Patch       Operation     `json:"patch,omitempty"`
	Trace       Operation     `json:"trace,omitempty"`
	Servers     []Server      `json:"servers,omitempty"`
	Parameters  []interface{} `json:"parameters,omitempty"`
}

type Operation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Externaldocs Externaldocumentation  `json:"externalDocs,omitempty"`
	Operationid  string                 `json:"operationId,omitempty"`
	Parameters   []interface{}          `json:"parameters,omitempty"`
	Requestbody  interface{}            `json:"requestBody,omitempty"`
	Responses    Responses              `json:"responses,omitempty"`
	Callbacks    map[string]interface{} `json:"callbacks,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty"`
	Security     []Securityrequirement  `json:"security,omitempty"`
	Servers      []Server               `json:"servers,omitempty"`
}

type Responses struct {
	Default interface{} `json:"default,omitempty"`
}

type Securityrequirement struct {
}

type Tag struct {
	Name         string                `json:"name,omitempty"`
	Description  string                `json:"description,omitempty"`
	Externaldocs Externaldocumentation `json:"externalDocs,omitempty"`
}

type Externaldocumentation struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type Parameter struct {
	Name            string                 `json:"name,omitempty"`
	In              string                 `json:"in,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Required        bool                   `json:"required,omitempty"`
	Deprecated      bool                   `json:"deprecated,omitempty"`
	Allowemptyvalue bool                   `json:"allowEmptyValue,omitempty"`
	Style           string                 `json:"style,omitempty"`
	Explode         bool                   `json:"explode,omitempty"`
	Allowreserved   bool                   `json:"allowReserved,omitempty"`
	Schema          interface{}            `json:"schema,omitempty"`
	Content         map[string]Mediatype   `json:"content,omitempty"`
	Example         interface{}            `json:"example,omitempty"`
	Examples        map[string]interface{} `json:"examples,omitempty"`
}

type Requestbody struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]Mediatype `json:"content,omitempty"`
	Required    bool                 `json:"required,omitempty"`
}

type Apikeysecurityscheme struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
}

type Httpsecurityscheme struct {
	Scheme       string `json:"scheme,omitempty"`
	Bearerformat string `json:"bearerFormat,omitempty"`
	Description  string `json:"description,omitempty"`
	Type         string `json:"type,omitempty"`
}

type Oauth2securityscheme struct {
	Type        string     `json:"type,omitempty"`
	Flows       Oauthflows `json:"flows,omitempty"`
	Description string     `json:"description,omitempty"`
}

type Openidconnectsecurityscheme struct {
	Type             string `json:"type,omitempty"`
	Openidconnecturl string `json:"openIdConnectUrl,omitempty"`
	Description      string `json:"description,omitempty"`
}

type Oauthflows struct {
	Implicit          Implicitoauthflow          `json:"implicit,omitempty"`
	Password          Passwordoauthflow          `json:"password,omitempty"`
	Clientcredentials Clientcredentialsflow      `json:"clientCredentials,omitempty"`
	Authorizationcode Authorizationcodeoauthflow `json:"authorizationCode,omitempty"`
}

type Implicitoauthflow struct {
	Authorizationurl string            `json:"authorizationUrl,omitempty"`
	Refreshurl       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

type Passwordoauthflow struct {
	Tokenurl   string            `json:"tokenUrl,omitempty"`
	Refreshurl string            `json:"refreshUrl,omitempty"`
	Scopes     map[string]string `json:"scopes,omitempty"`
}

type Clientcredentialsflow struct {
	Tokenurl   string            `json:"tokenUrl,omitempty"`
	Refreshurl string            `json:"refreshUrl,omitempty"`
	Scopes     map[string]string `json:"scopes,omitempty"`
}

type Authorizationcodeoauthflow struct {
	Authorizationurl string            `json:"authorizationUrl,omitempty"`
	Tokenurl         string            `json:"tokenUrl,omitempty"`
	Refreshurl       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

type Link struct {
	Operationid  string                 `json:"operationId,omitempty"`
	Operationref string                 `json:"operationRef,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	Requestbody  interface{}            `json:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Server       Server                 `json:"server,omitempty"`
}

type Callback struct {
}

type Encoding struct {
	Contenttype   string                 `json:"contentType,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty"`
	Style         string                 `json:"style,omitempty"`
	Explode       bool                   `json:"explode,omitempty"`
	Allowreserved bool                   `json:"allowReserved,omitempty"`
}

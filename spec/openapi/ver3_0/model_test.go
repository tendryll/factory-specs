package ver3_0

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func Test_v3_1(t *testing.T) {
	doc, err := LoadFromFile("../../../testing/resources/petstore.yaml")
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Spec.Info.Title
	if title != "Swagger Petstore - OpenAPI 3.0" {
		t.Error("incorrect title: expected 'Swagger Petstore - OpenAPI 3.0', got ", title)
	}
}

// LoadFromFile reads an OpenAPI JSON/YAML file and unmarshals into Document.
func LoadFromFile(path string) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var spec OpenAPI
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("unmarshal spec: %w", err)
	}
	return &Document{Spec: spec}, nil
}

// LoadFromBytes unmarshals a raw JSON document into Document.
func LoadFromBytes(b []byte) (*Document, error) {
	var spec OpenAPI
	if err := json.Unmarshal(b, &spec); err != nil {
		return nil, fmt.Errorf("unmarshal spec: %w", err)
	}
	return &Document{Spec: spec}, nil
}

// Marshal serializes the wrapped spec back to JSON.
func (d *Document) Marshal() ([]byte, error) {
	return json.MarshalIndent(d.Spec, "", "  ")
}

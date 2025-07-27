package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
	yaml "sigs.k8s.io/yaml/goyaml.v2"
)

var DefaultConfig = &Configuration{}

//go:embed config.yaml
var configYAML []byte

func init() {
	err := DefaultConfig.ParseYAML(configYAML)
	if err != nil {
		panic(fmt.Sprintf("failed to parse default config: %v", err))
	}
	err = DefaultConfig.ParseTemplates()
	if err != nil {
		panic(fmt.Sprintf("failed to compile templates: %v", err))
	}
}

type Configuration struct {
	Urls     []TemplateURL `yaml:"urls"`
	compiled map[string]*template.Template
}

type TemplateURL struct {
	Namespace string `yaml:"namespace"`
	URL       string `yaml:"url"`
	Alias     string `yaml:"alias,omitempty"`
}

func (c *Configuration) ParseYAML(yamlData []byte) error {
	err := yaml.Unmarshal(yamlData, c)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}
	return nil
}

func (c *Configuration) Merge(other *Configuration) {
	if other == nil {
		return
	}
	c.Urls = append(c.Urls, other.Urls...)
}

func (c *Configuration) ParseTemplates() error {
	c.compiled = make(map[string]*template.Template)
	for _, url := range c.Urls {
		tmpl, err := template.New(url.Namespace).Funcs(sprig.FuncMap()).Parse(url.URL)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", url.Namespace, err)
		}
		c.compiled[url.Namespace] = tmpl
	}
	return nil
}

func (c *Configuration) Render(purl Purl) (string, error) {
	tmpl, ok := c.compiled[purl.Type]
	if !ok {
		return "", fmt.Errorf("no template found for type %s", purl.Type)
	}
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, map[string]any{
		"Namespace":  purl.Namespace,
		"Name":       purl.Name,
		"Version":    purl.Version,
		"Subpath":    purl.Subpath,
		"Qualifiers": purl.Qualifiers.Map(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute template for %s: %w", purl.Type, err)
	}
	return buf.String(), nil
}

package main

import (
	_ "embed"

	"github.com/package-url/packageurl-go"
)

type Purl struct {
	packageurl.PackageURL
}

func Parse(purl string) (Purl, error) {
	parsed, err := packageurl.FromString(purl)
	if err != nil {
		return Purl{}, err
	}
	return Purl{parsed}, nil
}

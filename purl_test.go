package main

import (
	"testing"
)

func TestURL(t *testing.T) {
	var tests = []struct {
		input string
		url   string
	}{
		// keep-sorted start
		{"pkg:cargo/llm-pricing", "https://crates.io/crates/llm-pricing"},
		{"pkg:cargo/llm-pricing@0.1.0", "https://crates.io/crates/llm-pricing/0.1.0"},
		{"pkg:circleci/slack/slack", "https://circleci.com/developer/orbs/orb/slack/slack"},
		{"pkg:circleci/slack/slack@1", "https://circleci.com/developer/orbs/orb/slack/slack?version=1"},
		{"pkg:circleci/slack/slack@1.0.4", "https://circleci.com/developer/orbs/orb/slack/slack?version=1.0.4"},
		{"pkg:docker/alpine/git", "https://hub.docker.com/r/alpine/git"},
		{"pkg:docker/aquasecurity/trivy?repository_url=public.ecr.aws", "https://gallery.ecr.aws/aquasecurity/trivy"},
		{"pkg:docker/google_appengine/python?repository_url=gcr.io", "https://gcr.io/google_appengine/python"},
		{"pkg:docker/minio/minio?repository_url=quay.io", "https://quay.io/minio/minio"},
		{"pkg:docker/node", "https://hub.docker.com/r/library/node"},
		{"pkg:docker/puppeteer/puppeteer?repository_url=ghcr.io", "https://ghcr.io/puppeteer/puppeteer"},
		{"pkg:gem/rack", "https://rubygems.org/gems/rack"},
		{"pkg:gem/rails@7.1.3", "https://rubygems.org/gems/rails/versions/7.1.3"},
		{"pkg:github/actions/checkout", "https://github.com/actions/checkout"},
		{"pkg:github/actions/checkout#.github/workflows", "https://github.com/actions/checkout/tree/HEAD/.github/workflows"},
		{"pkg:github/actions/checkout@main", "https://github.com/actions/checkout/tree/main"},
		{"pkg:github/actions/checkout@main#.github/workflows", "https://github.com/actions/checkout/tree/main/.github/workflows"},
		{"pkg:githubactions/actions/checkout@v3", "https://github.com/actions/checkout/tree/v3"},
		{"pkg:githubactions/github/codeql-action/analyze@v3", "https://github.com/github/codeql-action/tree/v3/analyze"},
		{"pkg:githubactions/github/codeql-action@v3#analyze", "https://github.com/github/codeql-action/tree/v3/analyze"},
		{"pkg:golang/k8s.io/apimachinery/pkg/util/sets", "https://pkg.go.dev/k8s.io/apimachinery/pkg/util/sets"},
		{"pkg:npm/express", "https://www.npmjs.com/package/express"},
		{"pkg:npm/express@4.18.2", "https://www.npmjs.com/package/express/v/4.18.2"},
		{"pkg:nuget/CycloneDX", "https://www.nuget.org/packages/CycloneDX"},
		{"pkg:nuget/CycloneDX@1.3.0", "https://www.nuget.org/packages/CycloneDX/1.3.0"},
		{"pkg:pypi/uv", "https://pypi.org/project/uv"},
		{"pkg:pypi/uv@0.8.3", "https://pypi.org/project/uv/0.8.3"},
		// keep-sorted end
	}

	for _, test := range tests {
		p, err := Parse(test.input)
		if err != nil {
			t.Errorf("Parse(%q) = %q", test.input, err)
		}
		url, err := DefaultConfig.Render(p)
		if err != nil {
			t.Errorf("Render(%q) = %q", p, err)
		}
		if url != test.url {
			t.Errorf("Parse(%q).URL() = %q, want %q", test.input, url, test.url)
		}
	}
}

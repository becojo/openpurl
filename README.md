# openpurl

openpurl is a CLI tool to open or print canonical URLs for software packages using the [package-url](https://github.com/package-url/purl-spec) specification.

``` sh
openpurl pkg:pypi/uv@0.8.3
# Opens the URL https://pypi.org/project/uv/0.8.3
```

## Usage

```
usage: openpurl [--print|-p] [pkg:...]
```

## Building
Clone the repository and build:

```sh
git clone https://github.com/becojo/openpurl.git && cd openpurl
go build -o openpurl
```

URL templates for each package ecosystem are defined in `config.yaml` using Go templates.


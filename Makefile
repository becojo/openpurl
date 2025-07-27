GO := $(shell which go)
SRC := $(shell find . -name '*.go')
TARGET := openpurl

openpurl: $(SRC)
	CGO_ENABLED=0 $(GO) build -o $(TARGET)

clean:
	rm -f $(TARGET)

test:
	go test ./...

format: format.go format.keep-sorted

format.go:
	$(GO) fmt ./...

format.keep-sorted:
	keep-sorted $(shell grep -rl --exclude=openpurl 'keep-sorted'' start' .)

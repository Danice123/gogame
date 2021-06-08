.DEFAULT_GOAL := default
.PHONY: deps tidy fmt vet build dist
default: deps tidy fmt vet build dist

deps:
	go get -u github.com/go-bindata/go-bindata/...

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	mkdir -p dist
	go build -o dist/ ./...

dist:
	cp -R pkg/sheets dist
	cp -R pkg/maps dist
	zip -r dist.zip dist

clean:
	rm -rf dist


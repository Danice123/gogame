.DEFAULT_GOAL := default
.PHONY: deps tidy fmt vet test bindata build
default: deps tidy fmt vet test bindata build dist

deps:
	go get -u github.com/go-bindata/go-bindata/...
	asdf reshim golang

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... | grep -v 'no test files' || :

pkg/data/bindata.go:
	@mkdir -p pkg/data
	go-bindata -o pkg/data/bindata.go -pkg data data/...

bindata: pkg/data/bindata.go

build: bindata
	@mkdir -p dist
	go build -o dist/ ./...

dist/game: build

dist:
	@#cp -R pkg/sheets dist
	@#cp -R pkg/maps dist
	zip -r dist.zip dist

run: dist/game
	./dist/game

clean:
	rm -rf pkg/data
	rm -rf dist


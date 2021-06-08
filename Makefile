.DEFAULT_GOAL := default
.PHONY: deps tidy fmt vet test build
default: deps tidy fmt vet test build dist

deps:
	go get -u github.com/go-bindata/go-bindata/...

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... | grep -v 'no test files' || :

build:
	mkdir -p dist
	go build -o dist/ ./...

dist/game: build

dist:
	cp -R pkg/sheets dist
	cp -R pkg/maps dist
	zip -r dist.zip dist

run: dist/game
	(cd dist && ./game) # Need to bindata the files to remove cd

clean:
	rm -rf dist


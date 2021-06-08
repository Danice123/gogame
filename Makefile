.DEFAULT_GOAL := default
.PHONY: clean deps tidy fmt vet test build
default: clean deps tidy fmt vet test build dist

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
	go-bindata -nomemcopy -prefix data/ -o pkg/data/bindata.go -pkg data data/...

build:
	@rm pkg/data/bindata.go && $(MAKE) pkg/data/bindata.go
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
	rm -rf dist dist.zip


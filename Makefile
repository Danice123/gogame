.DEFAULT_GOAL := default

default: build dist

.PHONY: build
build:
	mkdir -p dist
	CGO_ENABLED=0 go build -o dist/gogame main.go

.PHONY: dist
dist:
	cp -R sheets dist
	cp -R maps dist
	zip -r dist.zip dist




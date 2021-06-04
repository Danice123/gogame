PHONY: build


build:
	mkdir -p dist
	go build -o dist/gogame main.go
	cp -R sheets dist
	cp -R maps dist
	zip -r dist.zip dist

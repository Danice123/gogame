PHONY: build


build:
	mkdir -p dist
	go build -o dist/gogame main.go
	cp -R sheets dist
	cp -R maps dist
	zip -r dist.zip dist

sprites:
	TexturePacker resources/sheets/mettaur.tps
	jq -s '.[0] * .[1]' resources/sheets/mettaur.json resources/sheets/mettaur.ani.json > resources/sheets/mettaur.json.new
	mv resources/sheets/mettaur.json.new resources/sheets/mettaur.json

	TexturePacker resources/sheets/generic.tps
	jq -s '.[0] * .[1]' resources/sheets/generic.json resources/sheets/generic.ani.json > resources/sheets/generic.json.new
	mv resources/sheets/generic.json.new resources/sheets/generic.json
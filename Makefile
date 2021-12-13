PHONY: build


build:
	mkdir -p dist
	go build -o dist/gogame main.go
	cp -R sheets dist
	cp -R maps dist
	zip -r dist.zip dist

sprites:
	TexturePacker resources/sheets/rockman.tps
	jq -s '.[0] * .[1]' resources/sheets/rockman.json resources/sheets/rockman.ani.json > resources/sheets/rockman.json.new
	mv resources/sheets/rockman.json.new resources/sheets/rockman.json
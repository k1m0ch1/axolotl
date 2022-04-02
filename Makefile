all:
	go fmt ./...
	go clean -testcache
	go test -v ./...

build:
	GOOS=windows GOARCH=amd64 go build -o .bin/axolotl-amd64.exe .
	GOOS=windows GOARCH=386 go build -o .bin/axolotl-386.exe .
	GOOS=darwin GOARCH=amd64 go build -o .bin/axolotl-amd64-darwin .
	GOOS=linux GOARCH=amd64 go build -o .bin/axolotl-amd64-linux .
	GOOS=linux GOARCH=386 go build -o .bin/axolotl-386-linux .

dev-install:
	sudo cp .bin/axolotl-amd64-linux /usr/bin/axolotl

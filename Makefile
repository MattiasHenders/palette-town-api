dep:
	go mod download
	go mod tidy

### Server ###
build-server: dep
	rm -rf bin/
	go build -o bin/ src/main.go

run-server: build-server
	./bin/server

build-server-windows: dep
	- rm -r .\bin
	- mkdir bin
	go build -o bin/server.exe src/main.go

run-server-windows: build-server-windows
	./bin/server.exe
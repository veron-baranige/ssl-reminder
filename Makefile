BINARY_NAME=ssl-reminder

build: clean
	mkdir dist
	CGO_ENABLED=0 go build -o dist/${BINARY_NAME}
	chmod +x dist/${BINARY_NAME}

clean: 
	@if [ -d dist ]; then rm -rf dist; fi
	go clean
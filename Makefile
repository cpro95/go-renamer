BINARY_NAME=go-renamer

run:
	go run *.go
build:
	go build -o $(BINARY_NAME) *.go
fmt:
	go fmt ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go get github.com/gdamore/tcell
	go get github.com/gdamore/tcell/encoding
	go get github.com/sirupsen/logrus


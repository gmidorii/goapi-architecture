BINARY='gapi'
PORT='3333'

build:
	go build -o $(BINARY)

run:
	go run *.go

gin:
	gin -a $(PORT)


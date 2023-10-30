BINARY_NAME=lazlo
CMD_PATH=./main.go

all: build

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

clean:
	rm -f bin/$(BINARY_NAME)
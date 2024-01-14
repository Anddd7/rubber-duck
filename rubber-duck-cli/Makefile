BIN		= ./bin
NAME	= rbcli

run:
	go run ./

build:
	go build -o $(BIN)/$(NAME) ./

clean:
	rm -rf $(BIN)

fmt:
	go fmt ./...
	gofumpt -l -w .
	go vet ./...

test:
	go test

lint:
	golangci-lint run -v

cover:
	go test -coverprofile coverage.out

coverweb: cover
	go tool cover -html=coverage.out

check: fmt lint cover

dep:
	go install mvdan.cc/gofumpt@latest
	go mod download

install: build
	mkdir -p ~/bin
	mv $(BIN)/$(NAME) ~/bin/$(NAME) 

uninstall:
	rm ~/bin/$(NAME)
VERSION := 0.0.1

dependency:
	apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	protoc \
		--go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		service.proto

docker:
	docker build -t docker.io/anddd7/grpcbin:$(VERSION) .
	docker push docker.io/anddd7/grpcbin:$(VERSION)

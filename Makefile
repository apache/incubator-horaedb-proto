.PHONY: all go rust java

all: go rust java

dependence:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go: dependence
	./generate-go.sh

rust:
	cargo build

java:
	mvn -f java/pom.xml clean install

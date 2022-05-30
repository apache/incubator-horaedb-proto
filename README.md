# ceresdbproto

Protocol buffer files for CeresDB

# Rust

## Usage

Add this to your Cargo.toml:

```
ceresdbproto = { git = "https://github.com/CeresDB/ceresdbproto.git"}
```

# Golang
## Prepare

- https://grpc.io/docs/languages/go/quickstart/
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## Generate
```
sh gen-go.sh
```

## Usage

```
go get github.com/CeresDB/ceresdbproto/go/ceresdbproto
```

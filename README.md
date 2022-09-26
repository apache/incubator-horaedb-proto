# ceresdbproto

Protocol buffer files for CeresDB.

# Usage
## Rust

Add this to your Cargo.toml:

```
ceresdbproto = { git = "https://github.com/CeresDB/ceresdbproto.git"}
```

## Go

```
go get github.com/CeresDB/ceresdbproto
```

# Generate code

Install [Protocol Buffers 3.20.1](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.1)+ compiler.

## Rust

Nothing required

## Go

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
sh gen-go.sh
```

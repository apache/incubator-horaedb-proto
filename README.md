# ceresdbproto

Protocol buffer files for CeresDB.

# User Guides

## Rust

Add the following dependency to your Cargo.toml, and specify a precise commit hash:

```
ceresdbproto = { git = "https://github.com/CeresDB/ceresdbproto.git"}
```

## Go

```sh
go get github.com/CeresDB/ceresdbproto/golang
```

## Java

Add a maven dependency to your project.

```xml
<dependency>
    <groupId>io.ceresdb</groupId>
    <artifactId>ceresdb-proto-internal</artifactId>
    <version>0.1.0</version>
</dependency>
```

# Developer Guides
After modifying the proto files, something else for different programming languages should be done.

## Java
1. Name the maven project to a new version.
2. Rebuild the maven project lies in the `java` directory (During the build process, java code for the latest proto will be generated).
2. Publish the build result to the central maven repository.

## Rust
As for the Rust projects that depends on this project, everything will be generated during the build process of themselves. So nothing else needs to be done for `Rust`.

## Golang
1. Install the prerequisites:
* Install [Protocol Buffers (above 3.20.1)](https://github.com/protocolbuffers/protobuf/releases) compiler.
* Install golang specific protocal and grpc toolsets.
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
2. execute the `gen-go.sh` in the `golang` directory.

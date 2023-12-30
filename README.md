# Protocol of [Apache HoraeDB™](https://github.com/apache/incubator-horaedb)

[![Crates.io](https://img.shields.io/crates/v/horaedb-proto.svg)](https://crates.io/crates/horaedb-proto)
[![Go Reference](https://pkg.go.dev/badge/github.com/apache/incubator-horaedb-proto.svg)](https://pkg.go.dev/github.com/apache/incubator-horaedb-proto)

Protocol buffer files for HoraeDB. Projects that manage generated code of different languages are also in this repository. They acts as underlying dependency of client, server and meta.

## User Guide

### Rust

```sh
cargo add horaedbproto
```

### Golang

```sh
go get github.com/apache/incubator-horaedb-proto/golang
```

### Java

Add a maven dependency to your project.

```xml
<dependency>
    <groupId>org.apache.horaedb</groupId>
    <artifactId>horaedb-proto-internal</artifactId>
    <version>${horaedb-proto-internal.version}</version>
</dependency>
```

## Developer Guide

After modifying the proto files, something else for different programming languages should be done.

### Java

1. Name the maven project to a new version.
2. Rebuild the maven project lies in the `java` directory (During the build process, java code for the latest proto will be generated).
2. Publish the build result to the central maven repository.

### Rust

As for the Rust projects that depends on this project, everything will be generated during the build process of themselves. So nothing else needs to be done for `Rust`.

### Golang

1. Install [Protocol Buffers v3.20.1](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.1) compiler.
2. Execute `make go`.

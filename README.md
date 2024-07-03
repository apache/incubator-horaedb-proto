# Apache HoraeDB (Incubating) Protocol Buffer Files

[![Crates.io](https://img.shields.io/crates/v/horaedbproto.svg)](https://crates.io/crates/horaedbproto)
[![Go Reference](https://pkg.go.dev/badge/github.com/apache/incubator-horaedb-proto.svg)](https://pkg.go.dev/github.com/apache/incubator-horaedb-proto)
[![Maven Central](https://img.shields.io/maven-central/v/org.apache.horaedb/horaedb-proto-internal.svg?logo=Apache+Maven&logoColor=blue)](https://central.sonatype.com/artifact/org.apache.horaedb/horaedb-proto-internal)

Protocol buffer files for [Apache HoraeDB (Incubating)](https://github.com/apache/incubator-horaedb). Projects that manage generated code of different languages are also in this repository. They act as underlying dependencies of the client, server, and meta.

> [!IMPORTANT]
> Apache HoraeDB (incubating) is an effort undergoing incubation at the Apache
> Software Foundation (ASF), sponsored by the Apache Incubator PMC.
>
> Please read the [DISCLAIMER](DISCLAIMER) and a full explanation of ["incubating"](https://incubator.apache.org/policy/incubation.html).

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
2. Rebuild the maven project in the `java` directory (During the build process, java code for the latest proto will be generated).
2. Publish the build result to the central maven repository.

### Rust

As for the Rust projects that depend on this project, everything will be generated during the build process themselves. So nothing else needs to be done for `Rust`.

### Golang

1. Install [Protocol Buffers v25.1](https://github.com/protocolbuffers/protobuf/releases/tag/v25.1) compiler.
2. Execute `make go`.

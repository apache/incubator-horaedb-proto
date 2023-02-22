// Copyright 2022 CeresDB Project Authors. Licensed under Apache-2.0.

//! Protobuf messages

pub use protos::*;

#[allow(dead_code)]
#[allow(clippy::all)]
mod protos {
    include!(concat!(env!("OUT_DIR"), "/mod.rs"));
}

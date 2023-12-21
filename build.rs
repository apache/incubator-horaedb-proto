// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

use std::{io::Write, path::Path};
use walkdir::WalkDir;

struct Builder {
    out_dir: String,
    proto_dir: String,
    proto_paths: Vec<String>,
    module_names: Vec<String>,
}

impl Builder {
    fn new(out_dir: String, proto_dir: String) -> Result<Self, Box<dyn std::error::Error>> {
        // Get each proto's path and related module name.
        let proto_dir = Path::new(proto_dir.as_str()).to_str().unwrap().to_owned();

        let mut proto_paths = Vec::new();
        let mut module_names = Vec::new();
        for entry in WalkDir::new(proto_dir.clone())
            .into_iter()
            .filter_map(Result::ok)
            .filter(|e| !e.file_type().is_dir() && !e.file_type().is_symlink())
        {
            proto_paths.push(entry.path().display().to_string());

            let module_name = entry
                .file_name()
                .to_str()
                .unwrap()
                .split(".")
                .into_iter()
                .next()
                .unwrap()
                .to_owned();
            module_names.push(module_name);
        }

        Ok(Self {
            out_dir,
            proto_dir,
            proto_paths,
            module_names,
        })
    }

    fn generate(self) -> Result<(), Box<dyn std::error::Error>> {
        self.generate_files()?.generate_mod_file()?;

        Ok(())
    }

    fn generate_files(self) -> Result<Self, Box<dyn std::error::Error>> {
        // Compile proto files.
        tonic_build::configure()
            .compile(&self.proto_paths, &[self.proto_dir.clone()])
            .map_err(|e| Box::new(e))?;

        Ok(self)
    }

    fn generate_mod_file(self) -> Result<Self, Box<dyn std::error::Error>> {
        // Generate mod.rs file.
        let mut mod_file = std::fs::File::create(format!("{}/mod.rs", self.out_dir))?;
        for module_name in self.module_names.iter() {
            write!(mod_file, "pub mod {};\n", module_name)?;
        }

        Ok(self)
    }
}

const ENABLE_VENDOR_ENV: &str = "HORAEDBPROTO_ENABLE_VENDORED";

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Download the protoc and set the path for tonic_build.

    println!("cargo:rerun-if-env-changed={}", ENABLE_VENDOR_ENV);

    let enable_vendor = std::env::var(ENABLE_VENDOR_ENV).unwrap_or("true".to_string());
    if "true" == enable_vendor {
        let protoc_path = protoc_bin_vendored::protoc_bin_path().map_err(|e| Box::new(e))?;
        std::env::set_var("PROTOC", protoc_path.as_os_str());
    }

    // Build protos.
    let out_dir = std::env::var("OUT_DIR").map_err(|e| Box::new(e))?;

    let protos_dir = "protos".to_string();
    let builder = Builder::new(out_dir, protos_dir)?;
    builder.generate()
}

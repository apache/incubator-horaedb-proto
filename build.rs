
fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure().out_dir("./src").compile(
        &[
            "protos/cluster.proto",
            "protos/common.proto",
            "protos/meta_event.proto",
            "protos/meta_service.proto",
            "protos/meta_storage.proto",
            "protos/prometheus.proto",
            "protos/storage.proto",
        ],
        &["protos"]
    )?;
    Ok(())
}

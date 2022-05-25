use protobuf_builder::Builder;

fn generate_pb() {
    Builder::new().search_dir_for_protos("protos").generate();
}

fn main() {
    generate_pb();
}

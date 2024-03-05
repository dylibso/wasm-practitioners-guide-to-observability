mod bindings;
use crate::bindings::dylibso::observe::api::*;

fn main() {
    span_enter("hello world");
    log(LogLevel::Info, b"hello world");
    span_exit();
    println!("Hello, world!");
}

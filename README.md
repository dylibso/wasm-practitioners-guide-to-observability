# Wasm Practitioners Guide to Observability

In this workshop, we’ll be covering the full spectrum of observability in Wasm - and what the options are to instrument your code, as well as how to ensure Wasm telemetry makes it out of the runtime and seamlessly blends into your existing observability stack.

We’ll discuss how WebAssembly is well suited for all kinds of automatic instrumentation that can be done after code is compiled, and how to precisely instrument code where needed.

The technology and libraries used is planned to undergo specification as “wasi-observe” (final name TBD) by the WASI subgroup, and is currently available as open source libraries from Dylibso, as the Observe SDK: https://github.com/dylibso/observe-sdk

Attendees will write some code to compile to Wasm, link import functions to handle telemetry integrated into their host applications, and then see live telemetry data visualized in their APM.


# Setup

## 1. Intstall Jaeger (viewing OpenTelemetry data)

```sh
docker run -d --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.54
```

Run this at http://localhost:16686/ in your browser.


## 2. Run the Go host service (executes Wasm code) 

```sh
cd src/host
docker run --network host workshop-host
# be sure to pass `--network host` here so that the Go application can reach Jaeger
```

This will start a server running at http://localhost:3000

## Upload instrumented wasm code to the Go host service

```sh
# from the root of the repo

# upload the Rust-based wasm module, name it `rust-manual`
curl -F wasm=@src/guest/rust/rust.wasm "http://localhost:3000/upload?name=rust-manual"

# upload the Go-based wasm module, name it `go-manual`
curl -F wasm=@src/guest/go/main.wasm "http://localhost:3000/upload?name=go-manual"
```

## Run instrumented wasm code on the Go host service

```sh
# run the Rust-based wasm module
curl -X POST "http://localhost:3000/run?name=rust-manual"

# run the Go-based wasm module
curl -X POST "http://localhost:3000/run?name=go-manual"
```

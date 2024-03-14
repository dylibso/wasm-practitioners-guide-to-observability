# Wasm Practitioners Guide to Observability

# Setup

## 1. Install Jaeger (viewing OpenTelemetry data)

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

docker build --tag workshop-host .
docker run --network host workshop-host
# be sure to pass `--network host` here so that the Go application can reach Jaeger
```

This will start a server running at http://localhost:3000

## 3. Upload instrumented wasm code to the Go host service

```sh
# from the root of the repo

# upload the Rust-based wasm module, name it `rust-manual`
curl -F wasm=@src/guest/modules/manual/rust/rust.wasm "http://localhost:3000/upload?name=rust-manual"

# upload the Go-based wasm module, name it `go-manual`
curl -F wasm=@src/guest/modules/manual/go/main.wasm "http://localhost:3000/upload?name=go-manual"
```

## 4. Run instrumented wasm code on the Go host service

```sh
# run the Rust-based wasm module
curl -X POST "http://localhost:3000/run?name=rust-manual"

# run the Go-based wasm module
curl -X POST "http://localhost:3000/run?name=go-manual"
```

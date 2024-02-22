# Wasm Practitioners Guide to Observability
Steve Manuel - Dylibso for WASMIO 2024

In this workshop, we’ll be covering the full spectrum of observability in Wasm - and what the options are to instrument your code, as well as how to ensure Wasm telemetry makes it out of the runtime and seamlessly blends into your existing observability stack.

We’ll discuss how WebAssembly is well suited for all kinds of automatic instrumentation that can be done after code is compiled, and how to precisely instrument code where needed.

The technology and libraries used is planned to undergo specification as “wasi-observe” (final name TBD) by the WASI subgroup, and is currently available as open source libraries from Dylibso, as the Observe SDK: https://github.com/dylibso/observe-sdk

Attendees will write some code to compile to Wasm, link import functions to handle telemetry integrated into their host applications, and then see live telemetry data visualized in their APM.

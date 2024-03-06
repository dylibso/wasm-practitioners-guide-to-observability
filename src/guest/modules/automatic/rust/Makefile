.PHONY: rundebug
rundebug: rust.debug.instr.wasm
	curl -F wasm=@rust.debug.instr.wasm "http://localhost:3000/upload?name=rust-auto-debug"
	curl -X POST "http://localhost:3000/run?name=rust-auto-debug"

.PHONY: rundebugconfig
rundebugconfig: rust.debug.config.instr.wasm
	curl -F wasm=@rust.debug.config.instr.wasm "http://localhost:3000/upload?name=rust-auto-debug-config"
	curl -X POST "http://localhost:3000/run?name=rust-auto-debug-config"

.PHONY: runrelease
runrelease: rust.instr.wasm
	curl -F wasm=@rust.instr.wasm "http://localhost:3000/upload?name=rust-auto"
	curl -X POST "http://localhost:3000/run?name=rust-auto"

rust.debug.wasm: src/main.rs
	cargo build --target=wasm32-wasi
	cp target/wasm32-wasi/debug/rust.wasm rust.debug.wasm

rust.debug.instr.wasm: rust.debug.wasm
	curl --fail -F wasm=@rust.debug.wasm \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> rust.debug.instr.wasm

rust.debug.config.instr.wasm: rust.debug.wasm config.json
	curl --fail -F wasm=@rust.debug.wasm \
	-F config=@config.json \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> rust.debug.config.instr.wasm

rust.wasm: src/main.rs
	cargo build --target=wasm32-wasi --release
	cp target/wasm32-wasi/release/rust.wasm rust.wasm

rust.instr.wasm: rust.wasm
	curl --fail -F wasm=@rust.wasm \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> rust.instr.wasm

.PHONY: rundebug
rundebug: main.debug.instr.wasm
	curl -F wasm=@main.debug.instr.wasm "http://localhost:3000/upload?name=go-auto-debug"
	curl -X POST "http://localhost:3000/run?name=go-auto-debug"

.PHONY: rundebugconfig
rundebugconfig: main.debug.config.instr.wasm
	curl -F wasm=@main.debug.config.instr.wasm "http://localhost:3000/upload?name=go-auto-debug-config"
	curl -X POST "http://localhost:3000/run?name=go-auto-debug-config"

.PHONY: runrelease
runrelease: main.instr.wasm
	curl -F wasm=@main.instr.wasm "http://localhost:3000/upload?name=go-auto"
	curl -X POST "http://localhost:3000/run?name=go-auto"

main.debug.wasm: main.go
	tinygo build -o main.debug.wasm -target wasi -opt=0 main.go

main.debug.instr.wasm: main.debug.wasm
	curl --fail -F wasm=@main.debug.wasm \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> main.debug.instr.wasm

main.debug.config.instr.wasm: main.debug.wasm config.json
	curl --fail -F wasm=@main.debug.wasm \
	-F config=@config.json \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> main.debug.config.instr.wasm

main.wasm: main.go
	tinygo build -o main.wasm -target wasi main.go

main.instr.wasm: main.wasm
	curl --fail -F wasm=@main.wasm \
	-H "Authorization: Bearer $(API_KEY)" \
	https://compiler-preview.dylibso.com/instrument \
	> main.instr.wasm

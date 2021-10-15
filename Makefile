build:
	tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go
run:
	wasmedge --dir .:.  main.wasm
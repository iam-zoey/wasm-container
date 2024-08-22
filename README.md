# wasm-container
This project demonstrates the use of WebAssembly (WASM) within Docker containers. 

The application code, `main.go`, is a simple JSON formatter that reads JSON data from the `data.json` file. Note that the environment variable `JSON_FILE_PATH` must be passed to `main.go` to specify the data file path.

## Prerequistie 
- Docker (version 4.15+ with containerd image store enabled)
- Go (version 1.21+)
- Wasmtime (a WebAssembly runtime)


#  Building the WASM module 
```shell 
# Compile main.go into main.wasm 
GOOS=wasip1 GOARCH=wasm go build -o main.wasm  main.go

# [opt] Compiile into .wat
GOOS=wasip1 GOARCH=wasm go build -o main.wat main.go
``` 

# Running WASM Module 
1. Locally 
```shell 
wasmtime run --mapdir /input::. --env JSON_FILE_PATH=/input/data.json main.wasm
```
2. Using Docker 
```shell 
# Build the Docker iamge 
docker build . -t <IMAGE-NAME>

# Run the container 
docker run --rm --runtime=io.containerd.wasmtime.v1 <IMAGE-NAME>
```

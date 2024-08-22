#Dockerfile 

# Use empty image 
FROM scratch

# Copy the WebAssembly binary and the data file into the container
COPY main.wasm /main.wasm 
COPY data.json /data.json

# Set environment variable for the JSON file path
ENV JSON_FILE_PATH=/data.json

# Command to run the WebAssembly module using Wasmtime
ENTRYPOINT [ "/main.wasm" ]
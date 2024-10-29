# gRPC Example in Go

This project demonstrates setting up a simple gRPC server and client in Go, using `buf` for managing Protocol Buffers.

## Setup Guide

### 1. Initialize Go Module

Run the following command to initialize your Go module:

```bash
go mod init github.com/yourusername/grpc-example
```

### 2. Download Protocol Buffer Compiler

Install `protoc`, the Protocol Buffer compiler. Follow the [official installation instructions](https://grpc.io/docs/protoc-installation/) for your operating system.

### 3. Install `buf`

Install `buf`, a tool for managing Protocol Buffers:

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
```

### 4. Create `buf.yaml`

Create a file named `buf.yaml` at the root of your project with the following content:

```yaml
version: v1
```

### 5. Create `buf.gen.yaml`

Create a file named `buf.gen.yaml` with the following configuration for Go and gRPC code generation:

```yaml
version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt: paths=source_relative
```

### 6. Add Dependencies

Install the necessary Go dependencies for gRPC and Protocol Buffers:

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

## Running the Project

1. Define your `.proto` file in the `proto` directory, specifying the gRPC services and messages.
2. Use `buf generate` to generate the Go code from your `.proto` file:
   ```bash
   buf generate
   ```
3. Implement the gRPC server and client according to the generated code.
4. Run the server and client to test the setup.

## Additional Resources

- [gRPC Documentation](https://grpc.io/docs/)
- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers)
- [`buf` Documentation](https://docs.buf.build/)

---

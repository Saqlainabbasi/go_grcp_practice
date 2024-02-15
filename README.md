# go_grcp_practice

This is a simple Project to practice gRPC with Go.Includes a simple server and client to demonstrate the basic functionality of gRPC.Like

- Unary RPC
- Server Streaming
- Client Streaming
- Bi-Directional Streaming

## useful commands

- To generate the gRPC code from the proto file

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto
```

- To set the path

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

- To run the server||client

```bash
go run *.go
```

## References

- [gRPC](https://grpc.io/)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)

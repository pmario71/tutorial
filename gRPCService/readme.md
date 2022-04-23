# Readme

## Recompiling protobufs

```bash
cd /workspaces/tutorial/gRPCService

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  protos/greeter.proto
```

## Buidling & Starting the Server

```bash
go build .
./go-grpcServer
```

## Testing the Server

### Using BloomRPC

* use url: `0.0.0.0:50051`

### Using grpcui

* service implements reflection

```ps1
grpcui -plaintext localhost:50051
```

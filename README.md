# golang-hexagonal-architecture
This is a sample code in golang to represent hexagonal architechture

**Generate the proto files for messages and service**
`protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/number_message.proto`
`protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto`

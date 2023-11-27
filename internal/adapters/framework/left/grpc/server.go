package grpc

import (
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/adapters/framework/left/grpc/pb"
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api}
}

// Run registers the ArithmeticServiceServer to a grpcServer and serves on
// the specified port
func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}

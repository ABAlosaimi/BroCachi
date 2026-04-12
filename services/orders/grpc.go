package orders

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register my services here in the grpcServer

	log.Printf("gRPC server has started on %s", lis.Addr())

	return grpcServer.Serve(lis)
}

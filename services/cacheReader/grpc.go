package cachereader

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

type CacheReaderServer struct{
	addr string 
}

func NewCacheReaderServer(addr string) *CacheReaderServer {
	return &CacheReaderServer{
		addr: addr,
	} 
}

func (s *CacheReaderServer) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// adding my services here 

	log.Printf("gRPC server has started on %s", lis.Addr())

	grpcServer := grpc.NewServer()
	return grpcServer.Serve(lis) 
}
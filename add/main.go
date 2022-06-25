package main

import (
	proto "github.com/CodeDSA/hello/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedGCDServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Compute(ctx context.Context, r *proto.GCDRequest) (*proto.GCDResponse, error) {
	a, b := r.A, r.B

	a = a + b

	return &proto.GCDResponse{Result: a}, nil
}

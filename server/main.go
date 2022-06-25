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
	proto.UnimplementedComputeServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterComputeServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) ComputeCode(ctx context.Context, r *proto.CodeRequest) (*proto.CodeResponse, error) {

	a := r.Problem + " " + r.Code

	return &proto.CodeResponse{Result: a}, nil
}

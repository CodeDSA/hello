package main

import (
	"log"
	"net"
	"fmt"
	"github.com/CodeDSA/hello/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *server) Compute(cxt context.Context, r *pb.GCDRequest) (*pb.GCDResponse , error){
	result := &pb.GCDResponse{}
	result.Result = r.A + r.B
	logMessage := fmt.Sprintf("A: %d    B: %d    sum: %d", r.A, r.B, result.Result)
	log.Println(logMessage)

	return result, nil
}

type server struct{}

func main(){
	lis, err := net.Listen("tcp", ":3000")

	if err != nil{
	log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s , &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
	log.Fatalf("Failed to serve: %v", err)
	}
}


package main

import (
	"context"
	"log"
	"net"

	pb "webconsole-proto/grpc/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConfigServiceServer
}

func (s *server) GetAccessAndMobilityConfig(ctx context.Context, req *pb.EmptyRequest) (*pb.AccessAndMobilityConfigResponse, error) {
	log.Println("Received a GetAccessAndMobilityConfig request")
	return &pb.AccessAndMobilityConfigResponse{
		Configs: []*pb.AccessAndMobilityConfig{
			{
				PlmnId: &pb.PlmnId{Mcc: "210", Mnc: "111"},
				Snssai: &pb.Snssai{Sst: "1", Sd: "11111"},
				Tacs:   []int32{1, 44, 5},
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterConfigServiceServer(grpcServer, &server{})

	log.Println("Server is running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

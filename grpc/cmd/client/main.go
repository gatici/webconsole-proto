package main

import (
	"context"
	"log"
	"time"

	pb "webconsole-proto/grpc/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewConfigServiceClient(conn)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		log.Println("Making a GetAccessAndMobilityConfig request")
		resp, err := client.GetAccessAndMobilityConfig(ctx, &pb.EmptyRequest{})
		if err != nil {
			log.Printf("could not get config: %v", err)
			continue
		}

		log.Println("Received config:")
		for _, cfg := range resp.Configs {
			log.Printf("PLMN: MCC %s MNC %s, SNSSAI: SST %s SD %s, TACs: %v",
				cfg.PlmnId.Mcc, cfg.PlmnId.Mnc,
				cfg.Snssai.Sst, cfg.Snssai.Sd,
				cfg.Tacs,
			)
		}
	}
}

package main

import (
	"context"
	api "inventory/internal/api/inventory/v1"
	"inventory/internal/config"
	repoPart "inventory/internal/repository/part"
	svcPart "inventory/internal/service/part"
	"log"
	"net"
	"os"
	"os/signal"
	desc "shared/pkg/proto/inventory/v1"
	"syscall"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	repo := repoPart.NewRepository(ctx, client, cfg.MongoDBName)
	svc := svcPart.NewService(repo)
	impl := api.NewImplementation(svc)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterInventoryServiceServer(s, impl)
	reflection.Register(s)

	log.Printf("Inventory Service запущен на %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Остановка Inventory Service...")
	s.GracefulStop()
}

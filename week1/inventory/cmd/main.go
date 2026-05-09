package main

import (
	api "inventory/internal/api/inventory/v1"
	repoPart "inventory/internal/repository/part"
	svcPart "inventory/internal/service/part"
	"log"
	"net"
	"os"
	"os/signal"
	desc "shared/pkg/proto/inventory/v1"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	repo := repoPart.NewRepository()
	svc := svcPart.NewService(repo)
	impl := api.NewImplementation(svc)

	lis, err := net.Listen("tcp", ":50051")
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

package main

import (
	"log"
	"net"
	paymentv1 "shared/pkg/proto/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderApi "order/internal/api/order/v1"
	clientGrpc "order/internal/client/grpc"
	inventoryClient "order/internal/client/grpc/inventory/v1"
	paymentClient "order/internal/client/grpc/payment/v1"
	repo "order/internal/repository/order"
	orderService "order/internal/service/order"
	inventoryv1 "shared/pkg/proto/inventory/v1"
	desc "shared/pkg/proto/order/v1"
)

const (
	grpcPort         = ":50051"          // Порт, который будет слушать наш Order Service
	inventoryAddress = "localhost:50052" // Адрес запущенного сервиса Inventory
	paymentAddress   = "localhost:50053" // Адрес запущенного сервиса Payment
)

func main() {

	invConn, err := grpc.NewClient(inventoryAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к Inventory: %v", err)
	}

	payConn, err := grpc.NewClient(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к Payment: %v", err)
	}

	grpcEngine := inventoryv1.NewInventoryServiceClient(invConn)
	payEngine := paymentv1.NewPaymentServiceClient(payConn)

	invCl := inventoryClient.NewClient(grpcEngine)
	payCl := paymentClient.NewClient(payEngine)

	allClients := clientGrpc.NewClients(invCl, payCl)

	repo := repo.NewRepository()

	service := orderService.NewService(repo, allClients)

	apiImpl := orderApi.NewImplementation(service)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("ошибка при прослушивании порта: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOrderServiceServer(s, apiImpl)

	log.Printf("🚀 Order Service запущен на %s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("ошибка при запуске сервера: %v", err)
	}
}

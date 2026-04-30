package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	desc "shared/pkg/proto/inventory/v1"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedInventoryServiceServer
}

func (s *server) GetPart(ctx context.Context, req *desc.GetPartRequest) (*desc.GetPartResponse, error) {
	log.Printf("Запрос детали с UUID: %s", req.GetUuid())

	part := mockPart()

	return &desc.GetPartResponse{
		Part: part,
	}, nil
}

func (s *server) ListParts(ctx context.Context, req *desc.ListPartsRequest) (*desc.ListPartsResponse, error) {
	return &desc.ListPartsResponse{
		Parts: []*desc.Part{
			mockPart(),
		},
	}, nil
}

func mockPart() *desc.Part {
	return &desc.Part{
		Uuid:          "11111111-1111-1111-1111-111111111111",
		Name:          "Двигатель Гипердрайва",
		Description:   "Ускоряет до гипера",
		Price:         999.99,
		StockQuantity: 5,
		Category:      desc.Category_CATEGORY_ENGINE,
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterInventoryServiceServer(s, &server{})
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

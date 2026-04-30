package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	desc "shared/pkg/proto/payment/v1"
	"syscall"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedPaymentServiceServer
}

func (s *server) PayOrder(ctx context.Context, req *desc.PayOrderRequest) (*desc.PayOrderResponse, error) {
	transactionUUID := uuid.New().String()

	fmt.Printf("Оплата прошла успешно, transaction_uuid: %s\n", transactionUUID)

	return &desc.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterPaymentServiceServer(s, &server{})

	reflection.Register(s)

	log.Printf("Payment Service запущен на %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Завершаем работу сервера...")
	s.GracefulStop()
	log.Println("Сервер успешно остановлен")
}

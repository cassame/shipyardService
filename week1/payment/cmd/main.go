package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	paymentApi "payment/internal/api/payment/v1"
	paymentSvc "payment/internal/service/payment"
	desc "shared/pkg/proto/payment/v1"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	svc := paymentSvc.NewService()
	impl := paymentApi.NewImplementation(svc)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	desc.RegisterPaymentServiceServer(s, impl)

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

package main

import (
	"context"
	"log"
	"net/http"
	"order/internal/migrator"
	"os"
	openapi "shared/pkg/openapi/order/v1"
	paymentv1 "shared/pkg/proto/payment/v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderApi "order/internal/api/order/v1"
	clientGrpc "order/internal/client/grpc"
	inventoryClient "order/internal/client/grpc/inventory/v1"
	paymentClient "order/internal/client/grpc/payment/v1"
	repo "order/internal/repository/order"
	orderService "order/internal/service/order"
	inventoryv1 "shared/pkg/proto/inventory/v1"
)

const (
	httpPort         = ":8082"
	inventoryAddress = "localhost:50051"
	paymentAddress   = "localhost:50053"
)

func main() {

	ctx := context.Background()

	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "postgresql://order-service-user:order-service-password@localhost:5432/order-service?sslmode=disable"
	}

	pool, err := pgxpool.New(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("db is unreachable: %v", err)
	}

	db := stdlib.OpenDB(*pool.Config().ConnConfig)
	m := migrator.NewMigrator(db, "migrations")
	if err := m.Up(); err != nil {
		_ = db.Close()
		log.Fatalf("failed to run migrations: %v", err)
	}
	_ = db.Close()

	log.Println("Миграции успешно применены!")

	invConn, err := grpc.NewClient(inventoryAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к Inventory: %v", err)
	}
	defer func() {
		_ = invConn.Close()
	}()

	payConn, err := grpc.NewClient(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к Payment: %v", err)
	}
	defer func() {
		_ = payConn.Close()
	}()

	grpcEngine := inventoryv1.NewInventoryServiceClient(invConn)
	payEngine := paymentv1.NewPaymentServiceClient(payConn)

	invCl := inventoryClient.NewClient(grpcEngine)
	payCl := paymentClient.NewClient(payEngine)

	allClients := clientGrpc.NewClients(invCl, payCl)

	repo := repo.NewRepository(pool)
	service := orderService.NewService(repo, allClients)
	apiImpl := orderApi.NewImplementation(service)

	srv, err := openapi.NewServer(apiImpl)
	if err != nil {
		log.Fatalf("ошибка при создании openapi сервера: %v", err)
	}

	srvWithLogs := logMiddleware(srv)
	log.Printf("🚀 Order Service (HTTP REST) запущен на %s", httpPort)
	if err := http.ListenAndServe(httpPort, srvWithLogs); err != nil {
		log.Fatalf("ошибка при запуске сервера: %v", err)
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📥 [HTTP] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

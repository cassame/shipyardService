package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	inventoryv1 "shared/pkg/proto/inventory/v1"
	paymentv1 "shared/pkg/proto/payment/v1"
	"strings"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	statusPendingPayment = "PENDING_PAYMENT"
	statusPaid           = "PAID"
	statusCancelled      = "CANCELLED"
)

type Order struct {
	UUID            string   `json:"uuid"`
	OrderUUID       string   `json:"order_uuid"`
	UserUUID        string   `json:"user_uuid"`
	PartUUIDs       []string `json:"part_uuids"`
	TotalPrice      float64  `json:"total_price"`
	TransactionUUID *string  `json:"transaction_uuid,omitempty"`
	PaymentMethod   *string  `json:"payment_method,omitempty"`
	Status          string   `json:"status"`
}

type Service struct {
	invClient inventoryv1.InventoryServiceClient
	payClient paymentv1.PaymentServiceClient

	mu     sync.RWMutex
	orders map[string]*Order
}

type CreateOrderRequest struct {
	UserUUID  string   `json:"user_uuid"`
	PartUUIDs []string `json:"part_uuids"`
}
type CreateOrderResponse struct {
	UUID       string  `json:"uuid"`
	OrderUUID  string  `json:"order_uuid"`
	TotalPrice float64 `json:"total_price"`
}

type PayOrderRequest struct {
	PaymentMethod string `json:"payment_method"`
}
type PayOrderResponse struct {
	TransactionUUID string `json:"transaction_uuid"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	//inv connect
	invConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Ошибка подключения к Inventory: %v", err)
	}
	defer invConn.Close()
	invClient := inventoryv1.NewInventoryServiceClient(invConn)

	//payment connect
	payConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Ошибка подключения к Payment: %v", err)
	}
	defer payConn.Close()
	payClient := paymentv1.NewPaymentServiceClient(payConn)

	svc := &Service{
		invClient: invClient,
		payClient: payClient,
		orders:    make(map[string]*Order),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/orders", svc.handleOrders)
	mux.HandleFunc("/api/v1/orders/", svc.handleOrderByUUID)

	log.Println("Order Service запущен на :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Ошибка запуска Order Service: %v", err)
	}
}

func (s *Service) handleOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	s.createOrder(w, r)
}

func (s *Service) handleOrderByUUID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/orders/")
	path = strings.Trim(path, "/")

	if path == "" {
		writeError(w, http.StatusBadRequest, "order uuid is required")
		return
	}

	parts := strings.Split(path, "/")
	orderUUID := parts[0]

	if len(parts) == 1 && r.Method == http.MethodGet {
		s.getOrder(w, r, orderUUID)
		return
	}

	if len(parts) == 2 && parts[1] == "pay" && r.Method == http.MethodPost {
		s.payOrder(w, r, orderUUID)
		return
	}

	if len(parts) == 2 && parts[1] == "cancel" && r.Method == http.MethodPost {
		s.cancelOrder(w, r, orderUUID)
		return
	}

	writeError(w, http.StatusNotFound, "not found")

}

func (s *Service) createOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.UserUUID == "" {
		writeError(w, http.StatusBadRequest, "user_uuid is required")
		return
	}

	if len(req.PartUUIDs) == 0 {
		writeError(w, http.StatusBadRequest, "part_uuids is required")
		return
	}

	partsResp, err := s.invClient.ListParts(r.Context(), &inventoryv1.ListPartsRequest{
		Filter: &inventoryv1.PartsFilter{
			Uuids: req.PartUUIDs,
		},
	})
	if err != nil {
		log.Printf("Inventory error: %v", err)
		writeError(w, http.StatusBadGateway, "inventory service error")
		return
	}

	if len(partsResp.GetParts()) == 0 {
		writeError(w, http.StatusBadRequest, "parts not found")
		return
	}

	var totalPrice float64

	for _, part := range partsResp.GetParts() {
		totalPrice += part.GetPrice()
	}

	orderUUID := uuid.New().String()

	order := &Order{
		UUID:       orderUUID,
		OrderUUID:  orderUUID,
		UserUUID:   req.UserUUID,
		PartUUIDs:  req.PartUUIDs,
		TotalPrice: totalPrice,
		Status:     statusPendingPayment,
	}

	s.mu.Lock()
	s.orders[orderUUID] = order
	s.mu.Unlock()

	writeJSON(w, http.StatusOK, CreateOrderResponse{
		UUID:       orderUUID,
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	})
}

func (s *Service) getOrder(w http.ResponseWriter, r *http.Request, orderUUID string) {
	s.mu.RLock()
	order, ok := s.orders[orderUUID]
	s.mu.RUnlock()

	if !ok {
		writeError(w, http.StatusNotFound, "order not found")
		return
	}

	writeJSON(w, http.StatusOK, order)
}

func (s *Service) payOrder(w http.ResponseWriter, r *http.Request, orderUUID string) {
	var req PayOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	paymentMethod, err := parsePaymentMethod(req.PaymentMethod)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.mu.RLock()
	order, ok := s.orders[orderUUID]
	s.mu.RUnlock()

	if !ok {
		writeError(w, http.StatusNotFound, "order not found")
		return
	}

	payResp, err := s.payClient.PayOrder(r.Context(), &paymentv1.PayOrderRequest{
		OrderUuid:     order.OrderUUID,
		UserUuid:      order.UserUUID,
		PaymentMethod: paymentMethod,
	})
	if err != nil {
		log.Printf("Payment error: %v", err)
		writeError(w, http.StatusBadGateway, "payment service error")
		return
	}
	transactionUUID := payResp.GetTransactionUuid()
	paymentMethodStr := req.PaymentMethod

	s.mu.Lock()
	order.Status = statusPaid
	order.TransactionUUID = &transactionUUID
	order.PaymentMethod = &paymentMethodStr
	s.mu.Unlock()

	writeJSON(w, http.StatusOK, PayOrderResponse{
		TransactionUUID: transactionUUID,
	})
}

func (s *Service) cancelOrder(w http.ResponseWriter, r *http.Request, orderUUID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[orderUUID]
	if !ok {
		writeError(w, http.StatusNotFound, "order not found")
		return
	}
	if order.Status == statusPaid {
		writeError(w, http.StatusConflict, "paid order cannot be cancelled")
		return
	}
	if order.Status == statusPendingPayment {
		order.Status = statusCancelled
	}
	log.Printf("Попытка отмены заказа: %s, текущий статус: %s", orderUUID, order.Status)

	w.WriteHeader(http.StatusNoContent)
}

func parsePaymentMethod(value string) (paymentv1.PaymentMethod, error) {
	switch value {
	case "CARD", "PAYMENT_METHOD_CARD":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_CARD, nil
	case "SBP", "PAYMENT_METHOD_SBP":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_SBP, nil
	case "CREDIT_CARD", "PAYMENT_METHOD_CREDIT_CARD":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD, nil
	case "INVESTOR_MONEY", "PAYMENT_METHOD_INVESTOR_MONEY":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY, nil
	default:
		return paymentv1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED, errors.New("unknown payment_method")
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, ErrorResponse{
		Error: message,
	})
}

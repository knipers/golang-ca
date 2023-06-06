package web

import (
	"encoding/json"
	"net/http"

	entity "github.com/knipers/golang-ca/internal/entity/order"
	"github.com/knipers/golang-ca/internal/usecase"
)

type OrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewOrderHandler(orderRepository entity.OrderRepositoryInterface) *OrderHandler {
	return &OrderHandler{
		OrderRepository: orderRepository,
	}
}

func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// get input
	var dto usecase.OrderInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(o.OrderRepository)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

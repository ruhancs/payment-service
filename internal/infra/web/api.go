package web

import (
	"log"
	"net/http"
	"payment-service/internal/application/usecase"
	"time"
)

type Application struct {
	GetOrderUseCase    *usecase.GetOrderUseCase
	UpdateOrderUseCase *usecase.UpdateOrderUseCase
}

func NewApplication(getOrderUsecase *usecase.GetOrderUseCase, updateOrderUsecase *usecase.UpdateOrderUseCase) *Application {
	return &Application{
		GetOrderUseCase: getOrderUsecase,
		UpdateOrderUseCase: updateOrderUsecase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8002",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}

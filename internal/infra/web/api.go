package web

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment-service/internal/application/usecase"
	"syscall"
	"time"
)

type Application struct {
	GetOrderUseCase *usecase.GetOrderUseCase
	PaymentUseCase  *usecase.PaymentUseCase
	SRV             *http.Server
	//WG              *sync.WaitGroup
}

func NewApplication(getOrderUsecase *usecase.GetOrderUseCase, paymentUseCase *usecase.PaymentUseCase,srv *http.Server) *Application {
	return &Application{
		GetOrderUseCase: getOrderUsecase,
		PaymentUseCase:  paymentUseCase,
		SRV: srv,
	}
}

func (app *Application) Server() error {
	app.SRV.Addr = ":8002"
	app.SRV.Handler = app.routes()
	app.SRV.IdleTimeout = 30 * time.Second
	app.SRV.ReadTimeout = 1 * time.Second
	app.SRV.ReadHeaderTimeout = 1 * time.Second
	app.SRV.WriteTimeout = 5 * time.Second

	log.Println("Runing server on port 8000...")
	return app.SRV.ListenAndServe()
}

func (app *Application) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//bloqueio ate se ter o sinal que terminou enviou as respostas para os usuarios
	<-quit
}

//func (app *Application) shutdown() {
//	fmt.Println("Cleaning all tasks")

//	app.WG.Wait()

//}

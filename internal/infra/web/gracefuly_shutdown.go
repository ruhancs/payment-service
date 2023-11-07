package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func(app *Application) GracefullyShutdown(chanError chan error) {
	//contexto ira receber os sinais de shutdown
	ctx,stop := signal.NotifyContext(context.Background(), os.Interrupt,syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		//ao receber o sinal de shutdown, o contexto recebe o done
		//para saber se o shutdown foi acionado
		<-ctx.Done()
		fmt.Println("Receive signal to shutdown, quitting...")

		//timeout de espera para fechar a aplicacao
		shutdownTimeout := 5 * time.Second
		//contexto com timeout
		ctxTimeout, cancel := context.WithTimeout(context.Background(),shutdownTimeout)

		defer func() {
			//stop para de escutar os sinais de desligar
			stop()
			//cancela o contexto com timeout
			cancel()
			//fechar o canal de erros
			close(chanError)
		}()

		//desligar o servidor
		err := app.SRV.Shutdown(ctxTimeout)
		if err != nil{
			return
		}
		fmt.Println("Shutdown Completed")
	}()

	go func ()  {
		if err := app.Server(); err != nil && err != http.ErrServerClosed {
			chanError <- fmt.Errorf("error while tryieng to start application, Error: %v", err)
		}
		return
	}()
}
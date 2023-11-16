package web

import (
	"encoding/json"
	"log"
	"net/http"
	"payment-service/internal/application/dto"
	"payment-service/pkg/worker"
	"time"

	"github.com/hibiken/asynq"
)

func(app *Application) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.PaymentInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	email := r.Context().Value("email").(string)
	inputDto.Email = email
	
	output,msg,err := app.PaymentUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	//criar payload para enviar tarefa de email em background
	taskPayload := &worker.PayloadSendEmail{
		Email: email,
	}
	// opcoes da tarefa, max de tentativas=10, tempo para iniciar processamento da tarefa 5 seg
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
	}
	err = app.TaskDistributor.DistributeTaskSendEmail(r.Context(),taskPayload,opts...)
	if err != nil {
		log.Printf("error distribute task to send email: %v",err.Error())
	}

	app.writeJson(w,http.StatusAccepted,output)
	app.writeJson(w,http.StatusAccepted,msg)
}
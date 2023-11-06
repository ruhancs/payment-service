package web

import (
	"encoding/json"
	"net/http"
	"payment-service/internal/application/dto"

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

	app.writeJson(w,http.StatusAccepted,output)
	app.writeJson(w,http.StatusAccepted,msg)
}
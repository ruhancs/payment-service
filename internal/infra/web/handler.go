package web

import (
	"encoding/json"
	"net/http"
	"payment-service/internal/application/dto"

	"github.com/go-chi/chi"
)

func(app *Application) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderID")
	var inputDto dto.PaymentInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	output,msg,err := app.UpdateOrderUseCase.Execute(r.Context(),orderID,inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	app.writeJson(w,http.StatusAccepted,output)
	app.writeJson(w,http.StatusAccepted,msg)
}
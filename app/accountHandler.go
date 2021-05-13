package app

import (
	"encoding/json"
	"net/http"

	"github.com/dapinto8/banking/dto"
	"github.com/dapinto8/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeRespoonse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId

	account, appError := ah.service.NewAccount(request)
	if appError != nil {
		writeRespoonse(w, appError.Code, appError.AsMessage())
		return
	}

	writeRespoonse(w, http.StatusAccepted, account)
}

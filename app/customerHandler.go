package app

import (
	"net/http"

	"github.com/dapinto8/banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeRespoonse(w, err.Code, err.AsMessage())
		return
	}

	writeRespoonse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeRespoonse(w, err.Code, err.AsMessage())
		return
	}

	writeRespoonse(w, http.StatusOK, customer)
}

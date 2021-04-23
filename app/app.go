package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dapinto8/banking/domain"
	"github.com/dapinto8/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryMysql())}

	router.HandleFunc("/health", health)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8200", router))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "health")
}

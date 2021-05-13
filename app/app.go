package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dapinto8/banking/domain"
	"github.com/dapinto8/banking/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Start() {

	loadEnv()

	router := mux.NewRouter()

	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryMysql())}

	router.HandleFunc("/health", health)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "health")
}

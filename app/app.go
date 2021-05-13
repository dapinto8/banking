package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dapinto8/banking/domain"
	"github.com/dapinto8/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	loadEnv()

	dbClient := getDbClient()
	customerRepositoryMysql := domain.NewCustomerRepositoryMysql(dbClient)
	accountRepositoryMysql := domain.NewAccountRepositoryMysql(dbClient)
	ch := CustomerHandler{service.NewCustomerService(customerRepositoryMysql)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryMysql)}

	router := mux.NewRouter()
	router.HandleFunc("/health", health)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "health")
}

package app

import (
	"bank-project/domain"
	"bank-project/service"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func GetStarted() {
	handler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router := mux.NewRouter()
	router.HandleFunc("/customerdetails", handler.CustomerDetails)
	http.ListenAndServe("localhost:8080", router)
}

func (ch *CustomerHandler) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	customerDetails, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customerDetails)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customerDetails)
	}
}

package main

import (
	"CiscoApr2023/transactionweb/stores"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var transactions []stores.Transaction

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	// Read-all
	router.HandleFunc("/transactions", getTransactionHandler).Methods("GET")

	// Initialize db connection
	stores.ConnectionHelperORM()

	//router.PathPrefix("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(http.ListenAndServe(":7072", router))
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Transaction
	transaction := stores.Transaction{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the department from the form info
	id, _ := strconv.Atoi(r.Form.Get("transactionId"))
	transaction.TransactionId = int64(id)
	transaction.TransactionDate = r.Form.Get("transactionDate")
	amount, _ := strconv.Atoi(r.Form.Get("amount"))
	transaction.Amount = int64(amount)
	// Append our existing list of departments with a new entry
	transactions = append(transactions, transaction)
	stores.SaveTransaction(&transaction)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func getTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := stores.GetAllTransactions()
	//Convert the "departments" variable to json
	transactionsListBytes, err := json.Marshal(transactions)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(transactionsListBytes)
}

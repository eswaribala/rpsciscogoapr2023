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

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	// Create
	router.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	// Read-all
	router.HandleFunc("/transactions", getTransactionHandler).Methods("GET")

	// Initialize db connection
	stores.ConnectionHelperORM()

	//router.PathPrefix("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(http.ListenAndServe(":7074", router))
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
	id, _ := strconv.Atoi(r.Form.Get("transaction_id"))
	transaction.TransactionId = int64(id)
	transaction.TransactionDate = r.Form.Get("transaction_date")
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

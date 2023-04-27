package main

import (
	"github.com/gorilla/mux"

	_ "CiscoApr2023/day4/docs"

	"CiscoApr2023/day4/stores"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Customer API
// @version 1.0
// @description This is api service for managing Customer
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7072
// @BasePath /
func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/customers", stores.SaveCustomer).Methods("POST")
	// Read
	router.HandleFunc("/customers/{accountNo}", stores.GetCustomerById).Methods("GET")
	// Read-all
	router.HandleFunc("/customers", stores.GetCustomers).Methods("GET")
	// Update
	//router.HandleFunc("/customers/{accountNo}", stores.UpdateCustomer).Methods("PUT")
	// Delete
	router.HandleFunc("/customers/{accountNo}", stores.DeleteCustomerById).Methods("DELETE")
	router.HandleFunc("/employees", stores.GetEmployees).Methods("GET")
	// Initialize db connection
	stores.InitDB()

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//router.PathPrefix("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(http.ListenAndServe(":6064", router))
}

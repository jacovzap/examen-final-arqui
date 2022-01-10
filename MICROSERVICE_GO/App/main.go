package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	repository "../Persistence"

)



func main(){
	router := mux.NewRouter()
	
	
	//endpoins
	router.HandleFunc("/products", repository.GetProductsEndPoint).Methods("GET")
	router.HandleFunc("/products/{name}", repository.GetProductEndPoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
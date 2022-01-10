package repository

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	data "../Data"
	product "../DomainModel"
)


func GetProductsEndPoint(w http.ResponseWriter, r *http.Request){
	data.Respository()
	json.NewEncoder(w).Encode(data.Products)
}

func GetProductEndPoint(w http.ResponseWriter, r *http.Request){
	data.Respository()
	params := mux.Vars(r)
	for _, item := range data.Products {
	  if item.Name == params["name"] {
		json.NewEncoder(w).Encode(item)
		return
	  }
	}
	json.NewEncoder(w).Encode(&product.Product{})
}
package data

import(
	product "../DomainModel"
)



var Products []product.Product


func Respository(){
	Products = append(Products, product.Product{Name: "Tv", Price: 5200, Amount: 15})
	Products = append(Products, product.Product{Name: "Computador", Price: 7000, Amount: 12})
	Products = append(Products, product.Product{Name: "Refigerador", Price: 3000, Amount: 23})
	Products = append(Products, product.Product{Name: "Cocina", Price: 2000, Amount: 10})
	Products = append(Products, product.Product{Name: "Lavadora", Price: 3000, Amount: 6})
	Products = append(Products, product.Product{Name: "Secadora", Price: 3000, Amount: 7})
}


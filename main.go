package main

import (
	"go-web/config"
	"go-web/controllers/categoriescontroller"
	"go-web/controllers/homecontrollers"
	"go-web/controllers/productscontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontrollers.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	// 3, Products
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/products/detail", productscontroller.Detail)
	http.HandleFunc("/products/add", productscontroller.Add)
	http.HandleFunc("/products/edit", productscontroller.Edit)
	http.HandleFunc("/products/delete", productscontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

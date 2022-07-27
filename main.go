package main

import (
	"log"
	"net/http"

	"github.com/djfemz/go-web-intro-2/controllers"
)

func main() {
	http.HandleFunc("/add-product", controllers.CreateProduct)
	http.HandleFunc("/delete-product/", controllers.DeleteProduct)
	http.HandleFunc("/products/all", controllers.GetAllProducts)
	log.Println("listening...")
	err := http.ListenAndServe(":9103", nil)
	if err != nil {
		log.Fatal(err)
	}
}

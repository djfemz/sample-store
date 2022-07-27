package controllers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/djfemz/go-web-intro-2/data/dtos"
	"github.com/djfemz/go-web-intro-2/data/repositories"
	"github.com/djfemz/go-web-intro-2/services"
	"github.com/djfemz/go-web-intro-2/utils"
)

func CreateProduct(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(rw, "Method not allowed", http.StatusBadRequest)
		return
	}
	product := utils.JsonToProduct(req.Body)
	_, err := services.AddProduct(&product)
	if err != nil {
		log.Fatal("error ", err)
	}
	bs := utils.ProductToJson(&dtos.ApiResponse{
		Message: "product added successfully",
		Code:    201,
		Products: repositories.ProductList,
	})
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bs)
}

func DeleteProduct(rw http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodDelete{
		http.Error(rw,"Request not allowed", http.StatusBadRequest)
		return
	}

	r, err:=regexp.Compile("[1-9]+")
	if err!=nil {
		log.Fatal("error: ", err)
	}
	captureGroup:=r.FindStringSubmatch(req.URL.Path)
	log.Println(captureGroup)
	id, err:=strconv.Atoi(captureGroup[0])
	log.Printf("%d:::%T", id, id)
	if err!=nil {
		log.Fatal("error: ", err)
	}
	err =services.DeleteById(id)
	if err!=nil {
		log.Fatal("error: ", err)
	}
}

func GetAllProducts(rw http.ResponseWriter, req *http.Request)  {
	products:=services.GetAllProducts()
	response:=&dtos.ApiResponse{
		Message: "successful",
		Code: http.StatusOK,
		Products: products,
	}
	res:=utils.ProductToJson(response)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}

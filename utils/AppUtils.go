package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/djfemz/go-web-intro-2/data/dtos"
	"github.com/djfemz/go-web-intro-2/data/models"
)

func JsonToProduct(requestBody io.Reader) models.Product {
	var product models.Product
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		log.Fatal("error reading from request ", err)
	}
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Fatal("error reading from request ", err)
	}
	return product
}

func ProductToJson(response *dtos.ApiResponse) []byte{
	jsonResponse, err:= json.Marshal(response)
	if err!=nil{
		log.Fatal(err)
	}
	return jsonResponse
}

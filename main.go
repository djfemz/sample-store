package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var ProductList []*Product

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type ApiResponse struct{
	Message string `json:"message"`
	Code int `json:"code"`
	Product Product `json:"-"`
}
//TODOS
//CRUD FUNCTIONS
//expose them as endpoints
func main() {
	http.HandleFunc("/add-product", createProduct)
	log.Println("listening...")
	err:=http.ListenAndServe(":9090", nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func createProduct(rw http.ResponseWriter, req *http.Request)  {
	var product = &Product{}
	if req.Method != http.MethodPost{
		apiResponse:=&ApiResponse{
			Message: "method not allowed",
			Code: http.StatusBadRequest,
		}
		bs, err:=json.Marshal(apiResponse)
		if err!=nil{
			log.Fatal("error serializing object apiResponse ", err)
		}
		fmt.Fprint(rw, bs)
	}	

	body, err:=ioutil.ReadAll(req.Body)
	if err!=nil{
		log.Fatal("error reading from request ", err)
	}
	json.Unmarshal(body, product)
	id, err:=AddProduct(product)
	if err!=nil{
		log.Fatal("error ", err)
	}
	fmt.Fprint(rw, &ApiResponse{
		Message: "product added successfully",
		Code: 201,
		Product: *ProductList[id-1],
	})
}

func AddProduct(product *Product) (id int, err error){
	initialProductCount:= len(ProductList)
	id, err = generateId()
	if err!= nil{
		err = errors.New("failed to generate id")
		return 
	}
	product.Id = id
	ProductList=append(ProductList, product)
	if len(ProductList) != initialProductCount+1{
		err=errors.New("failed to add product to store")
		return
	}
	fmt.Println(ProductList)
	fmt.Println(dbToString())
	return product.Id, nil
}

func generateId() (id int, err error){
	if len(ProductList) < 1{
		id = len(ProductList) + 1
		err=nil
		return
	}
	for index, product := range ProductList {
		if index == len(ProductList)-1{
			id=product.Id+1
			err = nil
			return 
		}
	}
	return -1, errors.New("failed to create id")
}

func DeleteById(id int) (err error){
	//check for product in the list
	var ProductListToReturn []*Product
	for _, product := range ProductList {
		if product.Id!=id{
			ProductListToReturn=append(ProductListToReturn, product)
		}
	}
	if len(ProductListToReturn)==len(ProductList) {
		return errors.New("failed to delete Product")
	}
	ProductList=ProductListToReturn
	return nil
}


func dbToString() string{
	s:="["
	for i, product := range ProductList {
		if i!=0{
			s+=","
		}
		s+=fmt.Sprint("{"+strconv.Itoa(product.Id)+" " + product.Name +" "+ product.Description +"}")
	}
	s+="]"
	return s
}
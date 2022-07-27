package main

import (
	"testing"

	"github.com/djfemz/go-web-intro-2/data/models"
	"github.com/djfemz/go-web-intro-2/data/repositories"
	"github.com/djfemz/go-web-intro-2/services"
	"github.com/stretchr/testify/assert"
)

var testProduct *models.Product
var testProduct1 *models.Product

func setUp() {
	testProduct = &models.Product{
		Name: "test product",
		Description: "a test product",
		Price: 20.00,
		Quantity: 2,
	}

	testProduct1= &models.Product{
		Name: "test product",
		Description: "a test product",
		Price: 20.00,
		Quantity: 2,
	}
}

func TestThatProductCanBeCreated(t *testing.T) {
	setUp()
	id, err:=services.AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "error adding product")
	}

	assert.Equal(t, 1, id)
}

func TestThatTwoProductsCanBeCreated(t *testing.T) {
	setUp()
	id, err:=services.AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "error adding product")
	}
	id1, err:=services.AddProduct(testProduct1)

	if err!=nil{
		assert.Fail(t, "error adding product")
	}
	assert.Equal(t, 1, id)
	assert.Equal(t, 2, id1)
}

func TestThatItemCanBeDeleted(t *testing.T) {
	setUp()
	//given that we have two items
	_, err:=services.AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "add failed")
	}
	id2, err:=services.AddProduct(testProduct1)
	if err!=nil{
		assert.Fail(t, "add failed")
	}
	//when we remove one product
	errr:=services.DeleteById(id2)
	if errr!=nil{
		assert.Fail(t, "add failed")
	}
	//assert that product was removed
	assert.Equal(t, 1, len(repositories.ProductList))
	indexOfLastProduct:=len(repositories.ProductList)-1
	lastProduct:=repositories.ProductList[indexOfLastProduct]
	assert.Equal(t, 1, lastProduct.Id)
}



package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testProduct *Product
var testProduct1 *Product

func setUp() {
	testProduct = &Product{
		Name: "test product",
		Description: "a test product",
		Price: 20.00,
		Quantity: 2,
	}

	testProduct1= &Product{
		Name: "test product",
		Description: "a test product",
		Price: 20.00,
		Quantity: 2,
	}
}

func TestThatProductCanBeCreated(t *testing.T) {
	setUp()
	id, err:=AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "error adding product")
	}

	assert.Equal(t, 1, id)
}

func TestThatTwoProductsCanBeCreated(t *testing.T) {
	setUp()
	id, err:=AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "error adding product")
	}
	id1, err:=AddProduct(testProduct1)

	if err!=nil{
		assert.Fail(t, "error adding product")
	}
	assert.Equal(t, 1, id)
	assert.Equal(t, 2, id1)
}

func TestThatItemCanBeDeleted(t *testing.T) {
	setUp()
	//given that we have two items
	_, err:=AddProduct(testProduct)
	if err!=nil{
		assert.Fail(t, "add failed")
	}
	id2, err:=AddProduct(testProduct1)
	if err!=nil{
		assert.Fail(t, "add failed")
	}
	//when we remove one product
	errr:=DeleteById(id2)
	if errr!=nil{
		assert.Fail(t, "add failed")
	}
	//assert that product was removed
	assert.Equal(t, 1, len(ProductList))
	indexOfLastProduct:=len(ProductList)-1
	lastProduct:=ProductList[indexOfLastProduct]
	assert.Equal(t, 1, lastProduct.Id)
}

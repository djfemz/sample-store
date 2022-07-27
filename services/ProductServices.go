package services

import (
	"errors"

	"github.com/djfemz/go-web-intro-2/data/models"
	"github.com/djfemz/go-web-intro-2/data/repositories"
)



func AddProduct(product *models.Product) (id int, err error) {
	initialProductCount := len(repositories.ProductList)
	id, err = generateId()
	if err != nil {
		err = errors.New("failed to generate id")
		return
	}
	product.Id = id
	repositories.ProductList = append(repositories.ProductList, product)
	if len(repositories.ProductList) != initialProductCount+1 {
		err = errors.New("failed to add product to store")
		return
	}
	return product.Id, nil
}


func generateId() (id int, err error) {
	if len(repositories.ProductList) < 1 {
		id = len(repositories.ProductList) + 1
		err = nil
		return
	}
	for index, product := range repositories.ProductList {
		if index == len(repositories.ProductList)-1 {
			id = product.Id + 1
			err = nil
			return
		}
	}
	return -1, errors.New("failed to create id")
}

func GetAllProducts() []*models.Product {
	return repositories.ProductList
}

func DeleteById(id int) (err error) {
	//list that will hold all products that are not to be deleted
	var newList []*models.Product
	//add products that are not to be deleted to new list
	for _, product := range repositories.ProductList {
		if product.Id != id {
			newList = append(newList, product)
		}
	}
	if len(newList) == len(repositories.ProductList) {
		return errors.New("failed to delete Product line 55")
	}
	//update ProductList to be new list
	repositories.ProductList = newList
	return nil
}
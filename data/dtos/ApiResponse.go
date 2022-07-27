package dtos

import "github.com/djfemz/go-web-intro-2/data/models"


type ApiResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Products []*models.Product `json:"products"`
}
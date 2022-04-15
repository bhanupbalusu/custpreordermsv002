package service

import (
	model "github.com/bhanupbalusu/custpreordermsv002/domain/model/product_details"
)

type ProductDetailsServiceInterface interface {
	Get() (*[]model.ProductDetailsModel, error)
	GetByID(id string) (*model.ProductDetailsModel, error)
	Create(pm *model.ProductDetailsModel) error
	Update(pm *model.ProductDetailsModel) error
	Delete(pid string) error
}

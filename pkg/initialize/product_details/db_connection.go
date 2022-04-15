package product_details

import (
	"fmt"

	pddb "github.com/bhanupbalusu/custpreordermsv002/data/mongodb/product_details"
	r "github.com/bhanupbalusu/custpreordermsv002/domain/app_interface/repo"
)

func NewDBConnection() r.ProductDetailsRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo := pddb.NewMongoRepository(pddb.NewConnection("custpreordermsv002"), "custpreordermsv002")
	return repo
}

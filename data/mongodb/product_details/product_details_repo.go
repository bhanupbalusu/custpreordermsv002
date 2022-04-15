package product_details

import (
	repo "github.com/bhanupbalusu/custpreordermsv002/domain/app_interface/repo"
	model "github.com/bhanupbalusu/custpreordermsv002/domain/model/product_details"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ProductDetailsCollection = "product_details"

type MongoRepository struct {
	c *mgo.Collection
}

func NewMongoRepository(conn Connection, db string) repo.ProductDetailsRepoInterface {
	return &MongoRepository{conn.DB(db).C(ProductDetailsCollection)}
}

// Get all products
func (r *MongoRepository) Get() (*[]model.ProductDetailsModel, error) {
	var results []model.ProductDetailsModel
	err := r.c.Find(bson.M{}).All(&results)
	return &results, err
}

// Get single product using id
func (r *MongoRepository) GetByID(id string) (*model.ProductDetailsModel, error) {
	var result model.ProductDetailsModel
	err := r.c.FindId(bson.ObjectIdHex(id)).One(&result)
	return &result, err
}

// Create or insert a new product
func (r *MongoRepository) Create(pm *model.ProductDetailsModel) error {

	return r.c.Insert(pm)
}

// Update existing product
func (r *MongoRepository) Update(pm *model.ProductDetailsModel) error {
	return r.c.UpdateId(pm.ProductId, pm)
}

// Delete an existing product
func (r *MongoRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

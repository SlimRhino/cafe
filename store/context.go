package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//DAO - Data Access Model
type dao struct {
	collection string
	db         *mgo.Database
}

func (d *dao) findByID(id string) (*Product, error) {
	var product Product
	err := d.db.C(d.collection).Find(bson.M{"_id": id}).One(&product)
	return &product, err
}

func (d *dao) add(product *Product) error {
	product.ID = bson.NewObjectId().Hex()
	err := d.db.C(d.collection).Insert(&product)
	return err
}

func (d *dao) delete(id string) error {
	err := d.db.C(d.collection).RemoveId(id)
	return err
}
